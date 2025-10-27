package main

import "time"

// Goal 学习目标
type Goal struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Deadline    *string   `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TimeRule 时间规则
type TimeRule struct {
	ID        int       `json:"id"`
	DayOfWeek int       `json:"day_of_week" binding:"required,min=0,max=6"`
	StartTime string    `json:"start_time" binding:"required"`
	EndTime   string    `json:"end_time" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LearningLog 学习记录
type LearningLog struct {
	ID        int       `json:"id"`
	GoalID    int       `json:"goal_id" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Duration  *int      `json:"duration"`
	LogDate   string    `json:"log_date" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Plan 学习计划
type Plan struct {
	ID        int       `json:"id"`
	GoalID    int       `json:"goal_id" binding:"required"`
	PlanDate  string    `json:"plan_date" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
