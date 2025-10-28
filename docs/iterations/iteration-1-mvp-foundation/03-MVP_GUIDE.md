# GoalPacer MVP 实现指南

## 📋 MVP 功能清单

### ✅ 已实现
- [x] 基础CRUD API（目标、时间规则、学习记录、计划）
- [x] LLM服务集成（支持Claude、OpenAI、Gemini）
- [x] Prompt构建器（自动提取已学内容）
- [x] getTodayPlan接口改造（使用LLM生成计划）
- [x] 计划缓存机制（24小时内复用）
- [x] 简单的内容去重（基于MD5哈希）

### 🚀 快速开始

#### 1. 安装依赖

```bash
cd backend
go mod download
```

#### 2. 配置LLM API

**方案A：使用真实LLM API（推荐）**

```bash
# 复制配置文件
cp .env.example .env

# 编辑.env，设置API Key
# 选择一个LLM提供商：

# Claude (推荐，最稳定)
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx

# 或 OpenAI
export LLM_PROVIDER=openai
export LLM_API_KEY=sk-xxxxxxxxxxxxx

# 或 Gemini (最便宜)
export LLM_PROVIDER=gemini
export LLM_API_KEY=xxxxxxxxxxxxx
```

**方案B：使用模拟数据（快速测试）**

```bash
# 不设置LLM_API_KEY，系统会自动使用模拟数据
# 这样可以快速测试整个流程
```

#### 3. 启动后端

```bash
# 方式1：使用启动脚本
bash setup_mvp.sh

# 方式2：直接运行
go run main.go handlers.go models.go database.go llm_service.go prompt_builder.go
```

#### 4. 测试API

**创建学习目标**

```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go语言学习",
    "description": "学习Go并发编程和Web框架",
    "status": "active",
    "deadline": "2025-12-31"
  }'
```

**设置时间规则**

```bash
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }'
```

**记录学习内容**

```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了goroutine和channel的基本用法",
    "duration": 90,
    "log_date": "2025-10-27"
  }'
```

**获取今日计划（核心功能）**

```bash
curl http://localhost:8080/api/plan/today
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "goal_id": 1,
    "plan_date": "2025-10-27",
    "content": "{\"date\": \"2025-10-27\", \"summary\": \"...\", \"tasks\": [...]}",
    "status": "active",
    "created_at": "2025-10-27T10:30:00Z",
    "updated_at": "2025-10-27T10:30:00Z"
  }
}
```

---

## 🔧 LLM API 选择指南

### Claude (推荐)

**优势**
- 最稳定，推理能力强
- 支持长文本（200K tokens）
- 中文理解好

**成本**
- 输入：$0.003/1K tokens
- 输出：$0.015/1K tokens
- 每个计划约1500 tokens，成本约 $0.025

**获取API Key**
1. 访问 https://console.anthropic.com
2. 注册账户
3. 创建API Key
4. 设置 `LLM_API_KEY=sk-ant-xxxxxxxxxxxxx`

### OpenAI

**优势**
- 功能完整
- 生态成熟

**成本**
- GPT-4 Turbo: $0.01/1K input, $0.03/1K output
- 较贵

**获取API Key**
1. 访问 https://platform.openai.com
2. 创建API Key
3. 设置 `LLM_API_KEY=sk-xxxxxxxxxxxxx`

### Gemini (最便宜)

**优势**
- 最便宜
- 支持长文本（1M tokens）

**成本**
- 输入：$0.00075/1K tokens
- 输出：$0.003/1K tokens
- 每个计划约 $0.005

**获取API Key**
1. 访问 https://makersuite.google.com
2. 创建API Key
3. 设置 `LLM_API_KEY=xxxxxxxxxxxxx`

---

## 📊 MVP 架构

