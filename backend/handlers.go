package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ===== 目标管理 =====

// getGoals 获取所有目标
func getGoals(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, description, status, deadline, created_at, updated_at FROM goals ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	defer rows.Close()

	var goals []Goal
	for rows.Next() {
		var goal Goal
		if err := rows.Scan(&goal.ID, &goal.Name, &goal.Description, &goal.Status, &goal.Deadline, &goal.CreatedAt, &goal.UpdatedAt); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		goals = append(goals, goal)
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: goals})
}

// createGoal 创建新目标
func createGoal(c *gin.Context) {
	var goal Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误: " + err.Error()})
		return
	}

	goal.Status = "active"
	goal.CreatedAt = time.Now()
	goal.UpdatedAt = time.Now()

	result, err := db.Exec(
		"INSERT INTO goals (name, description, status, deadline, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		goal.Name, goal.Description, goal.Status, goal.Deadline, goal.CreatedAt, goal.UpdatedAt,
	)
	if err != nil {
		log.Printf("插入失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败"})
		return
	}

	id, _ := result.LastInsertId()
	goal.ID = int(id)

	c.JSON(http.StatusOK, Response{Code: 0, Message: "创建成功", Data: goal})
}

// updateGoal 更新目标
func updateGoal(c *gin.Context) {
	id := c.Param("id")
	var goal Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}

	goal.UpdatedAt = time.Now()

	_, err := db.Exec(
		"UPDATE goals SET name=?, description=?, status=?, deadline=?, updated_at=? WHERE id=?",
		goal.Name, goal.Description, goal.Status, goal.Deadline, goal.UpdatedAt, id,
	)
	if err != nil {
		log.Printf("更新失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "更新失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "更新成功"})
}

// deleteGoal 删除目标
func deleteGoal(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM goals WHERE id=?", id)
	if err != nil {
		log.Printf("删除失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "删除成功"})
}

// ===== 时间规则 =====

// getTimeRules 获取时间规则
func getTimeRules(c *gin.Context) {
	rows, err := db.Query("SELECT id, day_of_week, start_time, end_time, created_at, updated_at FROM time_rules ORDER BY day_of_week")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	defer rows.Close()

	var rules []TimeRule
	for rows.Next() {
		var rule TimeRule
		if err := rows.Scan(&rule.ID, &rule.DayOfWeek, &rule.StartTime, &rule.EndTime, &rule.CreatedAt, &rule.UpdatedAt); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		rules = append(rules, rule)
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: rules})
}

// setTimeRules 设置时间规则
func setTimeRules(c *gin.Context) {
	var rule TimeRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}

	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()

	// 先删除该天的规则，再插入新规则
	db.Exec("DELETE FROM time_rules WHERE day_of_week=?", rule.DayOfWeek)

	result, err := db.Exec(
		"INSERT INTO time_rules (day_of_week, start_time, end_time, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		rule.DayOfWeek, rule.StartTime, rule.EndTime, rule.CreatedAt, rule.UpdatedAt,
	)
	if err != nil {
		log.Printf("插入失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "设置失败"})
		return
	}

	id, _ := result.LastInsertId()
	rule.ID = int(id)

	c.JSON(http.StatusOK, Response{Code: 0, Message: "设置成功", Data: rule})
}

// ===== 学习记录 =====

// getLogs 获取学习记录
func getLogs(c *gin.Context) {
	// 支持按目标ID筛选
	goalID := c.Query("goal_id")
	logDate := c.Query("log_date")

	query := "SELECT id, goal_id, content, duration, log_date, created_at, updated_at FROM learning_logs WHERE 1=1"
	var args []interface{}

	if goalID != "" {
		query += " AND goal_id=?"
		args = append(args, goalID)
	}

	if logDate != "" {
		query += " AND log_date=?"
		args = append(args, logDate)
	}

	query += " ORDER BY log_date DESC, created_at DESC"

	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	defer rows.Close()

	var logs []LearningLog
	for rows.Next() {
		var learningLog LearningLog
		if err := rows.Scan(&learningLog.ID, &learningLog.GoalID, &learningLog.Content, &learningLog.Duration, &learningLog.LogDate, &learningLog.CreatedAt, &learningLog.UpdatedAt); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		logs = append(logs, learningLog)
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: logs})
}

// createLog 创建学习记录
func createLog(c *gin.Context) {
	var learningLog LearningLog
	if err := c.ShouldBindJSON(&learningLog); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}

	learningLog.CreatedAt = time.Now()
	learningLog.UpdatedAt = time.Now()

	result, err := db.Exec(
		"INSERT INTO learning_logs (goal_id, content, duration, log_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		learningLog.GoalID, learningLog.Content, learningLog.Duration, learningLog.LogDate, learningLog.CreatedAt, learningLog.UpdatedAt,
	)
	if err != nil {
		log.Printf("插入失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败"})
		return
	}

	id, _ := result.LastInsertId()
	learningLog.ID = int(id)

	c.JSON(http.StatusOK, Response{Code: 0, Message: "创建成功", Data: learningLog})
}

// ===== 计划 =====

