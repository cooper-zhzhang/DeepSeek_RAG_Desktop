package global

const (
	MODEL_NAME     = "deepseek-r1:7b"
	MODEL_NAME_1_5 = "deepseek-r1:1.5b"
	GEN_URL        = "http://localhost:11434/api/generate"
	CHAT_URL       = "http://localhost:11434/api/chat"
	OLLAMA_URL     = "http://localhost:11434"
)

type RoleType string

const (
	RoleUser      RoleType = "user"
	RoleAssistant RoleType = "assistant"
)
