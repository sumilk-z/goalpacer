# 🚀 GoalPacer 本地启动指南

## ⚡ 快速启动（1分钟）

### 方式1：一键启动（推荐）

```bash
# 在项目根目录运行
bash start-local.sh
```

然后打开浏览器访问：**http://localhost:3000**

### 方式2：手动启动

#### 终端1：启动后端

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

#### 终端2：启动前端

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

---

## 🧪 快速测试

### 自动测试（推荐）

```bash
bash quick-test.sh
```

这会自动测试所有API接口，包括：
- ✅ 创建学习目标
- ✅ 获取目标列表
- ✅ 设置学习时间
- ✅ 记录学习内容
- ✅ 获取学习记录
- ✅ **生成今日计划（LLM）** ⭐
- ✅ 缓存测试

### 手动测试

#### 1️⃣ 创建学习目标

```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "学习Go语言",
    "description": "掌握Go并发编程",
    "deadline": "2025-12-31"
  }'
```

#### 2️⃣ 设置学习时间

```bash
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }'
```

#### 3️⃣ 记录学习内容

```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了goroutine和channel",
    "duration": 120
  }'
```

#### 4️⃣ 获取今日计划（LLM生成）⭐

```bash
curl -X GET http://localhost:8080/api/plan/today
```

---

## 🎯 前端功能体验

### 1. 创建学习目标
- 点击"新建目标"
- 输入目标名称和描述
- 设置截止日期
- 点击"创建"

### 2. 设置学习时间
- 点击"时间规则"
- 选择星期几
- 设置开始和结束时间
- 点击"保存"

### 3. 记录学习内容
- 点击"记录学习"
- 选择学习目标
- 输入学习内容
- 输入学习时长（分钟）
- 点击"保存"

### 4. 查看学习计划 ⭐
- 点击"今日计划"
- 系统会自动生成个性化学习计划
- 计划基于：
  - 你的学习目标
  - 你的学习时间规则
  - 你的历史学习记录
  - AI智能推荐

---

## 🔧 环境配置

### 使用模拟数据（推荐快速测试）

默认情况下，系统会使用模拟数据生成计划，无需配置API Key。

### 使用真实LLM API

#### 选项1：Claude API（推荐）

1. 获取API Key：https://console.anthropic.com
2. 设置环境变量：

```bash
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx
export LLM_MODEL=claude-3-5-sonnet-20241022
```

3. 启动后端：

```bash
cd backend
./goalpacer-backend
```

#### 选项2：OpenAI API

1. 获取API Key：https://platform.openai.com
2. 设置环境变量：

```bash
export LLM_PROVIDER=openai
export LLM_API_KEY=sk-xxxxxxxxxxxxx
export LLM_MODEL=gpt-4-turbo
```

#### 选项3：Gemini API

1. 获取API Key：https://ai.google.dev
2. 设置环境变量：

```bash
export LLM_PROVIDER=gemini
export LLM_API_KEY=xxxxxxxxxxxxx
export LLM_MODEL=gemini-1.5-flash
```

---

## 📊 系统架构

```
┌─────────────────────────────────────────────────────────┐
│                    前端 (React)                         │
│              http://localhost:3000                      │
└────────────────────┬────────────────────────────────────┘
                     │ HTTP/REST
┌────────────────────▼────────────────────────────────────┐
│                    后端 (Go)                            │
│              http://localhost:8080                      │
├─────────────────────────────────────────────────────────┤
│  • 目标管理 (Goals)                                     │
│  • 时间规则 (Time Rules)                               │
│  • 学习记录 (Learning Logs)                            │
│  • 计划生成 (Plan Generation)                          │
│  • LLM集成 (Claude/OpenAI/Gemini)                      │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  数据库 (SQLite)                        │
│              goalpacer.db                               │
└─────────────────────────────────────────────────────────┘
```

---

## 🐛 常见问题

### Q1: 后端启动失败

**错误：** `database initialization failed`

