package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化测试环境
func setupTestDB(t *testing.T) {
	if err := InitDB(); err != nil {
		t.Fatalf("初始化数据库失败: %v", err)
	}
}

// 清理测试环境
func teardownTestDB(t *testing.T) {
	if err := CloseDB(); err != nil {
		t.Fatalf("关闭数据库失败: %v", err)
	}
}

// 创建测试路由
func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/goals", getGoals)
		api.POST("/goals", createGoal)
		api.PUT("/goals/:id", updateGoal)
		api.DELETE("/goals/:id", deleteGoal)

		api.GET("/time-rules", getTimeRules)
		api.POST("/time-rules", setTimeRules)

		api.GET("/logs", getLogs)
		api.POST("/logs", createLog)

		api.GET("/plan/today", getTodayPlan)
		api.GET("/plan", getPlan)
		api.POST("/plan", createPlan)
		api.PUT("/plan/:id", updatePlan)
		api.DELETE("/plan/:id", deletePlan)
	}

	return r
}

// 辅助函数：创建目标
func createTestGoal(t *testing.T, router *gin.Engine, name string) int {
	goal := Goal{
		Name:        name + "_" + fmt.Sprintf("%d", time.Now().UnixNano()),
		Description: "测试目标",
	}
	body, _ := json.Marshal(goal)
	req, _ := http.NewRequest("POST", "/api/goals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("创建目标失败，状态码 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)
	dataBytes, _ := json.Marshal(resp.Data)
	var createdGoal Goal
	json.Unmarshal(dataBytes, &createdGoal)
	return createdGoal.ID
}

// ===== 目标管理测试 =====

