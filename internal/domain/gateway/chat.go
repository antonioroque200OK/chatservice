package gateway

import (
	"context"

	"github.com/antonioroque200OK/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatById(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