**解决：**
```bash
# 删除旧数据库
rm backend/goalpacer.db

# 重新启动
cd backend
./goalpacer-backend
```

### Q2: 前端无法连接后端

**错误：** `Failed to fetch from http://localhost:8080`

**解决：**
1. 确保后端已启动（检查 http://localhost:8080/api/goals）
2. 检查防火墙设置
3. 检查CORS配置

### Q3: LLM API调用失败

**错误：** `LLM generation failed`

**解决：**
1. 检查API Key是否正确
2. 检查网络连接
3. 使用模拟数据测试（不设置API Key）

### Q4: 计划生成很慢

**原因：** 首次调用需要调用LLM API（2-5秒）

**解决：** 第二次调用会使用缓存（<100ms）

### Q5: npm install 失败

**解决：**
```bash
# 清除npm缓存
npm cache clean --force

# 重新安装
npm install
```

---

## 📈 性能指标

| 操作 | 耗时 |
|------|------|
| 创建目标 | <50ms |
| 获取目标列表 | <50ms |
| 记录学习内容 | <50ms |
| 生成计划（首次） | 2-5秒 |
| 生成计划（缓存） | <100ms |

---

## 📚 详细文档

- 📖 [本地启动详细指南](./LOCAL_STARTUP_GUIDE.md)
- 📖 [MVP实现指南](./MVP_GUIDE.md)
- 📖 [架构设计](./MVP_ARCHITECTURE.md)
- 📖 [实现总结](./MVP_IMPLEMENTATION_SUMMARY.md)
- 📖 [快速参考](./QUICK_MVP_START.md)

---

## 🎬 完整体验流程

### 第1步：启动服务（1分钟）

```bash
bash start-local.sh
```

### 第2步：打开浏览器（1分钟）

访问 http://localhost:3000

### 第3步：创建学习目标（2分钟）

- 点击"新建目标"
- 输入"学习Go语言"
- 点击"创建"

### 第4步：设置学习时间（2分钟）

- 点击"时间规则"
- 设置周一 09:00-12:00
- 点击"保存"

### 第5步：记录学习内容（2分钟）

- 点击"记录学习"
- 输入"学习了goroutine"
- 输入"120"分钟
- 点击"保存"

### 第6步：查看AI生成的计划（1分钟）⭐

- 点击"今日计划"
- 查看系统生成的个性化学习计划
- 计划基于你的目标、时间和历史记录

**总耗时：约10分钟，完整体验MVP功能！**

---

## ✨ 核心特性

### 🤖 AI智能计划生成

系统使用LLM（Claude/OpenAI/Gemini）根据以下信息生成个性化学习计划：

- 📚 你的学习目标
- ⏰ 你的学习时间规则
- 📖 你的历史学习记录
- 🧠 AI智能推荐

### ⚡ 智能缓存

- 首次生成：2-5秒（调用LLM）
- 后续查询：<100ms（使用缓存）
- 缓存时间：24小时

### 🔄 内容去重

系统使用MD5哈希自动去重，避免重复的学习任务。

### 🔌 多LLM支持

支持多个LLM提供商：
- Claude（推荐）
- OpenAI
- Gemini
- 模拟数据（快速测试）

---

## 🚀 下一步

1. ✅ 启动本地服务
2. ✅ 体验前端功能
3. ✅ 测试API接口
4. ✅ 查看LLM生成的计划
5. 📝 根据反馈优化功能
6. 🚀 部署到云环境

---

## 📞 需要帮助？

- 查看详细文档：[LOCAL_STARTUP_GUIDE.md](./LOCAL_STARTUP_GUIDE.md)
- 查看API文档：[MVP_GUIDE.md](./MVP_GUIDE.md)
- 查看架构设计：[MVP_ARCHITECTURE.md](./MVP_ARCHITECTURE.md)

---

**现在就开始吧！** 🚀

```bash
bash start-local.sh
```

然后打开浏览器访问：**http://localhost:3000**
