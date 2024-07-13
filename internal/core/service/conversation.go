package service

import (
	"context"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/Coke3a/TalkPenguin/internal/core/port"
)

type ConversationService struct {
	ConversRepo  port.ConversationRepository
	messageSv port.MessageService
	promptSv port.PromptService
	// cache port.CacheRepository
}

func NewConversationService(ConversRepo port.ConversationRepository, messageSv port.MessageService, promptSv port.PromptService) *ConversationService {
	return &ConversationService{
		ConversRepo,
		messageSv,
		promptSv,
		// cache,
	}
}

func (cv *ConversationService) CreateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {


	// check prompt is exists

	conversation, err := cv.ConversRepo.CreateConversation(ctx, conversation)
	if err == nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return conversation, nil
}

func (cv *ConversationService) CreateConversationWithMessage(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, *domain.Message, error) {

	// check prompt is exists
	conversation, err := cv.ConversRepo.CreateConversation(ctx, conversation)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, nil, err
		}
		return nil, nil, domain.ErrInternal
	}

	message, err := cv.messageSv.CreateMessage(ctx, conversation, enum.MessageTypes.Ai)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, nil, err
		}
		return nil, nil, domain.ErrInternal
	}

	return conversation, message, nil
}



func (cv *ConversationService) GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error) {

	conversation, err := cv.ConversRepo.GetConversationById(ctx, conversationId)
	if err == nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return conversation, nil
}