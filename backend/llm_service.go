package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// LLMConfig LLM配置
type LLMConfig struct {
	Provider string // "claude", "openai", "gemini"
	APIKey   string
	Model    string
	BaseURL  string
}

// LLMService LLM服务
type LLMService struct {
	config LLMConfig
	client *http.Client
}

// Message 消息结构
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// NewLLMService 创建LLM服务
func NewLLMService() *LLMService {
	provider := os.Getenv("LLM_PROVIDER")
	if provider == "" {
		provider = "claude"
	}

	config := LLMConfig{
		Provider: provider,
		APIKey:   os.Getenv("LLM_API_KEY"),
		Model:    os.Getenv("LLM_MODEL"),
		BaseURL:  os.Getenv("LLM_BASE_URL"),
	}

	// 设置默认模型
	if config.Model == "" {
		switch provider {
		case "claude":
			config.Model = "claude-3-5-sonnet-20241022"
		case "openai":
			config.Model = "gpt-4-turbo"
		case "gemini":
			config.Model = "gemini-1.5-flash"
		}
	}

	return &LLMService{
		config: config,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

// Generate 生成计划
func (s *LLMService) Generate(systemPrompt, userPrompt string) (string, error) {
	if s.config.APIKey == "" {
		log.Println("⚠️  LLM_API_KEY 未设置，使用模拟数据")
		return s.generateMockPlan()
	}

	switch s.config.Provider {
	case "claude":
		return s.callClaude(systemPrompt, userPrompt)
	case "openai":
		return s.callOpenAI(systemPrompt, userPrompt)
	case "gemini":
		return s.callGemini(systemPrompt, userPrompt)
	default:
		return "", fmt.Errorf("unsupported provider: %s", s.config.Provider)
	}
}

// callClaude 调用Claude API
func (s *LLMService) callClaude(systemPrompt, userPrompt string) (string, error) {
	url := "https://api.anthropic.com/v1/messages"

	payload := map[string]interface{}{
		"model":       s.config.Model,
		"max_tokens":  2000,
		"system":      systemPrompt,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": userPrompt,
			},
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.config.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("❌ Claude API 调用失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Claude API 返回错误: %s", string(respBody))
		return "", fmt.Errorf("claude api error: %s", string(respBody))
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	// 提取响应内容
	if content, ok := result["content"].([]interface{}); ok && len(content) > 0 {
		if msg, ok := content[0].(map[string]interface{}); ok {
			if text, ok := msg["text"].(string); ok {
				return text, nil
			}
		}
	}

	return "", fmt.Errorf("invalid claude response format")
}

// callOpenAI 调用OpenAI API
func (s *LLMService) callOpenAI(systemPrompt, userPrompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model":       s.config.Model,
		"temperature": 0.7,
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.config.APIKey)

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("❌ OpenAI API 调用失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ OpenAI API 返回错误: %s", string(respBody))
		return "", fmt.Errorf("openai api error: %s", string(respBody))
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	// 提取响应内容
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if msg, ok := choice["message"].(map[string]interface{}); ok {
				if text, ok := msg["content"].(string); ok {
					return text, nil
				}
			}
		}
	}

	return "", fmt.Errorf("invalid openai response format")
}

// callGemini 调用Gemini API
func (s *LLMService) callGemini(systemPrompt, userPrompt string) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", s.config.Model, s.config.APIKey)

	payload := map[string]interface{}{
		"system_instruction": map[string]interface{}{
			"parts": []map[string]string{
				{"text": systemPrompt},
			},
		},
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": userPrompt},
				},
			},
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("❌ Gemini API 调用失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Gemini API 返回错误: %s", string(respBody))
		return "", fmt.Errorf("gemini api error: %s", string(respBody))
	}

	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	// 提取响应内容
	if candidates, ok := result["candidates"].([]interface{}); ok && len(candidates) > 0 {
		if candidate, ok := candidates[0].(map[string]interface{}); ok {
			if content, ok := candidate["content"].(map[string]interface{}); ok {
				if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
					if part, ok := parts[0].(map[string]interface{}); ok {
						if text, ok := part["text"].(string); ok {
							return text, nil
						}
					}
				}
			}
		}
	}

	return "", fmt.Errorf("invalid gemini response format")
}

// generateMockPlan 生成模拟计划（用于测试）
func (s *LLMService) generateMockPlan() (string, error) {
	mockPlan := `{
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
}`
	return mockPlan, nil
}
