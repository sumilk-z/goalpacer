# 🚀 本地启动指南

## 快速启动（3步）

### 1️⃣ 启动后端服务

```bash
cd backend
go build -o goalpacer-backend
./goalpacer-backend
```

**预期输出：**
```
✅ 数据库初始化成功
🎯 后端服务启动成功
📍 监听地址: http://localhost:8080
```

### 2️⃣ 启动前端服务（新终端）

```bash
cd frontend
npm install
npm start
```

**预期输出：**
```
✅ 前端服务启动成功
📍 访问地址: http://localhost:3000
```

### 3️⃣ 打开浏览器

访问 `http://localhost:3000`

---

## 📝 快速测试

### 测试1：创建学习目标

```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "学习Go语言",
    "description": "掌握Go并发编程",
    "deadline": "2025-12-31"
  }'
```

**预期响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "name": "学习Go语言",
    "description": "掌握Go并发编程",
    "status": "active",
    "deadline": "2025-12-31"
  }
}
```

### 测试2：设置学习时间规则

```bash
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }'
```

### 测试3：记录学习内容

```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了Go的goroutine和channel",
    "duration": 120
  }'
```

### 测试4：获取今日计划（LLM生成）⭐

```bash
curl -X GET http://localhost:8080/api/plan/today
```

**预期响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "plan_date": "2025-10-27",
    "content": "今日学习计划...",
    "status": "active"
  }
}
```

---

## 🔧 环境配置

### 使用模拟数据（推荐快速测试）

```bash
# 不设置 LLM_API_KEY，使用模拟数据
cd backend
./goalpacer-backend
```

### 使用真实LLM API

#### 选项1：Claude API（推荐）

```bash
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx
export LLM_MODEL=claude-3-5-sonnet-20241022
```

获取API Key：https://console.anthropic.com

#### 选项2：OpenAI API

```bash
export LLM_PROVIDER=openai
export LLM_API_KEY=sk-xxxxxxxxxxxxx
export LLM_MODEL=gpt-4-turbo
```

获取API Key：https://platform.openai.com

#### 选项3：Gemini API

```bash
export LLM_PROVIDER=gemini
export LLM_API_KEY=xxxxxxxxxxxxx
export LLM_MODEL=gemini-1.5-flash
```

获取API Key：https://ai.google.dev

---

## 📊 完整的API测试流程

### 步骤1：创建目标

```bash
# 创建目标1
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{"name": "学习Go", "description": "掌握Go编程", "deadline": "2025-12-31"}'

# 创建目标2
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{"name": "学习React", "description": "掌握React框架", "deadline": "2025-12-31"}'
```

### 步骤2：设置时间规则

```bash
# 周一 09:00-12:00
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{"day_of_week": 1, "start_time": "09:00", "end_time": "12:00"}'

# 周一 14:00-17:00
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{"day_of_week": 1, "start_time": "14:00", "end_time": "17:00"}'
```

### 步骤3：记录学习内容

```bash
# 记录学习1
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{"goal_id": 1, "content": "学习了goroutine", "duration": 60}'

# 记录学习2
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{"goal_id": 1, "content": "学习了channel", "duration": 90}'

# 记录学习3
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{"goal_id": 2, "content": "学习了React Hooks", "duration": 120}'
```

### 步骤4：获取今日计划

```bash
# 首次调用（会调用LLM，2-5秒）
curl -X GET http://localhost:8080/api/plan/today

# 第二次调用（使用缓存，<100ms）
curl -X GET http://localhost:8080/api/plan/today
```

---

## 🎯 前端功能体验

### 1. 目标管理
- ✅ 创建学习目标
- ✅ 查看目标列表
- ✅ 编辑目标
- ✅ 删除目标

### 2. 时间规则
- ✅ 设置每周学习时间
- ✅ 查看时间规则
- ✅ 编辑时间规则

### 3. 学习记录
- ✅ 记录学习内容
- ✅ 查看学习历史
- ✅ 统计学习时间

### 4. 智能计划 ⭐
- ✅ 自动生成今日计划
- ✅ 基于LLM的个性化推荐
- ✅ 智能缓存（24小时）
- ✅ 内容去重

---

## 🐛 常见问题

### Q1: 后端启动失败

**错误：** `database initialization failed`

**解决：**
```bash
# 删除旧数据库
rm goalpacer.db

# 重新启动
./goalpacer-backend
```

### Q2: 前端无法连接后端

**错误：** `Failed to fetch from http://localhost:8080`

**解决：**
1. 确保后端已启动
2. 检查后端是否监听在 8080 端口
3. 检查CORS配置

### Q3: LLM API调用失败

**错误：** `LLM generation failed`

**解决：**
1. 检查API Key是否正确
2. 检查网络连接
3. 使用模拟数据测试

### Q4: 计划生成很慢

**原因：** 首次调用需要调用LLM API（2-5秒）

**解决：** 第二次调用会使用缓存（<100ms）

---

## 📈 性能监控

### 查看后端日志

后端会输出详细的日志：

```
✅ 数据库初始化成功
🎯 后端服务启动成功
📍 监听地址: http://localhost:8080

[请求日志]
GET /api/goals
POST /api/logs
GET /api/plan/today
  ├─ 缓存未命中
  ├─ 构建Prompt
  ├─ 调用LLM
  ├─ 保存计划
  └─ 耗时: 3.2秒
```

### 查看前端日志

打开浏览器开发者工具（F12）查看：
- Network 标签：查看API请求
- Console 标签：查看错误信息
- Performance 标签：查看性能指标

---

## 🚀 一键启动脚本

创建 `start-all.sh`：

```bash
#!/bin/bash

# 启动后端
cd backend
go build -o goalpacer-backend
./goalpacer-backend &
BACKEND_PID=$!

# 启动前端
cd ../frontend
npm install
npm start &
FRONTEND_PID=$!

echo "✅ 服务已启动"
echo "📱 前端: http://localhost:3000"
echo "🔌 后端: http://localhost:8080"

# 等待中断
trap "kill $BACKEND_PID $FRONTEND_PID; exit 0" INT
wait
```

使用：
```bash
bash start-all.sh
```

---

## 📚 相关文档

- 📖 [MVP实现指南](./MVP_GUIDE.md)
- 📖 [架构设计](./MVP_ARCHITECTURE.md)
- 📖 [快速开始](./QUICK_MVP_START.md)
- 📖 [实现总结](./MVP_IMPLEMENTATION_SUMMARY.md)

---

## ✨ 下一步

1. ✅ 启动本地服务
2. ✅ 体验前端功能
3. ✅ 测试API接口
4. ✅ 查看LLM生成的计划
5. 📝 根据反馈优化功能

**现在就开始吧！** 🚀
