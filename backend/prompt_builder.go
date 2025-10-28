package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// PromptBuilder Prompt构建器
type PromptBuilder struct {
	db *sql.DB
}

// NewPromptBuilder 创建Prompt构建器
func NewPromptBuilder(db *sql.DB) *PromptBuilder {
	return &PromptBuilder{db: db}
}

// BuildPlanPrompt 构建生成计划的Prompt
func (pb *PromptBuilder) BuildPlanPrompt(date string) (string, string, error) {
	// 1. 获取所有目标
	goals, err := pb.getGoals()
	if err != nil {
		log.Printf("❌ 获取目标失败: %v", err)
		return "", "", err
	}

	// 2. 获取时间规则
	timeRules, err := pb.getTimeRules()
	if err != nil {
		log.Printf("❌ 获取时间规则失败: %v", err)
		return "", "", err
	}

	// 3. 获取最近7天的学习记录
	recentLogs, err := pb.getRecentLogs(7)
	if err != nil {
		log.Printf("❌ 获取学习记录失败: %v", err)
		return "", "", err
	}

	// 4. 提取已学习的内容
	learnedContent := pb.extractLearnedContent(recentLogs)

	// 5. 构建System Prompt
	systemPrompt := `你是一个专业的学习计划生成助手。
你的任务是根据用户的学习目标、时间安排和历史学习记录，生成一份详细的今日学习计划。

要求：
1. 计划要具体、可执行、有时间安排
2. 避免重复已学习的内容
3. 合理分配时间，不超过用户的可用时间
4. 考虑用户的学习进度和目标优先级
5. 返回有效的JSON格式

返回格式必须是以下JSON结构：
{
  "date": "YYYY-MM-DD",
  "summary": "今日计划总结",
  "tasks": [
    {
      "goal_id": 1,
      "title": "任务标题",
      "description": "详细描述",
      "duration_minutes": 60,
      "start_time": "HH:MM",
      "priority": "high|medium|low"
    }
  ],
  "total_duration_minutes": 180,
  "notes": "额外说明"
}`

	// 6. 构建User Prompt
	userPrompt := fmt.Sprintf(`
【用户学习目标】
%s

【今日可用时间段】
%s

【最近7天的学习记录】
%s

【已学习的内容（避免重复）】
%s

请为用户生成今日（%s）的学习计划。
返回格式必须是有效的JSON，不要包含任何其他文本。
`,
		pb.formatGoals(goals),
		pb.formatTimeRules(timeRules),
		pb.formatLogs(recentLogs),
		learnedContent,
		date,
	)

	return systemPrompt, userPrompt, nil
}

// getGoals 获取所有目标
func (pb *PromptBuilder) getGoals() ([]Goal, error) {
	rows, err := pb.db.Query("SELECT id, name, description, status, deadline FROM goals WHERE status='active' ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []Goal
	for rows.Next() {
		var goal Goal
		if err := rows.Scan(&goal.ID, &goal.Name, &goal.Description, &goal.Status, &goal.Deadline); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		goals = append(goals, goal)
	}

	return goals, nil
}

// getTimeRules 获取时间规则
func (pb *PromptBuilder) getTimeRules() ([]TimeRule, error) {
	// 获取今天的时间规则
	today := int(time.Now().Weekday())
	rows, err := pb.db.Query("SELECT id, day_of_week, start_time, end_time FROM time_rules WHERE day_of_week=? ORDER BY start_time", today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []TimeRule
	for rows.Next() {
		var rule TimeRule
		if err := rows.Scan(&rule.ID, &rule.DayOfWeek, &rule.StartTime, &rule.EndTime); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		rules = append(rules, rule)
	}

	return rules, nil
}

// getRecentLogs 获取最近N天的学习记录
func (pb *PromptBuilder) getRecentLogs(days int) ([]LearningLog, error) {
	query := `
		SELECT id, goal_id, content, duration, log_date 
		FROM learning_logs 
		WHERE log_date >= date('now', '-' || ? || ' days')
		ORDER BY log_date DESC, created_at DESC
	`
	rows, err := pb.db.Query(query, days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []LearningLog
	for rows.Next() {
		var learningLog LearningLog
		if err := rows.Scan(&learningLog.ID, &learningLog.GoalID, &learningLog.Content, &learningLog.Duration, &learningLog.LogDate); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		logs = append(logs, learningLog)
	}

	return logs, nil
}

// extractLearnedContent 提取已学习的内容
func (pb *PromptBuilder) extractLearnedContent(logs []LearningLog) string {
	if len(logs) == 0 {
		return "暂无学习记录"
	}

	// 使用map去重（基于内容哈希）
	contentMap := make(map[string]bool)
	var uniqueContents []string

	for _, log := range logs {
		// 计算内容哈希
		hash := pb.hashContent(log.Content)
		if !contentMap[hash] {
			contentMap[hash] = true
			// 格式化内容
			content := fmt.Sprintf("- [%s] %s", log.LogDate, log.Content)
			if log.Duration != nil {
				content += fmt.Sprintf(" (耗时: %d分钟)", *log.Duration)
			}
			uniqueContents = append(uniqueContents, content)
		}
	}

	if len(uniqueContents) == 0 {
		return "暂无学习记录"
	}

	return strings.Join(uniqueContents, "\n")
}

// hashContent 计算内容哈希
func (pb *PromptBuilder) hashContent(content string) string {
	// 简单的MD5哈希，用于去重
	hash := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", hash)
}

// formatGoals 格式化目标
func (pb *PromptBuilder) formatGoals(goals []Goal) string {
	if len(goals) == 0 {
		return "暂无学习目标"
	}

	var result []string
	for _, g := range goals {
		item := fmt.Sprintf("- %s: %s", g.Name, g.Description)
		if g.Deadline != nil {
			item += fmt.Sprintf(" (截止日期: %s)", *g.Deadline)
		}
		result = append(result, item)
	}

	return strings.Join(result, "\n")
}

// formatTimeRules 格式化时间规则
func (pb *PromptBuilder) formatTimeRules(rules []TimeRule) string {
	if len(rules) == 0 {
		return "暂无时间规则设置"
	}

	var result []string
	for _, r := range rules {
		result = append(result, fmt.Sprintf("- %s: %s - %s", getDayName(r.DayOfWeek), r.StartTime, r.EndTime))
	}

	return strings.Join(result, "\n")
}

// formatLogs 格式化学习记录
func (pb *PromptBuilder) formatLogs(logs []LearningLog) string {
	if len(logs) == 0 {
		return "暂无学习记录"
	}

	var result []string
	for _, l := range logs {
		item := fmt.Sprintf("- [%s] %s", l.LogDate, l.Content)
		if l.Duration != nil {
			item += fmt.Sprintf(" (耗时: %d分钟)", *l.Duration)
		}
		result = append(result, item)
	}

	return strings.Join(result, "\n")
}

// getDayName 获取星期名称
func getDayName(day int) string {
	days := []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	if day >= 0 && day < len(days) {
		return days[day]
	}
	return "未知"
}
