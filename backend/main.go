package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	if err := InitDB(); err != nil {
		log.Fatalf("❌ 数据库初始化失败: %v", err)
	}
	defer CloseDB()

	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// --- 路由定义 ---
	api := r.Group("/api")
	{
		// 目标管理
		api.GET("/goals", getGoals)
		api.POST("/goals", createGoal)
		api.PUT("/goals/:id", updateGoal)
		api.DELETE("/goals/:id", deleteGoal)

		// 时间规则
		api.GET("/time-rules", getTimeRules)
		api.POST("/time-rules", setTimeRules)

		// 学习记录
		api.GET("/logs", getLogs)
		api.POST("/logs", createLog)

		// 计划
		api.GET("/plan/today", getTodayPlan)
		api.GET("/plan", getPlan)
		api.POST("/plan", createPlan)
		api.PUT("/plan/:id", updatePlan)
		api.DELETE("/plan/:id", deletePlan)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("🚀 GoalPacer 后端服务启动在 http://0.0.0.0:8080")
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
