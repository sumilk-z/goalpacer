# GoalPacer Backend (Python)

GoalPacer是一个学习目标管理和智能计划生成系统的后端服务，使用Python和FastAPI开发。

## 技术栈

- **框架**: FastAPI
- **数据库**: SQLite (SQLAlchemy ORM)
- **语言模型**: LangChain (支持OpenAI、Claude、Gemini)
- **包管理**: uv
- **验证**: Pydantic V2

## 快速开始

### 安装uv

首先安装uv包管理工具：

```bash
# 使用pip安装uv
pip install uv

# 或者从官方安装
# Windows
iwr -useb https://astral.sh/uv/install.ps1 | iex

# macOS/Linux
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### 安装依赖

```bash
cd backend-py
uv sync
```

### 配置环境变量

复制环境变量模板并根据需要修改：

```bash
cp .env.example .env
```

编辑`.env`文件，配置必要的API密钥（至少需要配置一个LLM API密钥以启用完整功能）。

### 初始化数据库

```bash
uv run init-db
```

### 运行服务

```bash
# 使用开发模式（自动重载）
uv run dev

# 或者直接运行
uv run start
```

服务将在 http://localhost:8000 启动。

## API文档

启动服务后，可以访问以下地址查看自动生成的API文档：

- Swagger UI: http://localhost:8000/docs
- ReDoc: http://localhost:8000/redoc

## 项目结构

```
backend-py/
├── app/                    # 应用主目录
│   ├── __init__.py         # 包初始化文件
│   ├── api/               # API路由层
│   ├── models/            # 数据库模型
│   ├── schemas/           # Pydantic数据验证模型
│   ├── services/          # 业务逻辑层
│   └── database.py        # 数据库配置
├── main.py                # 应用入口
├── pyproject.toml         # 项目配置和依赖
├── requirements.txt       # 依赖列表（备用）
├── .env                   # 环境变量
├── .env.example           # 环境变量模板
└── start.sh               # 启动脚本
```

## 主要功能

1. **目标管理**: 创建、查询、更新和删除学习目标
2. **时间规则**: 设置每周可用学习时间
3. **学习记录**: 记录每日学习内容和时长
4. **智能计划**: 基于目标和时间规则自动生成每日学习计划

## 使用uv

### 常用命令

```bash
# 安装依赖
uv sync

# 添加新依赖
uv add <package>

# 运行脚本
uv run <script>

# 运行测试
uv run pytest

# 格式化代码
uv run black .
uv run isort .
```

### 自定义脚本

项目在`pyproject.toml`中定义了以下自定义脚本：

- `uv run start`: 启动应用
- `uv run dev`: 开发模式运行（带热重载）
- `uv run init-db`: 初始化数据库

## 环境变量说明

- `OPENAI_API_KEY`: OpenAI API密钥
- `ANTHROPIC_API_KEY`: Anthropic API密钥
- `GOOGLE_API_KEY`: Google Gemini API密钥
- `DATABASE_URL`: 数据库连接字符串
- `DEBUG`: 是否开启调试模式

## 注意事项

1. 确保Python版本 >= 3.8
2. 至少配置一个LLM API密钥以使用智能计划功能
3. 在生产环境中，建议修改CORS配置以限制允许的域名