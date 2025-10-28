# 迭代 2: 本地部署支持

**时间**: 中期开发阶段  
**目标**: 支持本地开发和测试环境  
**状态**: ✅ 完成

## 概述

本迭代为开发者提供了完整的本地部署和测试能力，包括：
- 一键启动脚本
- 本地开发环境配置
- 自动化测试工具
- 详细的启动指南

## 主要成果

### 启动脚本
- **start-local.sh**: 一键启动所有服务
- **start-backend-local.sh**: 启动后端服务
- **start-frontend-local.sh**: 启动前端服务
- **start-all.sh**: 启动所有服务

### 测试工具
- **quick-test.sh**: 自动化 API 测试脚本
- **integration-test.sh**: 集成测试脚本

### 文档
- 本地启动指南
- 快速开始文档
- 启动说明

## 使用流程

### 快速启动
```bash
# 一键启动所有服务
./start-local.sh

# 或分别启动
./start-backend-local.sh
./start-frontend-local.sh
```

### 测试 API
```bash
# 运行自动化测试
./quick-test.sh
```

## 环境配置

### 后端环境变量
```bash
# .env 文件
LLM_PROVIDER=claude
LLM_API_KEY=your-api-key
DATABASE_PATH=./goalpacer.db
```

### 前端配置
```javascript
// API 基础 URL
const API_BASE_URL = 'http://localhost:8080/api'
```

## 文件清单

| 文件 | 描述 |
|------|------|
| 01-LOCAL_STARTUP_GUIDE.md | 本地启动详细指南 |
| 02-LAUNCH_INSTRUCTIONS.md | 启动说明 |
| 03-QUICK_START.md | 快速开始 |
| 04-START_HERE.md | 从这里开始 |
| 05-READY_TO_RUN.md | 准备就绪 |

## 技术栈

| 组件 | 版本 | 端口 |
|------|------|------|
| Go 后端 | 1.19+ | 8080 |
| React 前端 | 18.x | 3000 |
| SQLite | 3.x | - |

## 故障排除

### 后端启动失败
```bash
# 检查端口占用
lsof -i :8080

# 清理数据库
./cleanup-db.sh
```

### 前端启动失败
```bash
# 清理依赖
rm -rf node_modules package-lock.json
npm install
```

## 下一步

- 功能增强（迭代 3）
- 云部署支持
- 性能优化
