# MVP 架构设计

## 🏗️ 系统架构图

```
┌─────────────────────────────────────────────────────────────┐
│                        前端 (React)                          │
│  ┌──────────────────────────────────────────────────────┐   │
│  │ TodayPlan.js                                         │   │
│  │ - 显示今日计划                                       │   │
│  │ - 调用 GET /api/plan/today                          │   │
│  └──────────────────────────────────────────────────────┘   │
└────────────────────────┬─────────────────────────────────────┘
                         │ HTTP
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    后端 (Go + Gin)                           │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  getTodayPlan 处理器                                        │
│  ┌──────────────────────────────────────────────────────┐   │
│  │ 1. 检查缓存                                          │   │
│  │    SELECT * FROM plans WHERE plan_date = TODAY      │   │
│  │    ├─ 有缓存 → 返回                                 │   │
│  │    └─ 无缓存 → 继续                                 │   │
│  │                                                      │   │
│  │ 2. 调用 PromptBuilder                               │   │
│  │    ├─ 获取目标                                      │   │
│  │    ├─ 获取时间规则                                  │   │
│  │    ├─ 获取学习记录                                  │   │
│  │    └─ 提取已学内容                                  │   │
│  │                                                      │   │
│  │ 3. 调用 LLMService                                  │   │
│  │    ├─ 构建Prompt                                    │   │
│  │    └─ 调用LLM API                                   │   │
│  │                                                      │   │
│  │ 4. 保存计划                                         │   │
│  │    INSERT INTO plans (...)                          │   │
│  │                                                      │   │
│  │ 5. 返回计划                                         │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                              │
│  PromptBuilder 模块                                         │
│  ┌──────────────────────────────────────────────────────┐   │
│  │ BuildPlanPrompt()                                    │   │
│  │ ├─ getGoals()          → 查询goals表                │   │
│  │ ├─ getTimeRules()      → 查询time_rules表           │   │
│  │ ├─ getRecentLogs()     → 查询learning_logs表        │   │
│  │ ├─ extractLearnedContent()  → MD5去重              │   │
│  │ └─ formatPrompt()      → 构建Prompt                │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                              │
│  LLMService 模块                                            │
│  ┌──────────────────────────────────────────────────────┐   │
│  │ Generate(systemPrompt, userPrompt)                   │   │
│  │ ├─ callClaude()   → Claude API                      │   │
│  │ ├─ callOpenAI()   → OpenAI API                      │   │
│  │ ├─ callGemini()   → Gemini API                      │   │
│  │ └─ generateMockPlan() → 模拟数据                    │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                              │
└────────────────┬──────────────────────────┬─────────────────┘
                 │                          │
                 ▼                          ▼
        ┌─────────────────┐        ┌──────────────────┐
        │  SQLite数据库   │        │  LLM API         │
        ├─────────────────┤        ├──────────────────┤
        │ goals           │        │ Claude           │
        │ time_rules      │        │ OpenAI           │
        │ learning_logs   │        │ Gemini           │
        │ plans           │        │ (或模拟数据)     │
        └─────────────────┘        └──────────────────┘
```

---

## 📊 数据流图

### 场景1：首次请求（无缓存）

```
用户请求
  │
  ▼
GET /api/plan/today
  │
  ▼
检查缓存 (SELECT * FROM plans WHERE plan_date = TODAY)
  │
  ├─ 缓存命中 ──→ 返回缓存 ──→ 响应 (< 100ms)
  │
  └─ 缓存未命中
      │
      ▼
    PromptBuilder.BuildPlanPrompt()
      │
      ├─ getGoals()
      │   └─ SELECT * FROM goals WHERE status = 'active'
      │
      ├─ getTimeRules()
      │   └─ SELECT * FROM time_rules WHERE day_of_week = TODAY
      │
      ├─ getRecentLogs()
      │   └─ SELECT * FROM learning_logs WHERE log_date >= DATE-7
      │
      └─ extractLearnedContent()
          └─ MD5哈希去重
      │
      ▼
    构建 systemPrompt 和 userPrompt
      │
      ▼
    LLMService.Generate()
      │
      ├─ 检查 LLM_API_KEY
      │   ├─ 有 → 调用真实LLM API
      │   └─ 无 → 返回模拟数据
      │
      ▼
    LLM API 调用 (2-5秒)
      │
      ├─ Claude API
      ├─ OpenAI API
      └─ Gemini API
      │
      ▼
    解析响应
      │
      ▼
    保存计划
      │
      └─ INSERT INTO plans (...)
      │
      ▼
    返回计划 (响应)
```

