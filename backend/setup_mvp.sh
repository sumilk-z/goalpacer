#!/bin/bash

# MVP 快速启动脚本

echo "🚀 GoalPacer MVP 快速启动"
echo "================================"

# 1. 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ 未找到Go环境，请先安装Go 1.21+"
    exit 1
fi

echo "✅ Go环境检查通过"

# 2. 下载依赖
echo "📦 下载依赖..."
go mod download

# 3. 检查.env文件
if [ ! -f .env ]; then
    echo "⚠️  未找到.env文件，复制.env.example"
    cp .env.example .env
    echo "📝 请编辑.env文件，设置LLM_API_KEY"
    echo "   或者使用模拟数据运行（不设置API_KEY）"
fi

# 4. 编译
echo "🔨 编译后端..."
go build -o goalpacer-backend

if [ $? -ne 0 ]; then
    echo "❌ 编译失败"
    exit 1
fi

echo "✅ 编译成功"

# 5. 启动
echo "🎯 启动后端服务..."
echo "📍 访问地址: http://localhost:8080"
echo "📍 API文档: http://localhost:8080/api"
echo ""
echo "按 Ctrl+C 停止服务"
echo ""

./goalpacer-backend
