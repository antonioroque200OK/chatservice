package chatcompletionstream

import (
	"github.com/antonioroque200OK/chatservice/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionConfigInputDto struct {
	Model               string
	ModelMaxString      int
	Temperature         float32
	TopP                float32
	N                   int
	Stop                []string
	MaxTokens           int
	PresencePenalty     float32
	FrequencyPenalty    float32
	InitalSystemMessage string
}

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openAiClient,
	}
}
