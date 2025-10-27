#!/bin/bash

# 自动化 API 测试脚本

set -e

BASE_URL="http://localhost:8080/api"
TIMESTAMP=$(date +%s%N)

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 计数器
TOTAL=0
PASSED=0
FAILED=0

# 测试函数
test_endpoint() {
  local method=$1
  local endpoint=$2
  local data=$3
  local description=$4
  
  TOTAL=$((TOTAL + 1))
  
  echo -e "\n${BLUE}[测试 $TOTAL] $description${NC}"
  echo "  方法: $method"
  echo "  端点: $endpoint"
  
  if [ -n "$data" ]; then
    echo "  数据: $data"
    response=$(curl -s -X $method "$BASE_URL$endpoint" \
      -H "Content-Type: application/json" \
      -d "$data")
  else
    response=$(curl -s -X $method "$BASE_URL$endpoint" \
      -H "Content-Type: application/json")
  fi
  
  # 检查响应
  if echo "$response" | grep -q '"code":0'; then
    echo -e "  ${GREEN}✅ 通过${NC}"
    PASSED=$((PASSED + 1))
    
    # 提取 ID 用于后续测试
    if echo "$response" | grep -q '"id"'; then
      ID=$(echo "$response" | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')
      echo "  ID: $ID"
      echo "$ID"
    fi
  else
    echo -e "  ${RED}❌ 失败${NC}"
    FAILED=$((FAILED + 1))
    echo "  响应: $response"
    return 1
  fi
}

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║         🚀 GoalPacer API 自动化测试                           ║"
echo "╚════════════════════════════════════════════════════════════════╝"

# 检查服务
echo -e "\n${YELLOW}检查后端服务...${NC}"
if ! curl -s http://localhost:8080/ping > /dev/null 2>&1; then
  echo -e "${RED}❌ 后端服务未启动${NC}"
  exit 1
fi
echo -e "${GREEN}✅ 后端服务已启动${NC}"

# ========== 目标管理 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}🎯 目标管理${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 创建目标 1
GOAL_DATA1="{\"name\":\"算法刷题_$TIMESTAMP\",\"description\":\"每天刷LeetCode题目\",\"status\":\"active\"}"
GOAL_ID1=$(test_endpoint "POST" "/goals" "$GOAL_DATA1" "创建目标 1")

# 创建目标 2
GOAL_DATA2="{\"name\":\"Golang学习_$TIMESTAMP\",\"description\":\"深入学习Go语言\",\"status\":\"active\"}"
GOAL_ID2=$(test_endpoint "POST" "/goals" "$GOAL_DATA2" "创建目标 2")

# 获取所有目标
test_endpoint "GET" "/goals" "" "获取所有目标"

# 更新目标
UPDATE_DATA="{\"name\":\"算法刷题-更新_$TIMESTAMP\",\"description\":\"深入学习算法\",\"status\":\"active\"}"
test_endpoint "PUT" "/goals/$GOAL_ID1" "$UPDATE_DATA" "更新目标"

# ========== 学习记录 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📚 学习记录${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 创建学习记录 1
LOG_DATA1="{\"goal_id\":$GOAL_ID1,\"content\":\"学习了二叉树的前序遍历\",\"duration\":90,\"record_date\":\"2025-10-27\"}"
LOG_ID1=$(test_endpoint "POST" "/logs" "$LOG_DATA1" "创建学习记录 1")

# 创建学习记录 2
LOG_DATA2="{\"goal_id\":$GOAL_ID2,\"content\":\"深入学习了channel的使用\",\"duration\":120,\"record_date\":\"2025-10-27\"}"
LOG_ID2=$(test_endpoint "POST" "/logs" "$LOG_DATA2" "创建学习记录 2")

# 获取所有学习记录
test_endpoint "GET" "/logs" "" "获取所有学习记录"

# 按目标筛选
test_endpoint "GET" "/logs?goal_id=$GOAL_ID1" "" "按目标筛选学习记录"

# 按日期筛选
test_endpoint "GET" "/logs?record_date=2025-10-27" "" "按日期筛选学习记录"

# ========== 时间规则 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}⏰ 时间规则${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 设置时间规则
RULE_DATA="{\"goal_id\":$GOAL_ID1,\"start_time\":\"09:00\",\"end_time\":\"11:00\",\"days\":\"1,2,3,4,5\"}"
RULE_ID=$(test_endpoint "POST" "/time-rules" "$RULE_DATA" "设置时间规则")

# 获取所有时间规则
test_endpoint "GET" "/time-rules" "" "获取所有时间规则"

# ========== 学习计划 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📅 学习计划${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 创建计划 1
PLAN_DATA1="{\"goal_id\":$GOAL_ID1,\"plan_date\":\"2025-10-27\",\"content\":\"完成LeetCode 5道题\",\"status\":\"pending\"}"
PLAN_ID1=$(test_endpoint "POST" "/plan" "$PLAN_DATA1" "创建学习计划 1")

# 创建计划 2
PLAN_DATA2="{\"goal_id\":$GOAL_ID2,\"plan_date\":\"2025-10-27\",\"content\":\"学习channel和select\",\"status\":\"pending\"}"
PLAN_ID2=$(test_endpoint "POST" "/plan" "$PLAN_DATA2" "创建学习计划 2")

# 获取今日计划
test_endpoint "GET" "/plan/today" "" "获取今日计划"

# 获取指定日期计划
test_endpoint "GET" "/plan?date=2025-10-27" "" "获取指定日期计划"

# 更新计划状态
UPDATE_PLAN="{\"goal_id\":$GOAL_ID1,\"plan_date\":\"2025-10-27\",\"content\":\"完成LeetCode 5道题\",\"status\":\"completed\"}"
test_endpoint "PUT" "/plan/$PLAN_ID1" "$UPDATE_PLAN" "更新计划状态"

# ========== 删除操作 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}🗑️  删除操作${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 删除计划
test_endpoint "DELETE" "/plan/$PLAN_ID2" "" "删除计划"

# 删除目标
test_endpoint "DELETE" "/goals/$GOAL_ID2" "" "删除目标"

# ========== 测试总结 ==========
echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📊 测试总结${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

echo ""
echo -e "总测试数:  ${BLUE}$TOTAL${NC}"
echo -e "通过数:    ${GREEN}$PASSED${NC}"
echo -e "失败数:    ${RED}$FAILED${NC}"

if [ $FAILED -eq 0 ]; then
  echo -e "\n${GREEN}✅ 所有测试通过！${NC}"
  exit 0
else
  echo -e "\n${RED}❌ 有 $FAILED 个测试失败${NC}"
  exit 1
fi
