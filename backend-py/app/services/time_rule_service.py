from typing import List
from sqlalchemy.orm import Session
from app.models.time_rule import TimeRule
from app.schemas.time_rule import TimeRuleCreate

class TimeRuleService:
    @staticmethod
    def create_time_rule(db: Session, time_rule: TimeRuleCreate) -> TimeRule:
        """创建单个时间规则"""
        db_time_rule = TimeRule(**time_rule.model_dump())
        db.add(db_time_rule)
        db.commit()
        db.refresh(db_time_rule)
        return db_time_rule
    
    @staticmethod
    def create_time_rules(db: Session, time_rules: List[TimeRuleCreate]) -> List[TimeRule]:
        """批量创建时间规则"""
        db_time_rules = [TimeRule(**rule.model_dump()) for rule in time_rules]
        db.add_all(db_time_rules)
        db.commit()
        for rule in db_time_rules:
            db.refresh(rule)
        return db_time_rules
    
    @staticmethod
    def get_time_rules(db: Session) -> List[TimeRule]:
        """获取所有时间规则"""
        return db.query(TimeRule).all()
    
    @staticmethod
    def delete_all_time_rules(db: Session) -> int:
        """删除所有时间规则"""
        deleted_count = db.query(TimeRule).delete()
        db.commit()
        return deleted_count