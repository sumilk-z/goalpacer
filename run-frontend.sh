#!/bin/bash

# 启动前端服务

cd /Users/zhucui/CodeBuddy/goalpacer/frontend

echo "📦 安装依赖..."
npm install > /dev/null 2>&1

echo "🚀 启动前端服务..."
echo "📍 访问地址: http://localhost:3000"
echo ""
echo "按 Ctrl+C 停止服务"
echo ""

npm start