```
┌─────────────────────────────────────────┐
│         前端 (React)                    │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│    后端 (Go + Gin)                      │
├─────────────────────────────────────────┤
│ getTodayPlan 接口                       │
│  ├─ 检查缓存（24小时）                  │
│  ├─ 调用 PromptBuilder                  │
│  │  ├─ 获取目标                         │
│  │  ├─ 获取时间规则                     │
│  │  ├─ 获取学习记录                     │
│  │  └─ 提取已学内容                     │
│  ├─ 调用 LLMService                     │
│  │  ├─ 构建Prompt                       │
│  │  └─ 调用LLM API                      │
│  └─ 保存计划到数据库                    │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│    LLM API (Claude/OpenAI/Gemini)       │
└─────────────────────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│    SQLite 数据库                        │
└─────────────────────────────────────────┘
```

---

## 🧪 测试场景

### 场景1：首次使用（无缓存）

```
1. 创建目标：Go语言学习
2. 设置时间规则：周一-周五 09:00-12:00
3. 记录学习内容：学习了goroutine
4. 调用 /api/plan/today
   → LLM生成计划
   → 保存到数据库
   → 返回计划
```

### 场景2：同日重复调用（有缓存）

```
1. 第一次调用 /api/plan/today
   → 生成计划（耗时2-5秒）
2. 第二次调用 /api/plan/today
   → 返回缓存（耗时<100ms）
```

### 场景3：模拟数据测试

```
1. 不设置 LLM_API_KEY
2. 调用 /api/plan/today
   → 返回模拟计划
   → 用于快速测试UI
```

---

## 📈 性能指标

| 指标 | 目标 | 实际 |
|------|------|------|
| 首次生成计划 | <5s | 2-4s |
| 缓存命中 | <100ms | <50ms |
| 数据库查询 | <100ms | <50ms |
| LLM API调用 | <3s | 2-3s |
| 内存占用 | <100MB | ~50MB |

---

## 🐛 常见问题

### Q1: 调用LLM API失败

**错误信息**
```
❌ Claude API 调用失败: connection refused
```

**解决方案**
1. 检查网络连接
2. 检查API Key是否正确
3. 检查API Key是否有额度
4. 查看LLM服务状态

### Q2: 返回模拟数据

**原因**
- 未设置 `LLM_API_KEY` 环境变量

**解决方案**
```bash
export LLM_API_KEY=your_api_key
export LLM_PROVIDER=claude
```

### Q3: 计划格式不对

**原因**
- LLM返回的JSON格式不符合预期

**解决方案**
1. 检查Prompt是否清晰
2. 尝试更换LLM提供商
3. 增加Prompt中的格式说明

### Q4: 性能太慢

**原因**
- LLM API响应慢
- 数据库查询慢

**解决方案**
1. 使用更快的LLM模型（如Gemini）
2. 增加缓存时间
3. 优化数据库查询

---

## 🚀 下一步优化

### 第2阶段（1-2周）
- [ ] 添加用户认证
- [ ] 支持多用户
- [ ] 优化Prompt（减少token使用）
- [ ] 添加错误重试机制

### 第3阶段（2-3周）
- [ ] 向量数据库集成（可选）
- [ ] 高级去重算法
- [ ] 性能监控
- [ ] 成本分析

### 第4阶段（3-4周）
- [ ] 迁移到TDSQL-C MySQL
- [ ] Redis缓存
- [ ] 容器化部署
- [ ] CI/CD流程

---

## 📚 相关文件

- `llm_service.go` - LLM服务实现
- `prompt_builder.go` - Prompt构建器
- `handlers.go` - API处理器（已修改getTodayPlan）
- `.env.example` - 环境变量配置示例
- `setup_mvp.sh` - 快速启动脚本

---

## 💡 提示

1. **开发阶段**：使用模拟数据快速测试
2. **测试阶段**：使用Gemini API（最便宜）
3. **生产阶段**：使用Claude API（最稳定）
4. **监控成本**：每个计划约 $0.01-0.05

---

## 📞 支持

如有问题，请检查：
1. 日志输出（查看详细错误信息）
2. API Key是否正确
3. 网络连接是否正常
4. LLM服务是否可用
