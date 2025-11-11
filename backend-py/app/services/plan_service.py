from typing import List, Optional
from sqlalchemy.orm import Session
from datetime import datetime
from app.models.plan import Plan
from app.models.goal import Goal
from app.models.time_rule import TimeRule
from app.models.learning_log import LearningLog
from app.schemas.plan import PlanCreate, PlanUpdate, PlanResponse
from app.services.llm_service import LLMService

class PlanService:
    @staticmethod
    def create_plan(db: Session, plan: PlanCreate) -> Plan:
        """创建计划"""
        db_plan = Plan(**plan.model_dump())
        db.add(db_plan)
        db.commit()
        db.refresh(db_plan)
        return db_plan
    
    @staticmethod
    def get_plans(db: Session, goal_id: Optional[int] = None, plan_date: Optional[str] = None) -> List[Plan]:
        """获取计划列表，可以按目标和日期过滤"""
        query = db.query(Plan)
        if goal_id:
            query = query.filter(Plan.goal_id == goal_id)
        if plan_date:
            query = query.filter(Plan.plan_date == plan_date)
        return query.order_by(Plan.plan_date.desc()).all()
    
    @staticmethod
    def get_plan(db: Session, plan_id: int) -> Optional[Plan]:
        """获取单个计划"""
        return db.query(Plan).filter(Plan.id == plan_id).first()
    
    @staticmethod
    def update_plan(db: Session, plan_id: int, plan_update: PlanUpdate) -> Optional[Plan]:
        """更新计划"""
        db_plan = db.query(Plan).filter(Plan.id == plan_id).first()
        if db_plan:
            update_data = plan_update.model_dump(exclude_unset=True)
            for field, value in update_data.items():
                setattr(db_plan, field, value)
            db.commit()
            db.refresh(db_plan)
        return db_plan
    
    @staticmethod
    def delete_plan(db: Session, plan_id: int) -> bool:
        """删除计划"""
        db_plan = db.query(Plan).filter(Plan.id == plan_id).first()
        if db_plan:
            db.delete(db_plan)
            db.commit()
            return True
        return False
    
    @staticmethod
    def get_today_plan(db: Session) -> PlanResponse:
        """获取今日计划"""
        today = datetime.now().strftime("%Y-%m-%d")
        plans = db.query(Plan).filter(Plan.plan_date == today).all()
        
        if not plans:
            # 如果没有今日计划，尝试生成
            plans = PlanService.generate_daily_plan(db)
            message = "今日计划已自动生成"
        else:
            message = "今日计划已存在"
        print(plans[0].content)
        # return TodayPlanResponse(plans=plans, message=message)
        #  @TODO 这里如果用户有多个 goal ，那么每个goal对应的plan 都要返回，前端不解析list，后续改一下
        return PlanResponse(
            id=plans[0].id,
            goal_id=plans[0].goal_id,
            plan_date=plans[0].plan_date,
            content=plans[0].content,
            status=plans[0].status,
            created_at=plans[0].created_at,
            updated_at=plans[0].updated_at,
        )
    
    @staticmethod
    def generate_daily_plan(db: Session) -> List[Plan]:
        """生成每日学习计划"""
        today = datetime.now().strftime("%Y-%m-%d")
        today_weekday = datetime.now().weekday()
        
        # 获取活跃的目标
        active_goals = db.query(Goal).filter(Goal.status == "active").all()
        if not active_goals:
            return []
        
        # 获取今日的时间规则
        time_rules = db.query(TimeRule).filter(TimeRule.day_of_week == today_weekday).all()
        
        # 获取最近的学习记录作为参考
        recent_logs = db.query(LearningLog).order_by(LearningLog.created_at.desc()).limit(10).all()
        
        # 构建LLM提示词
        print("构建LLM提示词 LLMService")
        llm_service = LLMService()
        plans = []
        
        for goal in active_goals:
            # 为每个活跃目标生成计划
            prompt = PlanService._build_plan_prompt(goal, time_rules, recent_logs, today)
            plan_content = llm_service.generate(prompt)
            
            if plan_content:
                # 创建计划记录
                plan = Plan(
                    goal_id=goal.id,
                    plan_date=today,
                    content=plan_content,
                    status="pending"
                )
                db.add(plan)
                plans.append(plan)
        
        if plans:
            db.commit()
        
        return plans
    
    @staticmethod
    def _build_plan_prompt(goal: Goal, time_rules: List[TimeRule], recent_logs: List[LearningLog], today: str) -> str:
        """构建计划生成的提示词"""
        prompt = f"""
        请为用户生成{today}的学习计划，基于以下信息：
        
        目标名称：{goal.name}
        目标描述：{goal.description}
        截止日期：{goal.deadline if goal.deadline else '无'}
        
        今日可用学习时间：
        """
        
        if time_rules:
            for rule in time_rules:
                prompt += f"- {rule.start_time} 至 {rule.end_time}\n"
        else:
            prompt += "- 今日无指定学习时间，建议安排1-2小时\n"
        
        prompt += "\n最近的学习记录：\n"
        if recent_logs:
            for log in recent_logs[:3]:  # 只显示最近3条
                prompt += f"- {log.log_date}: {log.content[:100]}...（{log.duration}分钟）\n"
        else:
            prompt += "- 暂无学习记录\n"
        
        prompt += "\n请生成一份详细、可执行的学习计划，包括具体的学习内容、建议时长和优先级。计划应当合理、可完成，并且与用户的目标相关。"
        
        return prompt