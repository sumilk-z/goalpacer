# 🚀 快速启动 - 体验刷新计划功能

## ⚡ 3步启动

### 1️⃣ 启动后端

```bash
cd /Users/zhucui/CodeBuddy/goalpacer/backend
./goalpacer-backend
```

### 2️⃣ 启动前端（新终端）

```bash
cd /Users/zhucui/CodeBuddy/goalpacer/frontend
npm install
npm start
```

### 3️⃣ 打开浏览器

访问：**http://localhost:3000**

---

## 📝 快速体验流程（5分钟）

### 步骤1：创建学习目标

1. 点击"目标管理"菜单
2. 点击"新建目标"按钮
3. 输入目标名称：`学习Go语言`
4. 输入描述：`掌握Go并发编程`
5. 点击"创建"

### 步骤2：设置学习时间

1. 点击"时间规则"菜单
2. 选择星期一
3. 设置时间：09:00 - 12:00
4. 点击"保存"

### 步骤3：记录学习内容

1. 点击"学习记录"菜单
2. 点击"新建记录"按钮
3. 选择目标：`学习Go语言`
4. 输入内容：`学习了goroutine和channel`
5. 输入时长：`120`分钟
6. 点击"保存"

### 步骤4：查看计划

1. 点击"今日计划"菜单
2. 查看系统生成的学习计划

### 步骤5：刷新计划 ⭐

1. 在"今日计划"页面右上角找到"🔄 刷新计划"按钮
2. 点击按钮
3. 等待2-5秒（LLM生成新计划）
4. 查看新生成的计划

---

## 🧪 API 测试

### 创建目标

```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "学习Go语言",
    "description": "掌握Go并发编程",
    "deadline": "2025-12-31"
  }'
```

### 设置时间规则

```bash
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "day_of_week": 1,
    "start_time": "09:00",
    "end_time": "12:00"
  }'
```

### 记录学习内容

```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了goroutine和channel",
    "duration": 120
  }'
```

### 获取今日计划

```bash
curl -X GET http://localhost:8080/api/plan/today
```

### 刷新今日计划 ⭐

```bash
curl -X POST http://localhost:8080/api/plan/today/refresh
```

---

## 📊 新功能说明

### 🔄 一键刷新计划

**位置：** 今日计划页面右上角

**功能：**
- 强制重新生成学习计划
- 忽略24小时缓存
- 基于最新的学习数据
- 每次都调用LLM生成新计划

**使用场景：**
- 学习目标或时间规则有变化
- 想要获得不同的计划建议
- 测试LLM的计划生成能力

**性能：**
- 首次生成：2-5秒
- 后续查询：<100ms（使用缓存）

---

## 🧹 数据库清理

### 清理脏数据

```bash
bash /Users/zhucui/CodeBuddy/goalpacer/cleanup-db.sh
```

**清理内容：**
- ✅ 删除所有目标
- ✅ 删除所有时间规则
- ✅ 删除所有学习记录
- ✅ 删除所有计划

**清理结果：**
```
清理前：69条记录
清理后：0条记录
```

---

## 🎯 功能完成情况

| 功能 | 状态 |
|------|------|
| 后端刷新接口 | ✅ |
| 前端刷新按钮 | ✅ |
| 计划自动解析 | ✅ |
| 加载状态显示 | ✅ |
| 成功/失败提示 | ✅ |
| 数据库清理 | ✅ |

---

## 📍 访问地址

- **前端：** http://localhost:3000
- **后端：** http://localhost:8080

---

## 🐛 常见问题

### Q: 刷新计划很慢？

**A:** 正常，首次需要2-5秒调用LLM API。第二次查询会使用缓存（<100ms）。

### Q: 没有API Key可以用吗？

**A:** 可以，系统会使用模拟数据生成计划。

### Q: 如何使用真实LLM API？

**A:** 设置环境变量后启动后端：
```bash
export LLM_PROVIDER=claude
export LLM_API_KEY=sk-ant-xxxxxxxxxxxxx
./goalpacer-backend
```

### Q: 如何清理数据库？

**A:** 运行清理脚本：
```bash
bash cleanup-db.sh
```

---

## 📚 详细文档

- 📖 [REFRESH_PLAN_FEATURE.md](./REFRESH_PLAN_FEATURE.md) - 完整功能说明
- 📖 [START_HERE.md](./START_HERE.md) - 快速启动指南
- 📖 [MVP_GUIDE.md](./MVP_GUIDE.md) - 完整实现指南

---

## ✨ 现在就开始！

```bash
# 启动后端
cd backend
./goalpacer-backend

# 新终端启动前端
cd frontend
npm start

# 打开浏览器
# http://localhost:3000
```

**祝你使用愉快！** 🎉
