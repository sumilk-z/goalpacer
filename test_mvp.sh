#!/bin/bash

# MVP 测试脚本

BASE_URL="http://localhost:8080/api"

echo "🧪 GoalPacer MVP 测试"
echo "================================"
echo ""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 测试计数
TOTAL=0
PASSED=0
FAILED=0

# 测试函数
test_api() {
    local name=$1
    local method=$2
    local endpoint=$3
    local data=$4
    
    TOTAL=$((TOTAL + 1))
    echo -n "测试 $TOTAL: $name ... "
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s -X GET "$BASE_URL$endpoint")
    else
        response=$(curl -s -X $method "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            -d "$data")
    fi
    
    # 检查响应是否包含 "code"
    if echo "$response" | grep -q '"code"'; then
        echo -e "${GREEN}✅ 通过${NC}"
        PASSED=$((PASSED + 1))
        echo "  响应: $(echo $response | head -c 100)..."
    else
        echo -e "${RED}❌ 失败${NC}"
        FAILED=$((FAILED + 1))
        echo "  响应: $response"
    fi
    echo ""
}

# 1. 创建目标
test_api "创建目标" "POST" "/goals" '{
    "name": "Go语言学习",
    "description": "学习Go并发编程和Web框架",
    "status": "active",
    "deadline": "2025-12-31"
}'

# 2. 获取目标
test_api "获取目标列表" "GET" "/goals" ""

# 3. 设置时间规则
test_api "设置时间规则" "POST" "/time-rules" '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
}'

# 4. 获取时间规则
test_api "获取时间规则" "GET" "/time-rules" ""

# 5. 记录学习内容
test_api "记录学习内容" "POST" "/logs" '{
    "goal_id": 1,
    "content": "学习了goroutine和channel的基本用法",
    "duration": 90,
    "log_date": "2025-10-27"
}'

# 6. 获取学习记录
test_api "获取学习记录" "GET" "/logs" ""

# 7. 获取今日计划（核心功能）
echo -e "${YELLOW}⏳ 调用LLM生成计划，请稍候...${NC}"
test_api "获取今日计划（LLM生成）" "GET" "/plan/today" ""

# 8. 再次获取今日计划（测试缓存）
echo -e "${YELLOW}⏳ 测试缓存...${NC}"
test_api "获取今日计划（缓存）" "GET" "/plan/today" ""

# 9. 获取指定日期计划
test_api "获取指定日期计划" "GET" "/plan?date=2025-10-27" ""

echo ""
echo "================================"
echo -e "测试结果: ${GREEN}通过 $PASSED${NC} / ${RED}失败 $FAILED${NC} / 总计 $TOTAL"
echo ""

if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✅ 所有测试通过！${NC}"
    exit 0
else
    echo -e "${RED}❌ 有测试失败${NC}"
    exit 1
fi
