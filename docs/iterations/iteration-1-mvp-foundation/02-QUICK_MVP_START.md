# 🚀 MVP 快速启动（5分钟）

## 1️⃣ 获取API Key（选一个）

### 选项A：Claude（推荐）
```bash
# 访问 https://console.anthropic.com
# 创建API Key，复制下来
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx
```

### 选项B：Gemini（最便宜）
```bash
# 访问 https://makersuite.google.com
# 创建API Key，复制下来
export LLM_PROVIDER=gemini
export LLM_API_KEY=xxxxxxxxxxxxx
```

### 选项C：不用API Key（快速测试）
```bash
# 直接运行，使用模拟数据
# 不需要设置任何环境变量
```

---

## 2️⃣ 启动后端

```bash
cd backend
bash setup_mvp.sh
```

或者：

```bash
cd backend
go run main.go handlers.go models.go database.go llm_service.go prompt_builder.go
```

看到这个输出说明成功：
```
✅ 数据库连接成功
✅ 数据表创建成功
🚀 GoalPacer 后端服务启动在 http://0.0.0.0:8080
```

---

## 3️⃣ 测试API

### 新开一个终端，运行测试脚本

```bash
bash test_mvp.sh
```

或者手动测试：

```bash
# 创建目标
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{"name":"Go学习","description":"学习Go","status":"active"}'

# 记录学习
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{"goal_id":1,"content":"学习了goroutine","duration":90,"log_date":"2025-10-27"}'

# 获取今日计划（核心功能）
curl http://localhost:8080/api/plan/today
```

---

## 4️⃣ 查看结果

### 如果设置了API Key

返回LLM生成的计划：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "content": "{\"date\":\"2025-10-27\",\"summary\":\"...\",\"tasks\":[...]}"
  }
}
```

### 如果没有设置API Key

返回模拟计划：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "content": "{\"date\":\"2025-10-27\",\"summary\":\"今日学习计划已生成\",\"tasks\":[...]}"
  }
}
```

---

## 📊 核心功能流程

```
1. 创建学习目标
   ↓
2. 设置学习时间
   ↓
3. 记录学习内容
   ↓
4. 调用 /api/plan/today
   ↓
5. LLM生成个性化计划
   ↓
6. 计划保存到数据库
   ↓
7. 返回计划给前端
```

---

## 🔑 关键文件

| 文件 | 说明 |
|------|------|
| `llm_service.go` | LLM API调用 |
| `prompt_builder.go` | Prompt构建 |
| `handlers.go` | API处理（getTodayPlan已改造） |
| `.env.example` | 环境变量配置 |
| `MVP_GUIDE.md` | 详细文档 |

---

## ⚡ 常见问题

**Q: 没有API Key可以用吗？**
A: 可以，会返回模拟数据，用于快速测试

**Q: 计划生成很慢？**
A: 正常，LLM API调用需要2-5秒，第二次会使用缓存

**Q: 如何切换LLM提供商？**
A: 修改环境变量 `LLM_PROVIDER` 和 `LLM_API_KEY`

**Q: 如何查看详细日志？**
A: 后端会打印所有操作日志，查看终端输出

---

## 💰 成本估算

| 提供商 | 每个计划成本 | 月成本（30个计划） |
|--------|------------|------------------|
| Claude | $0.025 | $0.75 |
| OpenAI | $0.05 | $1.50 |
| Gemini | $0.005 | $0.15 |

---

## 🎯 下一步

1. ✅ MVP完成
2. 📱 前端集成（显示计划）
3. 👤 多用户支持
4. 🗄️ 迁移到TDSQL-C MySQL
5. 📊 性能优化

---

**准备好了吗？开始吧！** 🚀
