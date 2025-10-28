# 📚 快速参考卡片

## 🚀 快速开始

```bash
# 一键启动所有服务
./start-local.sh

# 访问应用
http://localhost:3000
```

## 📖 文档导航

| 需求 | 文档位置 |
|------|---------|
| 了解项目 | [docs/INDEX.md](./INDEX.md) |
| 查看迭代 | [docs/iterations/README.md](./iterations/README.md) |
| 系统架构 | [迭代 1 - 架构设计](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md) |
| 本地部署 | [迭代 2 - 启动指南](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md) |
| 刷新功能 | [迭代 3 - 刷新功能](./iterations/iteration-3-feature-enhancement/01-REFRESH_PLAN_FEATURE.md) |
| 整理说明 | [ORGANIZATION_SUMMARY.md](./ORGANIZATION_SUMMARY.md) |

## 📁 目录结构速览

```
docs/
├── INDEX.md                          # 📍 从这里开始
├── QUICK_REFERENCE.md                # 本文件
├── ORGANIZATION_SUMMARY.md           # 整理说明
└── iterations/
    ├── README.md                     # 迭代总览
    ├── iteration-1-mvp-foundation/   # MVP 基础
    ├── iteration-2-local-deployment/ # 本地部署
    └── iteration-3-feature-enhancement/ # 功能增强
```

## 🎯 按角色查找

### 👨‍💻 开发者
1. [本地启动指南](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md)
2. [系统架构](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md)
3. [集成测试](./iterations/iteration-3-feature-enhancement/04-INTEGRATION_TEST.md)

### 🏗️ 架构师
1. [MVP 架构设计](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md)
2. [项目总结](./iterations/iteration-1-mvp-foundation/07-PROJECT_SUMMARY.md)
3. [前后端集成](./iterations/iteration-3-feature-enhancement/05-FRONTEND_BACKEND_INTEGRATION.md)

### 📊 项目经理
1. [迭代总览](./iterations/README.md)
2. [MVP 完成报告](./iterations/iteration-1-mvp-foundation/06-MVP_COMPLETION_REPORT.md)
3. [功能完成报告](./iterations/iteration-3-feature-enhancement/06-FEATURE_COMPLETE.md)

### 🧪 测试人员
1. [集成测试文档](./iterations/iteration-3-feature-enhancement/04-INTEGRATION_TEST.md)
2. [测试完成报告](./iterations/iteration-3-feature-enhancement/07-TESTING_COMPLETE.md)
3. [功能检查清单](./iterations/iteration-1-mvp-foundation/05-MVP_CHECKLIST.md)

## 📋 迭代概览

### 迭代 1: MVP 基础实现
- **文档数**: 8
- **关键文件**: 
  - `01-MVP_ARCHITECTURE.md` - 架构设计
  - `02-QUICK_MVP_START.md` - 快速开始
  - `06-MVP_COMPLETION_REPORT.md` - 完成报告

### 迭代 2: 本地部署支持
- **文档数**: 5
- **关键文件**:
  - `01-LOCAL_STARTUP_GUIDE.md` - 启动指南
  - `03-QUICK_START.md` - 快速开始
  - `05-READY_TO_RUN.md` - 准备就绪

### 迭代 3: 功能增强
- **文档数**: 8
- **关键文件**:
  - `01-REFRESH_PLAN_FEATURE.md` - 刷新功能
  - `04-INTEGRATION_TEST.md` - 集成测试
  - `06-FEATURE_COMPLETE.md` - 完成报告

## 🔧 常用命令

```bash
# 启动服务
./start-local.sh

# 运行测试
./quick-test.sh
./integration-test.sh

# 清理数据库
./cleanup-db.sh

# 查看日志
tail -f backend/goalpacer.db
```

## 💡 关键概念

| 概念 | 说明 | 文档 |
|------|------|------|
| LLM 集成 | 支持 Claude、OpenAI、Gemini | [MVP_ARCHITECTURE.md](./iterations/iteration-1-mvp-foundation/01-MVP_ARCHITECTURE.md) |
| 缓存机制 | 24 小时 TTL 缓存 | [MVP_GUIDE.md](./iterations/iteration-1-mvp-foundation/03-MVP_GUIDE.md) |
| 去重算法 | MD5 哈希去重 | [MVP_IMPLEMENTATION_SUMMARY.md](./iterations/iteration-1-mvp-foundation/04-MVP_IMPLEMENTATION_SUMMARY.md) |
| 刷新功能 | 一键重新生成计划 | [REFRESH_PLAN_FEATURE.md](./iterations/iteration-3-feature-enhancement/01-REFRESH_PLAN_FEATURE.md) |

## 📞 获取帮助

### 问题排查
1. 查看 [LOCAL_STARTUP_GUIDE.md](./iterations/iteration-2-local-deployment/01-LOCAL_STARTUP_GUIDE.md) 的故障排除部分
2. 查看 [INTEGRATION_TEST.md](./iterations/iteration-3-feature-enhancement/04-INTEGRATION_TEST.md) 的测试结果

### 功能说明
1. 查看对应迭代的 README
2. 查看详细的功能文档
3. 查看完成报告中的实现细节

## 📊 统计信息

- **总文档数**: 25
- **总大小**: 212 KB
- **迭代数**: 3
- **最后更新**: 2025-10-27

---

**💡 提示**: 从 [INDEX.md](./INDEX.md) 开始获得完整的文档导航！
