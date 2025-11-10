from typing import List
from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from app.database import get_db
from app.schemas.time_rule import TimeRuleCreate, TimeRuleResponse, TimeRulesCreate
from app.schemas.common import Response
from app.services.time_rule_service import TimeRuleService

router = APIRouter()

@router.post("/batch", response_model=Response)
def create_time_rules(time_rules: TimeRulesCreate, db: Session = Depends(get_db)):
    """批量创建时间规则"""
    try:
        # 先删除所有现有规则
        TimeRuleService.delete_all_time_rules(db)
        # 创建新规则
        db_time_rules = TimeRuleService.create_time_rules(db, time_rules.rules)
        return Response(code=0, message="更新成功", data=[TimeRuleResponse.model_validate(rule) for rule in db_time_rules])
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/", response_model=Response)
def get_time_rules(db: Session = Depends(get_db)):
    """获取所有时间规则"""
    time_rules = TimeRuleService.get_time_rules(db)
    return Response(code=0, message="获取成功", data=[TimeRuleResponse.model_validate(rule) for rule in time_rules])

@router.delete("/", response_model=Response)
def delete_all_time_rules(db: Session = Depends(get_db)):
    """删除所有时间规则"""
    deleted_count = TimeRuleService.delete_all_time_rules(db)
    return Response(code=0, message=f"删除成功，共删除{deleted_count}条规则")