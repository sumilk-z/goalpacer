# GoalPacer 项目文档索引

## 📚 文档导航

### 快速开始
- **[迭代记录总览](./iterations/README.md)** - 查看所有迭代的完整记录
- **[MVP 架构设计](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md)** - 系统架构和技术栈
- **[快速启动指南](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md)** - 本地开发环境配置

### 迭代详情

#### 迭代 1: MVP 基础实现
核心功能实现，包括 LLM 集成、缓存机制、数据库设计

| 文档 | 描述 |
|------|------|
| [01-MVP_ARCHITECTURE.md](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md) | 完整的架构设计 |
| [02-QUICK_MVP_START.md](./iterations/iteration-1-mvp-foundation/02-QUICK_MVP_START.md) | 快速开始指南 |
| [03-MVP_GUIDE.md](./iterations/iteration-1-mvp-foundation/03-MVP_GUIDE.md) | 详细的 MVP 指南 |
| [04-MVP_IMPLEMENTATION_SUMMARY.md](./iterations/iteration-1-mvp-foundation/04-MVP_IMPLEMENTATION_SUMMARY.md) | 实现总结 |
| [05-MVP_CHECKLIST.md](./iterations/iteration-1-mvp-foundation/05-MVP_CHECKLIST.md) | 功能检查清单 |
| [06-MVP_COMPLETION_REPORT.md](./iterations/iteration-1-mvp-foundation/06-MVP_COMPLETION_REPORT.md) | 完成报告 |
| [07-PROJECT_SUMMARY.md](./iterations/iteration-1-mvp-foundation/07-PROJECT_SUMMARY.md) | 项目总结 |

#### 迭代 2: 本地部署支持
开发环境配置和本地测试能力

| 文档 | 描述 |
|------|------|
| [01-LOCAL_STARTUP_GUIDE.md](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md) | 本地启动详细指南 |
| [02-LAUNCH_INSTRUCTIONS.md](./iterations/iteration-2-local-deployment/02-LAUNCH_INSTRUCTIONS.md) | 启动说明 |
| [03-QUICK_START.md](./iterations/iteration-2-local-deployment/03-QUICK_START.md) | 快速开始 |
| [04-START_HERE.md](./iterations/iteration-2-local-deployment/04-START_HERE.md) | 从这里开始 |
| [05-READY_TO_RUN.md](./iterations/iteration-2-local-deployment/05-READY_TO_RUN.md) | 准备就绪 |

#### 迭代 3: 功能增强
刷新按钮、数据清理、Bug 修复

| 文档 | 描述 |
|------|------|
| [01-REFRESH_PLAN_FEATURE.md](./iterations/iteration-3-feature-enhancement/01-REFRESH_PLAN_FEATURE.md) | 刷新功能详细文档 |
| [02-QUICK_START_REFRESH.md](./iterations/iteration-3-feature-enhancement/02-QUICK_START_REFRESH.md) | 快速开始刷新功能 |
| [03-INTEGRATION_COMPLETE.md](./iterations/iteration-3-feature-enhancement/03-INTEGRATION_COMPLETE.md) | 集成完成报告 |
| [04-INTEGRATION_TEST.md](./iterations/iteration-3-feature-enhancement/04-INTEGRATION_TEST.md) | 集成测试文档 |
| [05-FRONTEND_BACKEND_INTEGRATION.md](./iterations/iteration-3-feature-enhancement/05-FRONTEND_BACKEND_INTEGRATION.md) | 前后端集成详情 |
| [06-FEATURE_COMPLETE.md](./iterations/iteration-3-feature-enhancement/06-FEATURE_COMPLETE.md) | 功能完成报告 |
| [07-TESTING_COMPLETE.md](./iterations/iteration-3-feature-enhancement/07-TESTING_COMPLETE.md) | 测试完成报告 |

## 🏗️ 项目结构

```
goalpacer/
├── docs/
│   ├── INDEX.md (本文件)
│   └── iterations/
│       ├── README.md (迭代总览)
│       ├── iteration-1-mvp-foundation/
│       │   ├── README.md
│       │   ├── 01-MVP_ARCHITECTURE.md
│       │   ├── 02-QUICK_MVP_START.md
│       │   ├── 03-MVP_GUIDE.md
│       │   ├── 04-MVP_IMPLEMENTATION_SUMMARY.md
│       │   ├── 05-MVP_CHECKLIST.md
│       │   ├── 06-MVP_COMPLETION_REPORT.md
│       │   └── 07-PROJECT_SUMMARY.md
│       ├── iteration-2-local-deployment/
│       │   ├── README.md
│       │   ├── 01-LOCAL_STARTUP_GUIDE.md
│       │   ├── 02-LAUNCH_INSTRUCTIONS.md
│       │   ├── 03-QUICK_START.md
│       │   ├── 04-START_HERE.md
│       │   └── 05-READY_TO_RUN.md
│       └── iteration-3-feature-enhancement/
│           ├── README.md
│           ├── 01-REFRESH_PLAN_FEATURE.md
│           ├── 02-QUICK_START_REFRESH.md
│           ├── 03-INTEGRATION_COMPLETE.md
│           ├── 04-INTEGRATION_TEST.md
│           ├── 05-FRONTEND_BACKEND_INTEGRATION.md
│           ├── 06-FEATURE_COMPLETE.md
│           └── 07-TESTING_COMPLETE.md
├── backend/
├── frontend/
└── ...
```

## 🎯 按用途查找文档

### 我想快速开始
1. 阅读 [快速启动指南](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md)
2. 运行 `./start-local.sh`
3. 访问 `http://localhost:3000`

### 我想了解系统架构
1. 阅读 [MVP 架构设计](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md)
2. 查看 [项目总结](./iterations/iteration-1-mvp-foundation/07-PROJECT_SUMMARY.md)

### 我想了解刷新功能
1. 阅读 [刷新功能详细文档](./iterations/iteration-3-feature-enhancement/01-REFRESH_PLAN_FEATURE.md)
2. 查看 [前后端集成详情](./iterations/iteration-3-feature-enhancement/05-FRONTEND_BACKEND_INTEGRATION.md)

### 我想进行集成测试
1. 阅读 [集成测试文档](./iterations/iteration-3-feature-enhancement/04-INTEGRATION_TEST.md)
2. 运行 `./integration-test.sh`

### 我想清理数据库
1. 查看 [MVP 完成报告](./iterations/iteration-1-mvp-foundation/06-MVP_COMPLETION_REPORT.md)
2. 运行 `./cleanup-db.sh`

## 📊 关键指标

| 指标 | 数值 |
|------|------|
| 总迭代数 | 3 |
| 总文档数 | 22 |
| 首次计划生成 | 2-5 秒 |
| 缓存检索 | <100ms |
| 月度成本 | $0.15-1.50 |

## 🔧 技术栈

- **后端**: Go 1.19+, Gin Framework
- **前端**: React 18.x, JavaScript
- **数据库**: SQLite (MVP) / TDSQL-C MySQL (生产)
- **LLM**: Claude, OpenAI, Gemini

## 📝 文档维护

- 所有文档按迭代组织
- 每个迭代包含 README 和详细文档
- 文档按顺序编号（01-07）
- 使用 Markdown 格式

## 🚀 后续计划

- [ ] 向量数据库集成（RAG）
- [ ] 用户认证系统
- [ ] 多语言支持
- [ ] 移动端适配
- [ ] 高级分析仪表板

---

**最后更新**: 2025-10-27  
**项目状态**: MVP 完成，功能增强完成
