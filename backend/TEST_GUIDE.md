# GoalPacer 后端测试指南

## 快速开始

### 方式 1：运行所有测试
```bash
cd backend
go test -v
```

### 方式 2：使用测试脚本
```bash
cd backend
./test.sh
```

### 方式 3：运行特定测试
```bash
go test -v -run TestCreateGoal
go test -v -run TestGetGoals
go test -v -run TestUpdateGoal
go test -v -run TestDeleteGoal
```

## 测试覆盖范围

### 1. 目标管理 (Goals)

#### TestCreateGoal
- **功能**: 创建新的学习目标
- **测试内容**: 
  - 发送 POST 请求到 `/api/goals`
  - 验证响应状态码为 200
  - 验证响应码为 0（成功）
  - 验证返回的目标数据

```bash
go test -v -run TestCreateGoal
```

#### TestGetGoals
- **功能**: 获取所有学习目标
- **测试内容**:
  - 先创建一个目标
  - 发送 GET 请求到 `/api/goals`
  - 验证返回的目标列表不为空

```bash
go test -v -run TestGetGoals
```

#### TestUpdateGoal
- **功能**: 更新学习目标
- **测试内容**:
  - 创建一个目标
  - 发送 PUT 请求更新该目标
  - 验证更新成功

```bash
go test -v -run TestUpdateGoal
```

#### TestDeleteGoal
- **功能**: 删除学习目标
- **测试内容**:
  - 创建一个目标
  - 发送 DELETE 请求删除该目标
  - 验证删除成功

```bash
go test -v -run TestDeleteGoal
```

### 2. 时间规则 (Time Rules)

#### TestSetTimeRule
- **功能**: 设置时间规则
- **测试内容**:
  - 发送 POST 请求到 `/api/time-rules`
  - 设置周一 09:00-17:00 的学习时间
  - 验证设置成功

```bash
go test -v -run TestSetTimeRule
```

#### TestGetTimeRules
- **功能**: 获取所有时间规则
- **测试内容**:
  - 先设置一个时间规则
  - 发送 GET 请求到 `/api/time-rules`
  - 验证返回的规则列表

```bash
go test -v -run TestGetTimeRules
```

### 3. 学习记录 (Learning Logs)

#### TestCreateLog
- **功能**: 创建学习记录
- **测试内容**:
  - 先创建一个学习目标
  - 发送 POST 请求到 `/api/logs`
  - 记录学习内容和时长
  - 验证创建成功

```bash
go test -v -run TestCreateLog
```

#### TestGetLogs
- **功能**: 获取学习记录
- **测试内容**:
  - 创建目标和记录
  - 发送 GET 请求到 `/api/logs`
  - 验证返回的记录列表

```bash
go test -v -run TestGetLogs
```

### 4. 计划管理 (Plans)

#### TestCreatePlan
- **功能**: 创建学习计划
- **测试内容**:
  - 先创建一个学习目标
  - 发送 POST 请求到 `/api/plan`
  - 创建今日计划
  - 验证创建成功

```bash
go test -v -run TestCreatePlan
```

#### TestGetTodayPlan
- **功能**: 获取今日计划
- **测试内容**:
  - 创建目标和计划
  - 发送 GET 请求到 `/api/plan/today`
  - 验证返回今日的计划

```bash
go test -v -run TestGetTodayPlan
```

#### TestUpdatePlan
- **功能**: 更新计划
- **测试内容**:
  - 创建计划
  - 发送 PUT 请求更新计划状态为 "completed"
  - 验证更新成功

```bash
go test -v -run TestUpdatePlan
```

#### TestDeletePlan
- **功能**: 删除计划
- **测试内容**:
  - 创建计划
  - 发送 DELETE 请求删除计划
  - 验证删除成功

```bash
go test -v -run TestDeletePlan
```

## 测试数据流

```
创建目标 (Goal)
    ↓
创建时间规则 (TimeRule)
    ↓
创建学习记录 (LearningLog) - 关联目标
    ↓
创建学习计划 (Plan) - 关联目标
    ↓
更新/删除操作
```

## 预期结果

所有测试应该输出类似以下内容：

```
=== RUN   TestCreateGoal
--- PASS: TestCreateGoal (0.05s)
    handlers_test.go:XX: ✅ 创建目标成功: map[...]
=== RUN   TestGetGoals
--- PASS: TestGetGoals (0.03s)
    handlers_test.go:XX: ✅ 获取目标成功，共 1 个
...
```

## 常见问题

### Q: 测试失败，提示数据库错误？
A: 确保：
1. 没有其他进程占用 `goalpacer.db`
2. 有足够的磁盘空间
3. 数据库文件有读写权限

### Q: 如何只运行某个测试？
A: 使用 `-run` 标志：
```bash
go test -v -run TestCreateGoal
```

### Q: 如何查看详细的测试输出？
A: 使用 `-v` 标志获得详细输出：
```bash
go test -v
```

### Q: 如何测试特定的 API 端点？
A: 使用 curl 命令：
```bash
# 创建目标
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{"name":"算法刷题","description":"每天刷题"}'

# 获取所有目标
curl http://localhost:8080/api/goals

# 获取今日计划
curl http://localhost:8080/api/plan/today
```

## 性能测试

运行基准测试（如果需要）：
```bash
go test -bench=. -benchmem
```

## 测试覆盖率

查看测试覆盖率：
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 下一步

1. ✅ 单元测试已完成
2. 📝 集成测试（前后端联动）
3. 🔄 性能测试
4. 🚀 部署测试
