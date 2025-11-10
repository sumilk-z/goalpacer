from typing import List, Optional
from sqlalchemy.orm import Session
from app.models.goal import Goal
from app.schemas.goal import GoalCreate, GoalUpdate

class GoalService:
    @staticmethod
    def create_goal(db: Session, goal: GoalCreate) -> Goal:
        """创建新目标"""
        db_goal = Goal(**goal.model_dump())
        db.add(db_goal)
        db.commit()
        db.refresh(db_goal)
        return db_goal
    
    @staticmethod
    def get_goals(db: Session, skip: int = 0, limit: int = 100) -> List[Goal]:
        """获取目标列表"""
        return db.query(Goal).offset(skip).limit(limit).all()
    
    @staticmethod
    def get_goal(db: Session, goal_id: int) -> Optional[Goal]:
        """获取单个目标"""
        return db.query(Goal).filter(Goal.id == goal_id).first()
    
    @staticmethod
    def update_goal(db: Session, goal_id: int, goal_update: GoalUpdate) -> Optional[Goal]:
        """更新目标"""
        db_goal = db.query(Goal).filter(Goal.id == goal_id).first()
        if db_goal:
            update_data = goal_update.model_dump(exclude_unset=True)
            for field, value in update_data.items():
                setattr(db_goal, field, value)
            db.commit()
            db.refresh(db_goal)
        return db_goal
    
    @staticmethod
    def delete_goal(db: Session, goal_id: int) -> bool:
        """删除目标"""
        db_goal = db.query(Goal).filter(Goal.id == goal_id).first()
        if db_goal:
            db.delete(db_goal)
            db.commit()
            return True
        return False