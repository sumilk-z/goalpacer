from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from contextlib import asynccontextmanager
import os
from dotenv import load_dotenv

from app.database import init_db
from app.api import api_router

# 加载环境变量
load_dotenv()

# 应用生命周期管理
@asynccontextmanager
async def lifespan(app: FastAPI):
    # 启动时初始化数据库
    init_db()
    print("数据库初始化完成")
    yield
    # 关闭时的清理工作
    print("应用关闭")

# 创建FastAPI应用实例
app = FastAPI(
    title="GoalPacer API",
    description="学习目标管理和智能计划生成系统API",
    version="1.0.0",
    lifespan=lifespan
)

# 配置CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # 在生产环境中应该设置具体的前端域名
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 包含API路由
app.include_router(api_router, prefix="/api")

# 根路径
@app.get("/")
def read_root():
    return {
        "message": "Welcome to GoalPacer API",
        "version": "1.0.0",
        "docs": "/docs"
    }

# 健康检查端点
@app.get("/health")
def health_check():
    return {"status": "healthy"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=8080,
        reload=True
    )