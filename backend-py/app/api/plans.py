from typing import List, Optional
from fastapi import APIRouter, Depends, HTTPException, Query
from sqlalchemy.orm import Session
from app.database import get_db
from app.schemas.plan import PlanCreate, PlanUpdate, PlanResponse, TodayPlanResponse
from app.schemas.common import Response
from app.services.plan_service import PlanService

router = APIRouter()

@router.post("/", response_model=Response)
def create_plan(plan: PlanCreate, db: Session = Depends(get_db)):
    """创建计划"""
    try:
        db_plan = PlanService.create_plan(db, plan)
        return Response(code=0, message="success", data=PlanResponse.model_validate(db_plan))
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/today", response_model=Response)
def get_today_plan(db: Session = Depends(get_db)):
    """获取今日计划"""
    return Response(code=0, message="success", data=PlanService.get_today_plan(db))

@router.get("/", response_model=Response)
def get_plans(
    goal_id: int = Query(None, description="目标ID"),
    plan_date: str = Query(None, description="计划日期，格式YYYY-MM-DD"),
    db: Session = Depends(get_db)
):
    """获取计划列表"""
    plans = PlanService.get_plans(db, goal_id=goal_id, plan_date=plan_date)
    return Response(code=0, message="success", data=[PlanResponse.model_validate(plan) for plan in plans])

@router.get("/{plan_id}", response_model=Response)
def get_plan(plan_id: int, db: Session = Depends(get_db)):
    """获取单个计划"""
    plan = PlanService.get_plan(db, plan_id)
    if not plan:
        raise HTTPException(status_code=404, detail="计划不存在")
    return Response(code=0, message="success", data=PlanResponse.model_validate(plan))

@router.put("/{plan_id}", response_model=Response)
def update_plan(plan_id: int, plan_update: PlanUpdate, db: Session = Depends(get_db)):
    """更新计划"""
    plan = PlanService.update_plan(db, plan_id, plan_update)
    if not plan:
        raise HTTPException(status_code=404, detail="计划不存在")
    return Response(code=0, message="update success", data=PlanResponse.model_validate(plan))

@router.delete("/{plan_id}", response_model=Response)
def delete_plan(plan_id: int, db: Session = Depends(get_db)):
    """删除计划"""
    success = PlanService.delete_plan(db, plan_id)
    if not success:
        raise HTTPException(status_code=404, detail="计划不存在")
    return Response(code=0, message="delete success")
