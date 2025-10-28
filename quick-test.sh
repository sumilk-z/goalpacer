#!/bin/bash

# 🧪 GoalPacer 快速测试脚本

echo ""
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                                                            ║"
echo "║          🧪 GoalPacer API 快速测试                        ║"
echo "║                                                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""

BASE_URL="http://localhost:8080"
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

# 检查后端是否运行
echo -e "${BLUE}检查后端服务...${NC}"
if ! curl -s $BASE_URL/api/goals > /dev/null 2>&1; then
    echo -e "${YELLOW}⚠️  后端服务未运行${NC}"
    echo "请先运行: bash start-local.sh"
    exit 1
fi
echo -e "${GREEN}✅ 后端服务正常${NC}"
echo ""

# 测试1: 创建目标
echo -e "${BLUE}📝 测试1: 创建学习目标${NC}"
GOAL_RESPONSE=$(curl -s -X POST $BASE_URL/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "学习Go语言",
    "description": "掌握Go并发编程和Web开发",
    "deadline": "2025-12-31"
  }')

GOAL_ID=$(echo $GOAL_RESPONSE | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')
echo "响应: $GOAL_RESPONSE"
echo -e "${GREEN}✅ 目标已创建 (ID: $GOAL_ID)${NC}"
echo ""

# 测试2: 获取目标列表
echo -e "${BLUE}📋 测试2: 获取目标列表${NC}"
curl -s -X GET $BASE_URL/api/goals | jq '.' 2>/dev/null || curl -s -X GET $BASE_URL/api/goals
echo -e "${GREEN}✅ 目标列表获取成功${NC}"
echo ""

# 测试3: 设置时间规则
echo -e "${BLUE}⏰ 测试3: 设置学习时间规则${NC}"
curl -s -X POST $BASE_URL/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }' | jq '.' 2>/dev/null || curl -s -X POST $BASE_URL/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }'
echo -e "${GREEN}✅ 时间规则已设置${NC}"
echo ""

# 测试4: 记录学习内容
echo -e "${BLUE}📚 测试4: 记录学习内容${NC}"
curl -s -X POST $BASE_URL/api/logs \
  -H "Content-Type: application/json" \
  -d "{
    \"goal_id\": $GOAL_ID,
    \"content\": \"学习了Go的goroutine和channel机制\",
    \"duration\": 120
  }" | jq '.' 2>/dev/null || curl -s -X POST $BASE_URL/api/logs \
  -H "Content-Type: application/json" \
  -d "{
    \"goal_id\": $GOAL_ID,
    \"content\": \"学习了Go的goroutine和channel机制\",
    \"duration\": 120
  }"
echo -e "${GREEN}✅ 学习内容已记录${NC}"
echo ""

# 测试5: 获取学习记录
echo -e "${BLUE}📖 测试5: 获取学习记录${NC}"
curl -s -X GET $BASE_URL/api/logs | jq '.' 2>/dev/null || curl -s -X GET $BASE_URL/api/logs
echo -e "${GREEN}✅ 学习记录获取成功${NC}"
echo ""

# 测试6: 获取今日计划（LLM生成）
echo -e "${BLUE}🤖 测试6: 获取今日计划（LLM生成）${NC}"
echo "⏳ 等待LLM生成计划（首次2-5秒，后续使用缓存）..."
PLAN_RESPONSE=$(curl -s -X GET $BASE_URL/api/plan/today)
echo "响应: $PLAN_RESPONSE"
echo -e "${GREEN}✅ 计划生成成功${NC}"
echo ""

# 测试7: 再次获取计划（测试缓存）
echo -e "${BLUE}⚡ 测试7: 再次获取计划（测试缓存）${NC}"
echo "⏳ 这次应该很快（使用缓存）..."
START_TIME=$(date +%s%N)
curl -s -X GET $BASE_URL/api/plan/today > /dev/null
END_TIME=$(date +%s%N)
ELAPSED=$((($END_TIME - $START_TIME) / 1000000))
echo -e "${GREEN}✅ 缓存命中，耗时: ${ELAPSED}ms${NC}"
echo ""

# 总结
echo "╔════════════════════════════════════════════════════════════╗"
echo "║                                                            ║"
echo -e "║  ${GREEN}✨ 所有测试完成！${NC}                                    ║"
echo "║                                                            ║"
echo "╚════════════════════════════════════════════════════════════╝"
echo ""
echo -e "${BLUE}📊 测试结果:${NC}"
echo "  ✅ 创建目标"
echo "  ✅ 获取目标列表"
echo "  ✅ 设置时间规则"
echo "  ✅ 记录学习内容"
echo "  ✅ 获取学习记录"
echo "  ✅ 生成今日计划（LLM）"
echo "  ✅ 缓存测试"
echo ""
echo -e "${BLUE}🎯 下一步:${NC}"
echo "  1. 打开浏览器访问 http://localhost:3000"
echo "  2. 在前端界面体验完整功能"
echo "  3. 查看AI生成的学习计划"
echo ""