func TestCreateGoal(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	goal := Goal{
		Name:        "算法刷题_" + fmt.Sprintf("%d", time.Now().UnixNano()),
		Description: "每天刷 LeetCode 题目",
	}

	body, _ := json.Marshal(goal)
	req, _ := http.NewRequest("POST", "/api/goals", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 创建目标成功")
}

func TestGetGoals(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先创建一个目标
	createTestGoal(t, router, "Golang学习")

	// 获取所有目标
	req, _ := http.NewRequest("GET", "/api/goals", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 获取目标成功")
}

func TestUpdateGoal(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 创建目标
	goalID := createTestGoal(t, router, "Agent开发")

	// 更新目标
	updatedGoal := Goal{
		Name:        "Agent开发_" + fmt.Sprintf("%d", time.Now().UnixNano()),
		Description: "深入学习 LangChain 和 LLM",
		Status:      "active",
	}
	body, _ := json.Marshal(updatedGoal)
	req, _ := http.NewRequest("PUT", "/api/goals/"+fmt.Sprintf("%d", goalID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	t.Logf("✅ 更新目标成功")
}

func TestDeleteGoal(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 创建目标
	goalID := createTestGoal(t, router, "考研英语")

	// 删除目标
	req, _ := http.NewRequest("DELETE", "/api/goals/"+fmt.Sprintf("%d", goalID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	t.Logf("✅ 删除目标成功")
}

// ===== 时间规则测试 =====

func TestSetTimeRule(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	rule := TimeRule{
		DayOfWeek: 1,
		StartTime: "09:00",
		EndTime:   "17:00",
	}

	body, _ := json.Marshal(rule)
	req, _ := http.NewRequest("POST", "/api/time-rules", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 设置时间规则成功")
}

func TestGetTimeRules(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先设置一个规则
	rule := TimeRule{
		DayOfWeek: 2,
		StartTime: "10:00",
		EndTime:   "18:00",
	}
	body, _ := json.Marshal(rule)
	req, _ := http.NewRequest("POST", "/api/time-rules", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 获取所有规则
	req, _ = http.NewRequest("GET", "/api/time-rules", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 获取时间规则成功")
}

// ===== 学习记录测试 =====

func TestCreateLog(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先创建一个目标
	goalID := createTestGoal(t, router, "测试目标")

	// 创建学习记录
	learningLog := LearningLog{
		GoalID:  goalID,
		Content: "学习了二叉树的前序遍历",
		Duration: func() *int { i := 90; return &i }(),
		LogDate: time.Now().Format("2006-01-02"),
	}

	body, _ := json.Marshal(learningLog)
	req, _ := http.NewRequest("POST", "/api/logs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 创建学习记录成功")
}

func TestGetLogs(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先创建目标和记录
	goalID := createTestGoal(t, router, "测试目标2")

	learningLog := LearningLog{
		GoalID:  goalID,
		Content: "学习了动态规划",
		Duration: func() *int { i := 120; return &i }(),
		LogDate: time.Now().Format("2006-01-02"),
	}
	logBody, _ := json.Marshal(learningLog)
	req, _ := http.NewRequest("POST", "/api/logs", bytes.NewBuffer(logBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 获取记录
	req, _ = http.NewRequest("GET", "/api/logs", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 获取学习记录成功")
}

// ===== 计划测试 =====

func TestCreatePlan(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先创建一个目标
	goalID := createTestGoal(t, router, "测试目标3")

	// 创建计划
	plan := Plan{
		GoalID:   goalID,
		PlanDate: time.Now().Format("2006-01-02"),
		Content:  "完成 LeetCode 5 道题",
	}

	body, _ := json.Marshal(plan)
	req, _ := http.NewRequest("POST", "/api/plan", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 创建计划成功")
}

func TestGetTodayPlan(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 先创建目标和计划
	goalID := createTestGoal(t, router, "测试目标4")

	plan := Plan{
		GoalID:   goalID,
		PlanDate: time.Now().Format("2006-01-02"),
		Content:  "完成 LeetCode 3 道题",
	}
	planBody, _ := json.Marshal(plan)
	req, _ := http.NewRequest("POST", "/api/plan", bytes.NewBuffer(planBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 获取今日计划
	req, _ = http.NewRequest("GET", "/api/plan/today", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	var resp Response
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp.Code != 0 {
		t.Errorf("期望响应码 0，得到 %d", resp.Code)
	}

	t.Logf("✅ 获取今日计划成功")
}

func TestUpdatePlan(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 创建目标和计划
	goalID := createTestGoal(t, router, "测试目标5")

	plan := Plan{
		GoalID:   goalID,
		PlanDate: time.Now().Format("2006-01-02"),
		Content:  "完成 LeetCode 2 道题",
	}
	planBody, _ := json.Marshal(plan)
	req, _ := http.NewRequest("POST", "/api/plan", bytes.NewBuffer(planBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var planResp Response
	json.Unmarshal(w.Body.Bytes(), &planResp)
	dataBytes, _ := json.Marshal(planResp.Data)
	var createdPlan Plan
	json.Unmarshal(dataBytes, &createdPlan)

	// 更新计划
	updatedPlan := Plan{
		GoalID:   goalID,
		PlanDate: time.Now().Format("2006-01-02"),
		Content:  "完成 LeetCode 5 道题",
		Status:   "completed",
	}
	body, _ := json.Marshal(updatedPlan)
	req, _ = http.NewRequest("PUT", "/api/plan/"+fmt.Sprintf("%d", createdPlan.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	t.Logf("✅ 更新计划成功")
}

func TestDeletePlan(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	router := setupRouter()

	// 创建目标和计划
	goalID := createTestGoal(t, router, "测试目标6")

	plan := Plan{
		GoalID:   goalID,
		PlanDate: time.Now().Format("2006-01-02"),
		Content:  "完成 LeetCode 1 道题",
	}
	planBody, _ := json.Marshal(plan)
	req, _ := http.NewRequest("POST", "/api/plan", bytes.NewBuffer(planBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var planResp Response
	json.Unmarshal(w.Body.Bytes(), &planResp)
	dataBytes, _ := json.Marshal(planResp.Data)
	var createdPlan Plan
	json.Unmarshal(dataBytes, &createdPlan)

	// 删除计划
	req, _ = http.NewRequest("DELETE", "/api/plan/"+fmt.Sprintf("%d", createdPlan.ID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("期望状态码 200，得到 %d", w.Code)
	}

	t.Logf("✅ 删除计划成功")
}
