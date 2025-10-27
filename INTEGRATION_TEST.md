# 🚀 前后端集成测试指南

## 📋 测试概述

本文档提供完整的前后端集成测试步骤，验证所有 API 功能是否正常工作。

## 🔧 环境准备

### 1. 启动后端服务

```bash
cd backend
go run main.go
```

**预期输出**:
```
🚀 GoalPacer 后端服务启动在 http://0.0.0.0:8080
```

### 2. 启动前端服务

在另一个终端：

```bash
cd frontend
npm start
```

**预期输出**:
```
Compiled successfully!
You can now view the app in the browser.
  Local:            http://localhost:3000
```

### 3. 验证服务状态

```bash
# 检查后端
curl http://localhost:8080/ping

# 预期响应
{"message":"pong"}
```

---

## 📝 测试用例

### 🎯 测试 1: 目标管理

#### 1.1 创建目标

**请求**:
```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "算法刷题",
    "description": "每天刷LeetCode题目",
    "status": "active"
  }'
```

**预期响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "name": "算法刷题",
    "description": "每天刷LeetCode题目",
    "status": "active",
    "created_at": "2025-10-27T...",
    "updated_at": "2025-10-27T..."
  }
}
```

**前端验证**:
1. 打开浏览器访问 `http://localhost:3000`
2. 进入"目标管理"页面
3. 点击"添加目标"按钮
4. 填写表单：
   - 目标名称: `Golang学习`
   - 状态: `进行中`
   - 描述: `深入学习Go语言`
5. 点击"确定"
6. **验证**: 列表中应该出现新目标

#### 1.2 获取所有目标

**请求**:
```bash
curl http://localhost:8080/api/goals
```

**预期响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "算法刷题",
      "description": "每天刷LeetCode题目",
      "status": "active",
      "created_at": "2025-10-27T...",
      "updated_at": "2025-10-27T..."
    }
  ]
}
```

**前端验证**:
1. 进入"目标管理"页面
2. **验证**: 应该看到所有已创建的目标列表

#### 1.3 更新目标

**请求**:
```bash
curl -X PUT http://localhost:8080/api/goals/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "算法刷题-进阶",
    "description": "深入学习算法",
    "status": "active"
  }'
```

**前端验证**:
1. 在目标列表中找到目标
2. 点击"编辑"按钮
3. 修改目标信息
4. 点击"确定"
5. **验证**: 列表中的目标信息应该更新

#### 1.4 删除目标

**请求**:
```bash
curl -X DELETE http://localhost:8080/api/goals/1
```

**前端验证**:
1. 在目标列表中找到目标
2. 点击"删除"按钮
3. 确认删除
4. **验证**: 目标应该从列表中消失

---

### 📚 测试 2: 学习记录

#### 2.1 创建学习记录

**请求**:
```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了二叉树的前序遍历，完成了LeetCode 144题",
    "duration": 90,
    "record_date": "2025-10-27"
  }'
```

**前端验证**:
1. 进入"学习记录"页面
2. 点击"添加记录"按钮
3. 填写表单：
   - 学习目标: 选择一个目标
   - 学习日期: `2025-10-27`
   - 学习时长: `120`
   - 学习内容: `深入学习了channel的使用`
4. 点击"确定"
5. **验证**: 
   - 列表中应该出现新记录
   - 统计数据应该更新（总记录数、总学习时长）

#### 2.2 获取学习记录

**请求**:
```bash
curl http://localhost:8080/api/logs
```

**前端验证**:
1. 进入"学习记录"页面
2. **验证**: 应该看到所有学习记录

#### 2.3 按目标筛选

**请求**:
```bash
curl "http://localhost:8080/api/logs?goal_id=1"
```

**前端验证**:
1. 进入"学习记录"页面
2. 点击筛选下拉框
3. 选择一个目标
4. **验证**: 列表应该只显示该目标的记录

#### 2.4 按时间排序

**前端验证**:
1. 进入"学习记录"页面
2. 点击排序按钮
3. 选择"最新优先"或"最早优先"
4. **验证**: 列表顺序应该改变

---

### ⏰ 测试 3: 时间规则

#### 3.1 设置时间规则

**请求**:
```bash
curl -X POST http://localhost:8080/api/time-rules \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "start_time": "09:00",
    "end_time": "11:00",
    "days": "1,2,3,4,5"
  }'
