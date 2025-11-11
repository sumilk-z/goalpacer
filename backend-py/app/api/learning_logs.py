from typing import List
from fastapi import APIRouter, Depends, HTTPException, Query
from sqlalchemy.orm import Session
from app.database import get_db
from app.schemas.learning_log import LearningLogCreate, LearningLogResponse
from app.schemas.common import Response
from app.services.learning_log_service import LearningLogService

router = APIRouter()

@router.post("/", response_model=Response)
def create_learning_log(learning_log: LearningLogCreate, db: Session = Depends(get_db)):
    """创建学习记录"""
    try:
        db_learning_log = LearningLogService.create_learning_log(db, learning_log)
        return Response(code=0, message="创建成功", data=LearningLogResponse.model_validate(db_learning_log))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/", response_model=Response)
def get_learning_logs(
    goal_id: int = Query(None, description="目标ID，不传则获取所有"),
    skip: int = 0,
    limit: int = 100,
    db: Session = Depends(get_db)
):
    """获取学习记录列表"""
    learning_logs = LearningLogService.get_learning_logs(db, goal_id=goal_id, skip=skip, limit=limit)
    return Response(code=0, message="获取成功", data=[LearningLogResponse.model_validate(log) for log in learning_logs])

@router.get("/{log_id}", response_model=Response)
def get_learning_log(log_id: int, db: Session = Depends(get_db)):
    """获取单个学习记录"""
    learning_log = LearningLogService.get_learning_log(db, log_id)
    if not learning_log:
        raise HTTPException(status_code=404, detail="学习记录不存在")
    return Response(code=0, message="获取成功", data=LearningLogResponse.model_validate(learning_log))

@router.delete("/{log_id}", response_model=Response)
def delete_learning_log(log_id: int, db: Session = Depends(get_db)):
    """删除学习记录"""
    success = LearningLogService.delete_learning_log(db, log_id)
    if not success:
        raise HTTPException(status_code=404, detail="学习记录不存在")
    return Response(code=0, message="删除成功")