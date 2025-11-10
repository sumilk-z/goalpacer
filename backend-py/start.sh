#!/bin/bash

# GoalPacer 后端启动脚本

echo "=== GoalPacer 后端启动脚本 ==="

# 检查Python是否安装
if ! command -v python3 &> /dev/null; then
    echo "Python3 未安装，正在检查python命令..."
    if ! command -v python &> /dev/null; then
        echo "Error: Python 未安装，请先安装Python 3.8+"
        exit 1
    else
        PYTHON_CMD="python"
        echo "使用python命令"
    fi
else
    PYTHON_CMD="python3"
    echo "使用python3命令"
fi

# 检查Python版本
PYTHON_VERSION=$($PYTHON_CMD --version 2>&1 | awk '{print $2}')
echo "当前Python版本: $PYTHON_VERSION"

# 创建虚拟环境
if [ ! -d ".venv" ]; then
    echo "创建虚拟环境..."
    $PYTHON_CMD -m venv .venv
    if [ $? -ne 0 ]; then
        echo "Error: 虚拟环境创建失败"
        exit 1
    fi
fi

# 激活虚拟环境
echo "激活虚拟环境..."
if [[ "$OSTYPE" == "msys"* ]] || [[ "$OSTYPE" == "win32" ]]; then
    # Windows环境
    source .venv/Scripts/activate
else
    # Unix-like环境
    source .venv/bin/activate
fi



# 安装依赖
echo "安装项目依赖..."
if [ -f "requirements.txt" ]; then
    pip install -r requirements.txt
    if [ $? -ne 0 ]; then
        echo "Error: 依赖安装失败，请检查requirements.txt和网络连接"
        exit 1
    fi
else
    echo "Warning: requirements.txt 文件不存在"
fi

# 检查.env文件是否存在
if [ ! -f ".env" ]; then
    echo "创建.env文件..."
    if [ -f ".env.example" ]; then
        cp .env.example .env
        echo "已从.env.example创建.env文件，请根据需要配置API密钥"
    else
        echo "Warning: .env.example 文件不存在，无法创建.env文件"
    fi
fi

# 数据库初始化将由FastAPI应用启动时自动完成
echo "数据库将在应用启动时自动初始化..."

# 启动应用
echo "启动GoalPacer后端服务..."
echo "服务将在 http://localhost:8080 启动"
echo "API文档地址: http://localhost:8080/docs"

# 直接运行uvicorn启动应用
uvicorn main:app --reload --host 0.0.0.0 --port 8080