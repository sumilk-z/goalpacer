#!/bin/bash

# GoalPacer 后端测试脚本

echo "🧪 GoalPacer 后端测试"
echo "================================"

cd /Users/zhucui/CodeBuddy/20251027000233/backend

# 下载依赖
echo "📦 下载依赖..."
go mod download

# 运行测试
echo ""
echo "🔧 运行测试用例..."
echo ""

go test -v -run TestCreateGoal
go test -v -run TestGetGoals
go test -v -run TestUpdateGoal
go test -v -run TestDeleteGoal

go test -v -run TestSetTimeRule
go test -v -run TestGetTimeRules

go test -v -run TestCreateLog
go test -v -run TestGetLogs

go test -v -run TestCreatePlan
go test -v -run TestGetTodayPlan
go test -v -run TestUpdatePlan
go test -v -run TestDeletePlan

echo ""
echo "================================"
echo "✅ 所有测试完成！"
echo "================================"
