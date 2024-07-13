package port

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

type ConversationService interface {
	CreateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error)
	CreateConversationWithMessage(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, *domain.Message, error)
	GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error)
}

type ConversationRepository interface {
	// CreateUser inserts a new user into the database
	CreateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error)
	GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error)
}