```

**预期响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "goal_id": 1,
    "start_time": "09:00",
    "end_time": "11:00",
    "days": "1,2,3,4,5",
    "created_at": "2025-10-27T..."
  }
}
```

#### 3.2 获取时间规则

**请求**:
```bash
curl http://localhost:8080/api/time-rules
```

---

### 📅 测试 4: 学习计划

#### 4.1 创建计划

**请求**:
```bash
curl -X POST http://localhost:8080/api/plan \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "plan_date": "2025-10-27",
    "content": "完成LeetCode 5道题",
    "status": "pending"
  }'
```

#### 4.2 获取今日计划

**请求**:
```bash
curl http://localhost:8080/api/plan/today
```

#### 4.3 获取指定日期计划

**请求**:
```bash
curl "http://localhost:8080/api/plan?date=2025-10-27"
```

#### 4.4 更新计划

**请求**:
```bash
curl -X PUT http://localhost:8080/api/plan/1 \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "plan_date": "2025-10-27",
    "content": "完成LeetCode 5道题",
    "status": "completed"
  }'
```

#### 4.5 删除计划

**请求**:
```bash
curl -X DELETE http://localhost:8080/api/plan/1
```

---

## ✅ 完整测试流程

### 场景 1: 完整的学习流程

1. **创建学习目标**
   - 在前端"目标管理"页面添加目标
   - 验证后端数据库中是否保存

2. **添加学习记录**
   - 在前端"学习记录"页面添加记录
   - 验证统计数据是否更新

3. **筛选和排序**
   - 测试按目标筛选
   - 测试按时间排序

4. **编辑和删除**
   - 编辑目标信息
   - 删除不需要的记录

### 场景 2: 数据持久化

1. 添加数据
2. 刷新浏览器
3. **验证**: 数据应该仍然存在（从数据库加载）

### 场景 3: 错误处理

1. 尝试添加不完整的数据
2. **验证**: 应该显示错误提示

---

## 🐛 常见问题

### 问题 1: 后端无法连接

**症状**: 前端显示"加载失败"

**解决方案**:
```bash
# 检查后端是否运行
curl http://localhost:8080/ping

# 如果没有响应，启动后端
cd backend && go run main.go
```

### 问题 2: CORS 错误

**症状**: 浏览器控制台显示 CORS 错误

**解决方案**: 后端已配置 CORS，如果仍有问题，检查：
- 后端是否正确启动
- 前端是否使用正确的 API 地址

### 问题 3: 数据不显示

**症状**: 前端页面为空

**解决方案**:
1. 打开浏览器开发者工具（F12）
2. 查看 Network 标签，检查 API 请求
3. 查看 Console 标签，查看错误信息

---

## 📊 测试检查清单

- [ ] 后端服务启动成功
- [ ] 前端服务启动成功
- [ ] 创建目标成功
- [ ] 获取目标列表成功
- [ ] 更新目标成功
- [ ] 删除目标成功
- [ ] 创建学习记录成功
- [ ] 获取学习记录成功
- [ ] 按目标筛选成功
- [ ] 按时间排序成功
- [ ] 设置时间规则成功
- [ ] 创建学习计划成功
- [ ] 获取今日计划成功
- [ ] 数据持久化成功
- [ ] 错误处理正常

---

## 🎯 下一步

所有测试通过后，可以进行：

1. **性能测试** - 测试大数据量下的性能
2. **压力测试** - 测试并发请求处理能力
3. **LLM 集成** - 集成 AI 分析和计划生成功能
4. **部署** - 部署到生产环境

---

**测试日期**: 2025-10-27  
**版本**: 1.0.0-beta
