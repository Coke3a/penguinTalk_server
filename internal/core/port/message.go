package port

import (
	"context"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

type MessageService interface {
	CreateMessage(ctx context.Context, conversation *domain.Conversation, messageType enum.MessageType) (*domain.Message, error)
	ExchangingMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetAllMessagesByConversationId(ctx context.Context, conversationId uint64) (*[]domain.Message, error)
	GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error)
}

type MessageRepository interface {
	// CreateUser inserts a new user into the database
	SaveMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetAllMessagesByConversationId(ctx context.Context, conversationId uint64) ([]domain.Message, error)
}