### 场景2：同日重复请求（有缓存）

```
用户请求
  │
  ▼
GET /api/plan/today
  │
  ▼
检查缓存 (SELECT * FROM plans WHERE plan_date = TODAY)
  │
  ├─ 缓存命中 ──→ 返回缓存 ──→ 响应 (< 100ms) ✅
  │
  └─ 缓存未命中 (不会发生)
```

---

## 🔄 模块交互图

```
┌──────────────────────────────────────────────────────────┐
│                    getTodayPlan()                        │
└────────────────────┬─────────────────────────────────────┘
                     │
        ┌────────────┼────────────┐
        │            │            │
        ▼            ▼            ▼
    ┌────────┐  ┌──────────────┐  ┌───────────┐
    │ 数据库 │  │PromptBuilder │  │LLMService │
    └────────┘  └──────────────┘  └───────────┘
        │            │                  │
        │            ├─→ 查询数据库     │
        │            │   ├─ goals       │
        │            │   ├─ time_rules  │
        │            │   └─ logs        │
        │            │                  │
        │            ├─→ 提取内容       │
        │            │   └─ MD5去重     │
        │            │                  │
        │            └─→ 构建Prompt     │
        │                               │
        │                ┌──────────────┤
        │                │              │
        │                ▼              ▼
        │            ┌─────────────────────┐
        │            │  LLM API 调用       │
        │            │  ├─ Claude          │
        │            │  ├─ OpenAI          │
        │            │  ├─ Gemini          │
        │            │  └─ Mock Data       │
        │            └─────────────────────┘
        │                    │
        │                    ▼
        │            ┌─────────────────────┐
        │            │  解析响应            │
        │            │  提取计划JSON        │
        │            └─────────────────────┘
        │                    │
        └────────────────────┼────────────────┐
                             │                │
                             ▼                ▼
                        ┌─────────────┐  ┌──────────┐
                        │ 保存计划    │  │ 返回响应 │
                        │ INSERT      │  │ JSON     │
                        └─────────────┘  └──────────┘
```

---

## 📦 文件结构

```
backend/
├── main.go                    # 入口点，路由定义
├── handlers.go                # API处理器
│   └─ getTodayPlan()         # ✨ 改造：LLM集成
├── models.go                  # 数据模型
├── database.go                # 数据库初始化
├── llm_service.go             # ✨ 新增：LLM服务
│   ├─ NewLLMService()
│   ├─ Generate()
│   ├─ callClaude()
│   ├─ callOpenAI()
│   ├─ callGemini()
│   └─ generateMockPlan()
├── prompt_builder.go          # ✨ 新增：Prompt构建
│   ├─ NewPromptBuilder()
│   ├─ BuildPlanPrompt()
│   ├─ getGoals()
│   ├─ getTimeRules()
│   ├─ getRecentLogs()
│   ├─ extractLearnedContent()
│   └─ formatXxx()
├── .env.example               # ✨ 新增：环境变量示例
├── setup_mvp.sh               # ✨ 新增：启动脚本
└── goalpacer.db               # SQLite数据库文件
```

---

## 🔌 API 接口

### getTodayPlan 接口