// getTodayPlan 获取今日计划（使用LLM生成）
func getTodayPlan(c *gin.Context) {
	today := time.Now().Format("2006-01-02")

	// 1. 先查询是否已有缓存的计划（24小时内）
	var existingPlan Plan
	err := db.QueryRow(
		"SELECT id, goal_id, plan_date, content, status, created_at, updated_at FROM plans WHERE plan_date=? LIMIT 1",
		today,
	).Scan(&existingPlan.ID, &existingPlan.GoalID, &existingPlan.PlanDate, &existingPlan.Content, &existingPlan.Status, &existingPlan.CreatedAt, &existingPlan.UpdatedAt)

	if err == nil {
		// 缓存命中，直接返回
		log.Printf("✅ 使用缓存的计划")
		c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: existingPlan})
		return
	}

	// 2. 缓存未命中，调用LLM生成计划
	log.Printf("🔄 生成新的学习计划...")

	// 构建Prompt
	promptBuilder := NewPromptBuilder(db)
	systemPrompt, userPrompt, err := promptBuilder.BuildPlanPrompt(today)
	if err != nil {
		log.Printf("❌ 构建Prompt失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "构建计划失败: " + err.Error()})
		return
	}

	// 调用LLM
	llmService := NewLLMService()
	planContent, err := llmService.Generate(systemPrompt, userPrompt)
	if err != nil {
		log.Printf("❌ LLM生成计划失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "生成计划失败: " + err.Error()})
		return
	}

	log.Printf("✅ LLM生成计划成功")

	// 3. 保存计划到数据库
	result, err := db.Exec(
		"INSERT INTO plans (goal_id, plan_date, content, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		1, today, planContent, "active", time.Now(), time.Now(),
	)
	if err != nil {
		log.Printf("❌ 保存计划失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "保存计划失败: " + err.Error()})
		return
	}

	lastID, _ := result.LastInsertId()
	newPlan := Plan{
		ID:        int(lastID),
		GoalID:    1,
		PlanDate:  today,
		Content:   planContent,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: newPlan})
}

// getPlan 获取指定日期的计划
func getPlan(c *gin.Context) {
	planDate := c.Query("date")
	if planDate == "" {
		planDate = time.Now().Format("2006-01-02")
	}

	rows, err := db.Query(
		"SELECT id, goal_id, plan_date, content, status, created_at, updated_at FROM plans WHERE plan_date=? ORDER BY created_at",
		planDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	defer rows.Close()

	var plans []Plan
	for rows.Next() {
		var plan Plan
		if err := rows.Scan(&plan.ID, &plan.GoalID, &plan.PlanDate, &plan.Content, &plan.Status, &plan.CreatedAt, &plan.UpdatedAt); err != nil {
			log.Printf("扫描行失败: %v", err)
			continue
		}
		plans = append(plans, plan)
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "success", Data: plans})
}

// createPlan 创建计划
func createPlan(c *gin.Context) {
	var plan Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}

	plan.Status = "pending"
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	result, err := db.Exec(
		"INSERT INTO plans (goal_id, plan_date, content, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		plan.GoalID, plan.PlanDate, plan.Content, plan.Status, plan.CreatedAt, plan.UpdatedAt,
	)
	if err != nil {
		log.Printf("插入失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败"})
		return
	}

	id, _ := result.LastInsertId()
	plan.ID = int(id)

	c.JSON(http.StatusOK, Response{Code: 0, Message: "创建成功", Data: plan})
}

// updatePlan 更新计划
func updatePlan(c *gin.Context) {
	id := c.Param("id")
	var plan Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}

	plan.UpdatedAt = time.Now()

	_, err := db.Exec(
		"UPDATE plans SET goal_id=?, plan_date=?, content=?, status=?, updated_at=? WHERE id=?",
		plan.GoalID, plan.PlanDate, plan.Content, plan.Status, plan.UpdatedAt, id,
	)
	if err != nil {
		log.Printf("更新失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "更新失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "更新成功"})
}

// deletePlan 删除计划  改成软删？
func deletePlan(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM plans WHERE id=?", id)
	if err != nil {
		log.Printf("删除失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败"})
		return
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "删除成功"})
}

// refreshTodayPlan 刷新今日计划（强制重新生成，忽略缓存）
func refreshTodayPlan(c *gin.Context) {
	today := time.Now().Format("2006-01-02")

	log.Printf("🔄 刷新今日计划（强制重新生成）...")

	// 1. 删除旧计划
	_, err := db.Exec("DELETE FROM plans WHERE plan_date=?", today)
	if err != nil {
		log.Printf("❌ 删除旧计划失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除旧计划失败: " + err.Error()})
		return
	}

	// 2. 构建Prompt
	promptBuilder := NewPromptBuilder(db)
	systemPrompt, userPrompt, err := promptBuilder.BuildPlanPrompt(today)
	if err != nil {
		log.Printf("❌ 构建Prompt失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "构建计划失败: " + err.Error()})
		return
	}

	// 3. 调用LLM生成新计划
	llmService := NewLLMService()
	planContent, err := llmService.Generate(systemPrompt, userPrompt)
	if err != nil {
		log.Printf("❌ LLM生成计划失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "生成计划失败: " + err.Error()})
		return
	}

	log.Printf("✅ LLM生成新计划成功")

	// 4. 保存新计划到数据库
	result, err := db.Exec(
		"INSERT INTO plans (goal_id, plan_date, content, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		1, today, planContent, "active", time.Now(), time.Now(),
	)
	if err != nil {
		log.Printf("❌ 保存计划失败: %v", err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "保存计划失败: " + err.Error()})
		return
	}

	lastID, _ := result.LastInsertId()
	newPlan := Plan{
		ID:        int(lastID),
		GoalID:    1,
		PlanDate:  today,
		Content:   planContent,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(http.StatusOK, Response{Code: 0, Message: "计划已刷新", Data: newPlan})
}
