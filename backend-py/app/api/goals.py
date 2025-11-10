from typing import List
from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.database import get_db
from app.schemas.goal import GoalCreate, GoalUpdate, GoalResponse
from app.schemas.common import Response
from app.services.goal_service import GoalService

router = APIRouter()

@router.post("/", response_model=Response)
def create_goal(goal: GoalCreate, db: Session = Depends(get_db)):
    print(goal)
    """创建新目标"""
    try:
        db_goal = GoalService.create_goal(db, goal)
        return Response(code=0, message="创建成功", data=GoalResponse.model_validate(db_goal))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/", response_model=Response)
def get_goals(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    """获取目标列表"""
    goals = GoalService.get_goals(db, skip=skip, limit=limit)
    return Response(code=0, message="获取成功", data=[GoalResponse.model_validate(goal) for goal in goals])

@router.get("/{goal_id}", response_model=Response)
def get_goal(goal_id: int, db: Session = Depends(get_db)):
    """获取单个目标"""
    goal = GoalService.get_goal(db, goal_id)
    if not goal:
        raise HTTPException(status_code=404, detail="目标不存在")
    return Response(code=0, message="获取成功", data=GoalResponse.model_validate(goal))

@router.put("/{goal_id}", response_model=Response)
def update_goal(goal_id: int, goal_update: GoalUpdate, db: Session = Depends(get_db)):
    """更新目标"""
    goal = GoalService.update_goal(db, goal_id, goal_update)
    if not goal:
        raise HTTPException(status_code=404, detail="目标不存在")
    return Response(code=0, message="更新成功", data=GoalResponse.model_validate(goal))

@router.delete("/{goal_id}", response_model=Response)
def delete_goal(goal_id: int, db: Session = Depends(get_db)):
    """删除目标"""
    success = GoalService.delete_goal(db, goal_id)
    if not success:
        raise HTTPException(status_code=404, detail="目标不存在")
    return Response(code=0, message="删除成功")