# GoalPacer 后端服务

基于 Golang + Gin + SQLite 的学习规划工具后端

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod download
```

### 2. 运行服务

```bash
go run main.go database.go models.go handlers.go
```

或者使用启动脚本：

```bash
./start-backend.sh
```

服务将在 `http://localhost:8080` 启动

### 3. 测试 API

```bash
curl http://localhost:8080/ping
```

## API 文档

### 目标管理

#### 获取所有目标
```
GET /api/goals
```

#### 创建目标
```
POST /api/goals
Content-Type: application/json

{
  "name": "算法刷题",
  "description": "每天刷 LeetCode 题目",
  "deadline": "2025-12-31"
}
```

#### 更新目标
```
PUT /api/goals/:id
Content-Type: application/json

{
  "name": "算法刷题",
  "description": "每天刷 LeetCode 题目",
  "status": "active",
  "deadline": "2025-12-31"
}
```

#### 删除目标
```
DELETE /api/goals/:id
```

### 时间规则

#### 获取时间规则
```
GET /api/time-rules
```

#### 设置时间规则
```
POST /api/time-rules
Content-Type: application/json

{
  "day_of_week": 1,
  "start_time": "09:00",
  "end_time": "17:00"
}
```

### 学习记录

#### 获取学习记录
```
GET /api/logs
GET /api/logs?goal_id=1
GET /api/logs?log_date=2025-10-26
```

#### 创建学习记录
```
POST /api/logs
Content-Type: application/json

{
  "goal_id": 1,
  "content": "学习了二叉树的前序遍历",
  "duration": 90,
  "log_date": "2025-10-26"
}
```

### 计划

#### 获取今日计划
```
GET /api/plan/today
```

#### 获取指定日期计划
```
GET /api/plan?date=2025-10-26
```

#### 创建计划
```
POST /api/plan
Content-Type: application/json

{
  "goal_id": 1,
  "plan_date": "2025-10-26",
  "content": "完成 LeetCode 5 道题"
}
```

#### 更新计划
```
PUT /api/plan/:id
Content-Type: application/json

{
  "goal_id": 1,
  "plan_date": "2025-10-26",
  "content": "完成 LeetCode 5 道题",
  "status": "completed"
}
```

#### 删除计划
```
DELETE /api/plan/:id
```

## 数据库

SQLite 数据库文件：`goalpacer.db`

### 表结构

- **goals** - 学习目标
- **time_rules** - 时间规则
- **learning_logs** - 学习记录
- **plans** - 学习计划

## 项目结构

```
backend/
├── main.go          # 主程序入口
├── database.go      # 数据库初始化
├── models.go        # 数据模型
├── handlers.go      # API 处理器
├── go.mod           # Go 模块定义
└── README.md        # 本文件
```

## 技术栈

- **框架**: Gin Web Framework
- **数据库**: SQLite3
- **语言**: Go 1.21+
