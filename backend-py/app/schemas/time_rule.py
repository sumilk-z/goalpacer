from pydantic import BaseModel, Field, field_validator
from datetime import datetime
from typing import List

class TimeRuleBase(BaseModel):
    day_of_week: int = Field(..., ge=0, le=6, description="星期几，0表示周日")
    start_time: str = Field(..., description="开始时间，格式HH:MM")
    end_time: str = Field(..., description="结束时间，格式HH:MM")
    
    @field_validator('start_time', 'end_time')
    def validate_time_format(cls, v):
        try:
            hours, minutes = map(int, v.split(':'))
            if not (0 <= hours <= 23 and 0 <= minutes <= 59):
                raise ValueError("时间格式必须为HH:MM，且在有效范围内")
            return v
        except ValueError:
            raise ValueError("时间格式必须为HH:MM")

class TimeRuleCreate(TimeRuleBase):
    pass

class TimeRuleResponse(TimeRuleBase):
    id: int
    created_at: datetime
    updated_at: datetime
    
    class Config:
        from_attributes = True

class TimeRulesCreate(BaseModel):
    rules: List[TimeRuleCreate]