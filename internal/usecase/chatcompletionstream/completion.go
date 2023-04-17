package chatcompletionstream

import (
	"context"
	"errors"

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

type ChatCompletionInputDto struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      ChatCompletionConfigInputDto
}

type ChatCompletionOutputDto struct {
	ChatID  string
	UserID  string
	Content string
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

func (uc *ChatCompletionUseCase) Execute(ctx context.Context, input ChatCompletionInputDto) (*ChatCompletionOutputDto, error) {
	chat, err := uc.ChatGateway.FindChatById(ctx, input.ChatID)
	if err != nil {
		if err.Error() == "chat not found" {
			// * create new chat (entity)
			// TODO: implemet 'createNewChat'
			chat, err = createNewChat(input)
			if err != nil {
				return nil, errors.New("error creating new chat: " + err.Error())
			}
			// * save on database
			err = uc.ChatGateway.CreateChat(ctx, chat)
			if err != nil {
				return nil, errors.New("error persisting new chat: " + err.Error())
			}
		} else {
			return nil, errors.New("error fetching chat: " + err.Error())
		}
	}
}
