# 迭代记录

本目录记录了 GoalPacer 项目的完整开发迭代过程。

## 迭代概览

| 迭代 | 标题 | 时间 | 主要成果 |
|------|------|------|--------|
| [迭代 1](./iteration-1-mvp-foundation/) | MVP 基础实现 | 初期 | LLM 集成、缓存机制、数据库设计 |
| [迭代 2](./iteration-2-local-deployment/) | 本地部署支持 | 中期 | 本地启动脚本、开发环境配置 |
| [迭代 3](./iteration-3-feature-enhancement/) | 功能增强 | 后期 | 刷新按钮、数据清理、Bug 修复 |

## 快速导航

- **项目概览**: 查看 [MVP_ARCHITECTURE.md](./iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md)
- **快速开始**: 查看 [QUICK_MVP_START.md](./iteration-1-mvp-foundation/02-QUICK_MVP_START.md)
- **本地部署**: 查看 [LOCAL_STARTUP_GUIDE.md](./iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md)
- **功能详情**: 查看 [REFRESH_PLAN_FEATURE.md](./iteration-3-feature-enhancement/01-REFRESH_PLAN_FEATURE.md)

## 文档结构

```
docs/iterations/
├── README.md (本文件)
├── iteration-1-mvp-foundation/
│   ├── README.md
│   ├── 01-MVP_ARCHITECTURE.md
│   ├── 02-QUICK_MVP_START.md
│   ├── 03-MVP_GUIDE.md
│   ├── 04-MVP_IMPLEMENTATION_SUMMARY.md
│   ├── 05-MVP_CHECKLIST.md
│   ├── 06-MVP_COMPLETION_REPORT.md
│   └── 07-PROJECT_SUMMARY.md
├── iteration-2-local-deployment/
│   ├── README.md
│   ├── 01-LOCAL_STARTUP_GUIDE.md
│   ├── 02-LAUNCH_INSTRUCTIONS.md
│   ├── 03-QUICK_START.md
│   ├── 04-START_HERE.md
│   └── 05-READY_TO_RUN.md
└── iteration-3-feature-enhancement/
    ├── README.md
    ├── 01-REFRESH_PLAN_FEATURE.md
    ├── 02-QUICK_START_REFRESH.md
    ├── 03-INTEGRATION_COMPLETE.md
    ├── 04-INTEGRATION_TEST.md
    ├── 05-FRONTEND_BACKEND_INTEGRATION.md
    ├── 06-FEATURE_COMPLETE.md
    └── 07-TESTING_COMPLETE.md
```

## 关键技术决策

### 迭代 1: MVP 基础实现
- **LLM 选择**: 支持 Claude、OpenAI、Gemini 多提供商
- **缓存策略**: 24 小时 TTL 缓存，减少 API 调用
- **数据库**: SQLite（MVP）/ TDSQL-C MySQL（生产）
- **去重机制**: MD5 哈希基础的内容去重

### 迭代 2: 本地部署支持
- **开发环境**: 本地 Go + React 开发服务器
- **启动脚本**: 一键启动所有服务
- **测试工具**: 自动化 API 测试脚本

### 迭代 3: 功能增强
- **刷新机制**: POST 端点强制重新生成计划
- **前端状态**: 加载状态管理和用户反馈
- **数据清理**: 数据库清理工具和脚本

## 性能指标

| 指标 | 数值 |
|------|------|
| 首次计划生成 | 2-5 秒（LLM API 调用） |
| 缓存计划检索 | <100ms |
| 数据库清理 | <10ms |
| 月度成本估算 | $0.15-1.50（Gemini 到 Claude） |

## 后续计划

- [ ] 向量数据库集成（RAG 增强）
- [ ] 用户认证系统
- [ ] 多语言支持
- [ ] 移动端适配
- [ ] 高级分析仪表板
