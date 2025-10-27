#!/bin/bash

# 前后端集成测试脚本

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║         🚀 GoalPacer 前后端集成测试                           ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 测试计数
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# 测试函数
test_api() {
  local method=$1
  local endpoint=$2
  local data=$3
  local description=$4
  
  TOTAL_TESTS=$((TOTAL_TESTS + 1))
  
  echo -e "${BLUE}测试 $TOTAL_TESTS: $description${NC}"
  echo "  请求: $method $endpoint"
  
  if [ -z "$data" ]; then
    response=$(curl -s -X $method "http://localhost:8080/api$endpoint" \
      -H "Content-Type: application/json")
  else
    echo "  数据: $data"
    response=$(curl -s -X $method "http://localhost:8080/api$endpoint" \
      -H "Content-Type: application/json" \
      -d "$data")
  fi
  
  # 检查响应
  if echo "$response" | grep -q '"code":0'; then
    echo -e "  ${GREEN}✅ 通过${NC}"
    PASSED_TESTS=$((PASSED_TESTS + 1))
    echo "  响应: $response" | head -c 100
    echo ""
    echo ""
    return 0
  else
    echo -e "  ${RED}❌ 失败${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
    echo "  响应: $response"
    echo ""
    return 1
  fi
}

# 等待服务启动
echo -e "${YELLOW}⏳ 等待服务启动...${NC}"
sleep 2

# 检查后端服务
echo -e "${YELLOW}检查后端服务...${NC}"
if curl -s http://localhost:8080/ping > /dev/null 2>&1; then
  echo -e "${GREEN}✅ 后端服务已启动${NC}"
else
  echo -e "${RED}❌ 后端服务未启动${NC}"
  exit 1
fi

echo ""
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║                    📋 测试用例                                 ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# ========== 目标管理测试 ==========
echo -e "${BLUE}🎯 目标管理测试${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# 创建目标
GOAL_DATA='{"name":"算法刷题","description":"每天刷LeetCode题目","status":"active"}'
test_api "POST" "/goals" "$GOAL_DATA" "创建目标"

# 获取所有目标
test_api "GET" "/goals" "" "获取所有目标"

# 更新目标（假设ID为1）
UPDATE_DATA='{"name":"算法刷题-更新","description":"深入学习算法","status":"active"}'
test_api "PUT" "/goals/1" "$UPDATE_DATA" "更新目标"

echo ""

# ========== 学习记录测试 ==========
echo -e "${BLUE}📚 学习记录测试${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# 创建学习记录
LOG_DATA='{"goal_id":1,"content":"学习了二叉树的前序遍历","duration":90,"record_date":"2025-10-27"}'
test_api "POST" "/logs" "$LOG_DATA" "创建学习记录"

# 获取学习记录
test_api "GET" "/logs" "" "获取所有学习记录"

# 按目标筛选
test_api "GET" "/logs?goal_id=1" "" "按目标筛选学习记录"

echo ""

# ========== 时间规则测试 ==========
echo -e "${BLUE}⏰ 时间规则测试${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# 设置时间规则
RULE_DATA='{"goal_id":1,"start_time":"09:00","end_time":"11:00","days":"1,2,3,4,5"}'
test_api "POST" "/time-rules" "$RULE_DATA" "设置时间规则"

# 获取时间规则
test_api "GET" "/time-rules" "" "获取所有时间规则"

echo ""

# ========== 学习计划测试 ==========
echo -e "${BLUE}📅 学习计划测试${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# 创建计划
PLAN_DATA='{"goal_id":1,"plan_date":"2025-10-27","content":"完成LeetCode 5道题","status":"pending"}'
test_api "POST" "/plan" "$PLAN_DATA" "创建学习计划"

# 获取今日计划
test_api "GET" "/plan/today" "" "获取今日计划"

# 获取指定日期计划
test_api "GET" "/plan?date=2025-10-27" "" "获取指定日期计划"

echo ""

# ========== 测试总结 ==========
echo "╔════════════════════════════════════════════════════════════════╗"
echo "║                    📊 测试总结                                 ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""
echo -e "总测试数:  ${BLUE}$TOTAL_TESTS${NC}"
echo -e "通过数:    ${GREEN}$PASSED_TESTS${NC}"
echo -e "失败数:    ${RED}$FAILED_TESTS${NC}"

if [ $FAILED_TESTS -eq 0 ]; then
  echo -e "\n${GREEN}✅ 所有测试通过！${NC}"
  exit 0
else
  echo -e "\n${RED}❌ 有测试失败${NC}"
  exit 1
fi
