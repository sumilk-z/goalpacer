from pydantic import BaseModel, Field
from datetime import datetime
from typing import Optional

class LearningLogBase(BaseModel):
    goal_id: int = Field(..., description="关联的目标ID")
    content: str = Field(..., description="学习内容")
    duration: int = Field(..., ge=1, description="学习时长（分钟）")
    log_date: Optional[str] = Field(None, description="学习日期，格式YYYY-MM-DD")

class LearningLogCreate(LearningLogBase):
    pass

class LearningLogResponse(LearningLogBase):
    id: int
    created_at: datetime
    updated_at: datetime
    
    class Config:
        from_attributes = True