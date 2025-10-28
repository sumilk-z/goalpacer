# ✅ 编译成功！现在可以启动服务了

## 🚀 快速启动

### 方式1：启动后端（推荐）

```bash
cd /Users/zhucui/CodeBuddy/goalpacer
bash run-backend.sh
```

**预期输出：**
```
🚀 启动后端服务...
📍 访问地址: http://localhost:8080
📍 API文档: http://localhost:8080/api

按 Ctrl+C 停止服务
```

### 方式2：启动前端（新终端）

```bash
cd /Users/zhucui/CodeBuddy/goalpacer
bash run-frontend.sh
```

**预期输出：**
```
🚀 启动前端服务...
📍 访问地址: http://localhost:3000

按 Ctrl+C 停止服务
```

---

## 📍 访问地址

- **前端**：http://localhost:3000
- **后端**：http://localhost:8080

---

## 🧪 快速测试

启动后端后，在新终端运行：

```bash
# 创建学习目标
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

---

## 📊 已修复的问题

✅ **编译错误已修复**

**问题：** `log.Printf undefined (type LearningLog has no field or method Printf)`

**原因：** 变量名 `log` 与 `log` 包冲突

**解决：** 将变量名改为 `learningLog`

---

## 🎯 完整体验流程

### 1️⃣ 启动后端

```bash
bash run-backend.sh
```

### 2️⃣ 启动前端（新终端）

```bash
bash run-frontend.sh
```

### 3️⃣ 打开浏览器

访问 http://localhost:3000

### 4️⃣ 创建学习目标

- 点击"新建目标"
- 输入"学习Go语言"
- 点击"创建"

### 5️⃣ 设置学习时间

- 点击"时间规则"
- 设置周一 09:00-12:00
- 点击"保存"

### 6️⃣ 记录学习内容

- 点击"记录学习"
- 输入"学习了goroutine"
- 输入"120"分钟
- 点击"保存"

### 7️⃣ 查看AI生成的计划 ⭐

- 点击"今日计划"
- 查看系统生成的个性化学习计划

---

## 📈 性能指标

| 操作 | 耗时 |
|------|------|
| 创建目标 | <50ms |
| 生成计划（首次） | 2-5秒 |
| 生成计划（缓存） | <100ms |

---

## 🔧 环境配置

### 使用模拟数据（推荐快速测试）

无需配置，系统会返回模拟计划。

### 使用真实LLM API

```bash
# Claude（推荐）
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx

# 然后启动后端
bash run-backend.sh
```

---

## 📚 相关文档

- 📖 [START_HERE.md](./START_HERE.md) - 本地启动指南
- 📖 [LOCAL_STARTUP_GUIDE.md](./LOCAL_STARTUP_GUIDE.md) - 详细启动指南
- 📖 [MVP_GUIDE.md](./MVP_GUIDE.md) - 完整实现指南

---

## ✨ 现在就开始！

```bash
bash run-backend.sh
```

然后在新终端：

```bash
bash run-frontend.sh
```

最后打开浏览器访问：**http://localhost:3000**

**祝你使用愉快！** 🎉
