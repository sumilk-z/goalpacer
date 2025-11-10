from pydantic import BaseModel, Field
from datetime import datetime
from typing import Optional

class GoalBase(BaseModel):
    name: str = Field(..., description="目标名称")
    description: Optional[str] = Field(None, description="目标描述")
    status: Optional[str] = Field("active", description="目标状态")
    deadline: Optional[str] = Field(None, description="截止日期")

class GoalCreate(GoalBase):
    pass

class GoalUpdate(BaseModel):
    name: Optional[str] = None
    description: Optional[str] = None
    status: Optional[str] = None
    deadline: Optional[str] = None

class GoalResponse(GoalBase):
    id: int
    created_at: datetime
    updated_at: datetime
    
    class Config:
        from_attributes = True