```
请求：
  GET /api/plan/today

响应：
  {
    "code": 0,
    "message": "success",
    "data": {
      "id": 1,
      "goal_id": 1,
      "plan_date": "2025-10-27",
      "content": "{\"date\":\"2025-10-27\",\"summary\":\"...\",\"tasks\":[...]}",
      "status": "active",
      "created_at": "2025-10-27T10:30:00Z",
      "updated_at": "2025-10-27T10:30:00Z"
    }
  }

content 字段是JSON字符串，包含：
  {
    "date": "2025-10-27",
    "summary": "今日计划总结",
    "tasks": [
      {
        "goal_id": 1,
        "title": "任务标题",
        "description": "详细描述",
        "duration_minutes": 60,
        "start_time": "09:00",
        "priority": "high"
      }
    ],
    "total_duration_minutes": 180,
    "notes": "额外说明"
  }
```

---

## 🗄️ 数据库表

### 现有表

```sql
-- 学习目标
CREATE TABLE goals (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  status TEXT,
  deadline DATE,
  created_at DATETIME,
  updated_at DATETIME
);

-- 时间规则
CREATE TABLE time_rules (
  id INTEGER PRIMARY KEY,
  day_of_week INTEGER,
  start_time TEXT,
  end_time TEXT,
  created_at DATETIME,
  updated_at DATETIME
);

-- 学习记录
CREATE TABLE learning_logs (
  id INTEGER PRIMARY KEY,
  goal_id INTEGER,
  content TEXT,
  duration INTEGER,
  log_date DATE,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (goal_id) REFERENCES goals(id)
);

-- 学习计划
CREATE TABLE plans (
  id INTEGER PRIMARY KEY,
  goal_id INTEGER,
  plan_date DATE,
  content TEXT,
  status TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (goal_id) REFERENCES goals(id)
);
```

---

## 🔐 环境变量

```bash
# LLM 配置
LLM_PROVIDER=claude              # 提供商：claude, openai, gemini
LLM_API_KEY=sk-ant-xxxxx         # API Key
LLM_MODEL=claude-3-5-sonnet-...  # 模型名称（可选）

# 数据库配置（可选）
DATABASE_URL=...                 # 数据库连接字符串
```

---

## ⚡ 性能优化

### 缓存策略

```
请求 /api/plan/today
  │
  ├─ 第1次：无缓存
  │   └─ 调用LLM (2-5秒)
  │       └─ 保存到数据库
  │
  ├─ 第2-N次（同日）：有缓存
  │   └─ 直接返回 (<100ms)
  │
  └─ 第二天：缓存过期
      └─ 调用LLM (2-5秒)
```

### 数据库查询优化

```
getGoals()
  └─ SELECT * FROM goals WHERE status='active'
     └─ 索引：status

getTimeRules()
  └─ SELECT * FROM time_rules WHERE day_of_week=?
     └─ 索引：day_of_week

getRecentLogs()
  └─ SELECT * FROM learning_logs WHERE log_date >= DATE-7
     └─ 索引：log_date
```

---

## 🧪 测试流程

```
1. 启动后端
   bash setup_mvp.sh

2. 创建测试数据
   POST /api/goals
   POST /api/time-rules
   POST /api/logs

3. 调用核心接口
   GET /api/plan/today
   ├─ 第1次：生成计划 (2-5秒)
   └─ 第2次：返回缓存 (<100ms)

4. 验证结果
   ├─ 计划格式正确
   ├─ 包含所有任务
   └─ 时间分配合理
```

---

## 📈 扩展点

### 后续可以扩展的地方

```
1. 向量数据库
   ├─ 替换MD5去重
   └─ 使用向量相似度检测

2. 多用户支持
   ├─ 添加user_id字段
   └─ 修改所有查询

3. 高级Prompt
   ├─ 学习风格分析
   ├─ 个性化推荐
   └─ 动态难度调整

4. 性能优化
   ├─ Redis缓存
   ├─ 数据库连接池
   └─ 异步处理

5. 监控和分析
   ├─ 成本追踪
   ├─ 性能监控
   └─ 用户行为分析
```

---

## 🎯 总结

MVP架构设计遵循以下原则：

1. **简洁性** - 最小化实现，快速验证
2. **可扩展性** - 为后续优化预留接口
3. **可维护性** - 清晰的模块划分
4. **灵活性** - 支持多个LLM提供商
5. **可靠性** - 完整的错误处理和日志

现在可以开始使用了！🚀
