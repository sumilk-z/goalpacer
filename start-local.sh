#!/bin/bash

# 🚀 GoalPacer 本地启动脚本

set -e

echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                                                            ║"
echo "║          🚀 GoalPacer 本地启动                            ║"
echo "║                                                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 检查依赖
echo -e "${BLUE}📋 检查依赖...${NC}"

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ 未找到Go环境${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Go环境检查通过${NC}"

if ! command -v node &> /dev/null; then
    echo -e "${RED}❌ 未找到Node.js环境${NC}"
    exit 1
fi
echo -e "${GREEN}✅ Node.js环境检查通过${NC}"

# 编译后端
echo ""
echo -e "${BLUE}🔨 编译后端...${NC}"
cd backend
go build -o goalpacer-backend 2>&1 | grep -v "^#" || true
if [ ! -f goalpacer-backend ]; then
    echo -e "${RED}❌ 后端编译失败${NC}"
    exit 1
fi
echo -e "${GREEN}✅ 后端编译成功${NC}"

# 启动后端
echo ""
echo -e "${BLUE}🚀 启动后端服务...${NC}"
./goalpacer-backend > /tmp/goalpacer-backend.log 2>&1 &
BACKEND_PID=$!
echo -e "${GREEN}✅ 后端已启动 (PID: $BACKEND_PID)${NC}"

# 等待后端启动
sleep 2

# 检查后端是否运行
if ! kill -0 $BACKEND_PID 2>/dev/null; then
    echo -e "${RED}❌ 后端启动失败${NC}"
    cat /tmp/goalpacer-backend.log
    exit 1
fi

# 启动前端
echo ""
echo -e "${BLUE}📦 安装前端依赖...${NC}"
cd ../frontend
npm install > /dev/null 2>&1 || npm install
echo -e "${GREEN}✅ 前端依赖安装完成${NC}"

echo ""
echo -e "${BLUE}🚀 启动前端服务...${NC}"
npm start > /tmp/goalpacer-frontend.log 2>&1 &
FRONTEND_PID=$!
echo -e "${GREEN}✅ 前端已启动 (PID: $FRONTEND_PID)${NC}"

# 等待前端启动
sleep 3

# 显示启动信息
echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                                                            ║"
echo -e "║  ${GREEN}✨ 服务已启动！${NC}                                      ║"
echo "║                                                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""
echo -e "${BLUE}📱 前端地址:${NC} ${GREEN}http://localhost:3000${NC}"
echo -e "${BLUE}🔌 后端地址:${NC} ${GREEN}http://localhost:8080${NC}"
echo ""
echo -e "${YELLOW}💡 提示:${NC}"
echo "  • 打开浏览器访问 http://localhost:3000"
echo "  • 创建学习目标"
echo "  • 设置学习时间"
echo "  • 记录学习内容"
echo "  • 查看AI生成的学习计划"
echo ""
echo -e "${YELLOW}📝 查看日志:${NC}"
echo "  • 后端日志: tail -f /tmp/goalpacer-backend.log"
echo "  • 前端日志: tail -f /tmp/goalpacer-frontend.log"
echo ""
echo -e "${YELLOW}⏹️  停止服务:${NC}"
echo "  • 按 Ctrl+C 停止所有服务"
echo ""

# 处理中断信号
cleanup() {
    echo ""
    echo -e "${YELLOW}⏹️  停止服务...${NC}"
    kill $BACKEND_PID 2>/dev/null || true
    kill $FRONTEND_PID 2>/dev/null || true
    sleep 1
    echo -e "${GREEN}✅ 服务已停止${NC}"
    exit 0
}

trap cleanup INT TERM

# 等待进程
wait
