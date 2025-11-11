from pydantic import BaseModel, Field
from datetime import datetime
from typing import Optional, List

class PlanBase(BaseModel):
    goal_id: int = Field(..., description="关联的目标ID")
    plan_date: str = Field(..., description="计划日期，格式YYYY-MM-DD")
    content: str = Field(..., description="计划内容")
    status: Optional[str] = Field("pending", description="计划状态")

class PlanCreate(PlanBase):
    pass

class PlanUpdate(BaseModel):
    content: Optional[str] = None
    status: Optional[str] = None

class PlanResponse(PlanBase):
    id: int
    created_at: datetime
    updated_at: datetime
    
    class Config:
        from_attributes = True

class TodayPlanResponse(BaseModel):
    plans: List[PlanResponse]
    message: str
    code: int = Field(0, description="success")   