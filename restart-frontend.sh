#!/bin/bash

# GoalPacer 前端服务重启脚本
# 用途：一键停止旧服务并启动新服务，确保代码更新生效

echo "🔄 正在重启 GoalPacer 前端服务..."

# 1. 查找并停止占用 3000 端口的进程
echo "📍 步骤 1/3: 停止旧服务..."
PORT=3000
PID=$(lsof -ti:$PORT)

if [ -z "$PID" ]; then
  echo "   ℹ️  端口 $PORT 未被占用"
else
  echo "   🛑 发现进程 $PID 占用端口 $PORT，正在停止..."
  kill -9 $PID
  sleep 2
  echo "   ✅ 旧服务已停止"
fi

# 2. 清理缓存（可选，但能解决很多热更新问题）
echo "📍 步骤 2/3: 清理缓存..."
cd /Users/zhucui/CodeBuddy/20251027000233/frontend
rm -rf node_modules/.cache
echo "   ✅ 缓存已清理"

# 3. 启动新服务
echo "📍 步骤 3/3: 启动新服务..."
echo "   🚀 正在启动开发服务器..."
npm start

# 注意：npm start 会阻塞终端，按 Ctrl+C 可停止服务
