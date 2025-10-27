#!/bin/bash
echo "========================================="
echo "启动 GoalPacer 后端服务"
echo "========================================="
cd /Users/zhucui/CodeBuddy/20251027000233/backend
echo "当前目录: $(pwd)"
echo "开始运行..."
go run main.go database.go models.go handlers.go
