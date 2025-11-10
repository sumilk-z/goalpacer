from app.api import goals, time_rules, learning_logs, plans
from fastapi import APIRouter

api_router = APIRouter()

api_router.include_router(goals.router, prefix="/goals", tags=["goals"])
api_router.include_router(time_rules.router, prefix="/time-rules", tags=["time_rules"])
api_router.include_router(learning_logs.router, prefix="/learning-logs", tags=["learning_logs"])
api_router.include_router(plans.router, prefix="/plan", tags=["plan"])