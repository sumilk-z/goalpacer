# GoalPacer 快速开始指南

## 项目概述

**GoalPacer** 是一个基于 LLM 的通用学习规划工具，帮助用户制定和执行个性化的学习计划。

- **前端**: React + TDesign
- **后端**: Golang + Gin + SQLite
- **数据库**: SQLite3

## 一键启动

### 方式 1：启动前后端（推荐）

```bash
cd /Users/zhucui/CodeBuddy/20251027000233
./start-all.sh
```

然后访问：
- 🌐 前端: http://localhost:3000
- 🔧 后端: http://localhost:8080

### 方式 2：分别启动

**启动前端**：
```bash
cd frontend
npm start
```

**启动后端**：
```bash
cd backend
./start-backend.sh
```

## 前端开发

### 快速重启前端
```bash
./restart-frontend.sh
```

### 前端页面

1. **目标管理** - 创建和管理学习目标
2. **时间配置** - 设置每周的学习时间规则
3. **学习记录** - 记录每日学习内容和时长
4. **今日计划** - 查看今日的学习计划
5. **设置** - 系统配置

## 后端开发

### 运行测试
```bash
cd backend
go test -v
```

### 查看测试结果
```bash
cat TEST_RESULTS.md
```

### API 文档
```bash
cat README.md
```

## 项目结构

```
/Users/zhucui/CodeBuddy/20251027000233/
├── frontend/                 # React 前端项目
│   ├── src/
│   │   ├── pages/           # 页面组件
│   │   ├── App.js           # 主应用
│   │   └── index.js         # 入口
│   ├── package.json
│   └── README.md
├── backend/                  # Golang 后端项目
│   ├── main.go              # 主程序
│   ├── database.go          # 数据库初始化
│   ├── models.go            # 数据模型
│   ├── handlers.go          # API 处理器
│   ├── handlers_test.go     # 单元测试
│   ├── go.mod               # Go 模块
│   ├── README.md            # 后端文档
│   └── TEST_GUIDE.md        # 测试指南
├── start-all.sh             # 一键启动脚本
├── restart-frontend.sh      # 前端重启脚本
└── QUICK_START.md           # 本文件
```

## 常见命令

### 前端相关
```bash
# 启动前端开发服务器
cd frontend && npm start

# 重启前端（清理缓存）
./restart-frontend.sh

# 构建前端
cd frontend && npm run build
```

### 后端相关
```bash
# 启动后端服务
cd backend && ./start-backend.sh

# 运行所有测试
cd backend && go test -v

# 运行特定测试
cd backend && go test -v -run TestCreateGoal

# 查看测试覆盖率
cd backend && go test -cover
```

## API 快速参考

### 创建目标
```bash
curl -X POST http://localhost:8080/api/goals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "算法刷题",
    "description": "每天刷 LeetCode 题目"
  }'
```

### 获取所有目标
```bash
curl http://localhost:8080/api/goals
```

### 创建学习记录
```bash
curl -X POST http://localhost:8080/api/logs \
  -H "Content-Type: application/json" \
  -d '{
    "goal_id": 1,
    "content": "学习了二叉树",
    "duration": 90,
    "log_date": "2025-10-27"
  }'
```

### 获取今日计划
```bash
curl http://localhost:8080/api/plan/today
```

## 数据库

SQLite 数据库文件位置：`backend/goalpacer.db`

### 表结构

**goals** - 学习目标
```sql
CREATE TABLE goals (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  status TEXT DEFAULT 'active',
  deadline DATE,
  created_at DATETIME,
  updated_at DATETIME
);
```

**time_rules** - 时间规则
```sql
CREATE TABLE time_rules (
  id INTEGER PRIMARY KEY,
  day_of_week INTEGER,
  start_time TEXT,
  end_time TEXT,
  created_at DATETIME,
  updated_at DATETIME
);
```

**learning_logs** - 学习记录
```sql
CREATE TABLE learning_logs (
  id INTEGER PRIMARY KEY,
  goal_id INTEGER,
  content TEXT,
  duration INTEGER,
  log_date DATE,
  created_at DATETIME,
  updated_at DATETIME
);
```

**plans** - 学习计划
```sql
CREATE TABLE plans (
  id INTEGER PRIMARY KEY,
  goal_id INTEGER,
  plan_date DATE,
  content TEXT,
  status TEXT DEFAULT 'pending',
  created_at DATETIME,
  updated_at DATETIME
);
```

## 故障排除

### 前端问题

**问题**: 页面没有更新
**解决**: 运行 `./restart-frontend.sh` 重启前端服务

**问题**: 端口 3000 被占用
**解决**: 
```bash
lsof -ti:3000 | xargs kill -9
```

### 后端问题

**问题**: 数据库连接失败
**解决**: 确保 `backend/goalpacer.db` 文件存在且有读写权限

**问题**: 测试失败
**解决**: 
```bash
cd backend
rm -f goalpacer.db
go test -v
```

## 开发工作流

1. **修改前端代码** → 自动热更新（或运行 `./restart-frontend.sh`）
2. **修改后端代码** → 运行 `go test -v` 验证 → 重启后端
3. **测试 API** → 使用 curl 或 Postman
4. **查看日志** → 检查终端输出

## 下一步

- [ ] 实现 LLM 分析功能
- [ ] 实现计划生成功能
- [ ] 前后端集成测试
- [ ] 性能优化
- [ ] 部署到生产环境

## 文档

- [前端 README](frontend/README.md)
- [后端 README](backend/README.md)
- [后端测试指南](backend/TEST_GUIDE.md)
- [测试结果](backend/TEST_RESULTS.md)

## 联系方式

如有问题，请查看相应的文档或运行测试进行诊断。

---

**最后更新**: 2025-10-27
**版本**: 1.0.0
