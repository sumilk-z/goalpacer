import os
from langchain.prompts import ChatPromptTemplate
from langchain.schema.output_parser import StrOutputParser
from langchain_openai import ChatOpenAI

class LLMService:
    def __init__(self):
        """初始化LLM服务，按优先级选择可用的LLM提供商"""
        self.llm = self._init_llm()
        self.output_parser = StrOutputParser()
    
    def _init_llm(self):
        """初始化LLM模型，按优先级尝试"""
        # 首先尝试Claude
        claude_api_key = os.getenv("ANTHROPIC_API_KEY")
        if claude_api_key:
            try:
                return ChatAnthropic(model="claude-3-opus-20240229", api_key=claude_api_key)
            except Exception as e:
                print(f"Claude初始化失败: {e}")
        
        # 然后尝试OpenAI
        openai_api_key = os.getenv("OPENAI_API_KEY")
        if openai_api_key:
            try:
                return ChatOpenAI(model="gpt-4-turbo", api_key=openai_api_key)
            except Exception as e:
                print(f"OpenAI初始化失败: {e}")
        
        # 最后尝试Gemini
        gemini_api_key = os.getenv("GOOGLE_API_KEY")
        if gemini_api_key:
            try:
                return ChatGoogleGenerativeAI(model="gemini-pro", api_key=gemini_api_key)
            except Exception as e:
                print(f"Gemini初始化失败: {e}")
        
        # 如果没有配置任何API Key，返回None
        print("warning: 未配置任何LLM API Key，将使用模拟数据")
        return None
    
    def generate(self, prompt: str) -> str:
        """生成LLM响应"""
        if not self.llm:
            # 返回模拟数据
            return self._generate_mock_response(prompt)
        
        try:
            # 使用LangChain的链来处理
            chat_prompt = ChatPromptTemplate.from_template("{prompt}")
            chain = chat_prompt | self.llm | self.output_parser
            response = chain.invoke({"prompt": prompt})
            return response
        except Exception as e:
            print(f"LLM生成失败: {e}")
            # 失败时返回模拟数据
            return self._generate_mock_response(prompt)
    
    def _generate_mock_response(self, prompt: str) -> str:
        """生成模拟的LLM响应"""
        return """
{
  "date": "2025-10-27",
  "summary": "今日学习计划已生成",
  "tasks": [
    {
      "goal_id": 1,
      "title": "Go语言基础学习",
      "description": "学习Go语言的并发编程和goroutine",
      "duration_minutes": 90,
      "start_time": "09:00",
      "priority": "high"
    },
    {
      "goal_id": 2,
      "title": "React组件开发",
      "description": "完成TDesign组件集成和优化",
      "duration_minutes": 60,
      "start_time": "14:00",
      "priority": "medium"
    },
    {
      "goal_id": 3,
      "title": "数据库设计复习",
      "description": "复习TDSQL-C MySQL的最佳实践",
      "duration_minutes": 45,
      "start_time": "16:00",
      "priority": "medium"
    }
  ],
  "total_duration_minutes": 195,
  "notes": "这是一个模拟计划，用于测试。请设置LLM_API_KEY环境变量以使用真实的LLM服务。"
}
"""