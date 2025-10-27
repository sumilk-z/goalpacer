#!/bin/bash

# GoalPacer 后端启动脚本

echo "🚀 启动 GoalPacer 后端服务..."

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "❌ 未找到 Go，请先安装 Go 1.21+"
    exit 1
fi

# 下载依赖
echo "📦 下载依赖..."
go mod download

# 运行服务
echo "🔧 启动服务..."
go run main.go database.go models.go handlers.go
