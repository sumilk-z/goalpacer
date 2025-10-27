#!/bin/bash

# GoalPacer 前后端一键启动脚本

echo "🎯 GoalPacer 项目启动"
echo "================================"

# 启动后端
echo ""
echo "📍 步骤 1/2: 启动后端服务..."
cd /Users/zhucui/CodeBuddy/20251027000233/backend
go mod download > /dev/null 2>&1
go run main.go database.go models.go handlers.go &
BACKEND_PID=$!
echo "✅ 后端服务已启动 (PID: $BACKEND_PID)"

# 等待后端启动
sleep 3

# 启动前端
echo ""
echo "📍 步骤 2/2: 启动前端服务..."
cd /Users/zhucui/CodeBuddy/20251027000233/frontend

# 停止旧的前端进程
PORT=3000
PID=$(lsof -ti:$PORT)
if [ ! -z "$PID" ]; then
  kill -9 $PID 2>/dev/null
  sleep 1
fi

rm -rf node_modules/.cache 2>/dev/null
npm start &
FRONTEND_PID=$!
echo "✅ 前端服务已启动 (PID: $FRONTEND_PID)"

echo ""
echo "================================"
echo "🚀 所有服务已启动！"
echo ""
echo "📱 前端: http://localhost:3000"
echo "🔧 后端: http://localhost:8080"
echo ""
echo "按 Ctrl+C 停止所有服务"
echo "================================"

# 等待进程
wait
