package request

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatReq struct {
	ConversationID string    `json:"conversationId"` // uuid
	Model          Model     `json:"model"`
	Messages       []Message `json:"messages"`
	Key            string    `json:"key"`
	Prompt         string    `json:"prompt"` // 提示词，就是openai的System
}