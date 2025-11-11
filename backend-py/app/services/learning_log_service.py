from typing import List, Optional
from sqlalchemy.orm import Session
from datetime import datetime
from app.models.learning_log import LearningLog
from app.schemas.learning_log import LearningLogCreate

class LearningLogService:
    @staticmethod
    def create_learning_log(db: Session, learning_log: LearningLogCreate) -> LearningLog:
        """创建学习记录"""
        # 如果没有提供log_date，使用当前日期
        if not learning_log.log_date:
            learning_log.log_date = datetime.now().strftime("%Y-%m-%d")
        
        db_learning_log = LearningLog(**learning_log.model_dump())
        db.add(db_learning_log)
        db.commit()
        db.refresh(db_learning_log)
        return db_learning_log
    
    @staticmethod
    def get_learning_logs(db: Session, goal_id: Optional[int] = None, skip: int = 0, limit: int = 100) -> List[LearningLog]:
        """获取学习记录列表，可以按目标过滤"""
        query = db.query(LearningLog)
        if goal_id:
            query = query.filter(LearningLog.goal_id == goal_id)
        return query.order_by(LearningLog.log_date.desc()).offset(skip).limit(limit).all()
    
    @staticmethod
    def get_learning_log(db: Session, log_id: int) -> Optional[LearningLog]:
        """获取单个学习记录"""
        return db.query(LearningLog).filter(LearningLog.id == log_id).first()
    
    @staticmethod
    def delete_learning_log(db: Session, log_id: int) -> bool:
        """删除学习记录"""
        db_learning_log = db.query(LearningLog).filter(LearningLog.id == log_id).first()
        if db_learning_log:
            db.delete(db_learning_log)
            db.commit()
            return True
        return False