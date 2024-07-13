package service

import (
	"context"
	"log/slog"

	enum "github.com/Coke3a/TalkPenguin/internal/adapter/enum"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/Coke3a/TalkPenguin/internal/core/port"
)

type MessageService struct {
	messageRepo  port.MessageRepository
	conversRepo  port.ConversationRepository
	promptSv  port.PromptService
	mtSv 		port.ModelTransactionService
	// cache port.CacheRepository
}

func NewMessageService(messageRepo port.MessageRepository, conversRepo  port.ConversationRepository, promptSv  port.PromptService, mtSv port.ModelTransactionService) *MessageService {
	return &MessageService{
		messageRepo,
		conversRepo,
		promptSv,
		mtSv,
		// cache,
	}
}


func (ms *MessageService) CreateMessage(ctx context.Context, conversation *domain.Conversation,  messageType enum.MessageType) (*domain.Message, error) {

	prompt, err := ms.promptSv.GetPrompt(ctx, conversation.PromptId)
	if err != nil {
		return nil, err
	}

	modelTransaction, err :=  ms.mtSv.RequestToModel(ctx, prompt.Prompt, prompt.Prompt2, nil)
	if err != nil {
		return nil, err
	}
	generatedMessage, err := ms.mtSv.GetResponseMessage(ctx, modelTransaction)
	if err != nil {
		return nil, err
	}

	// conversation.PromptId := 123
	messageAudio := "123"

	message := &domain.Message{
		ConversationId: conversation.ConversationId,
		UserId:         conversation.UserId,
		MtId:			modelTransaction.MtId,
		MessageText: generatedMessage,
		MessageAudio: messageAudio,
		MessageType: messageType,
	}
	message, err = ms.messageRepo.SaveMessage(ctx, message)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return message, nil
}

func (ms *MessageService) ExchangingMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {

	conversation, err := ms.conversRepo.GetConversationById(ctx, message.ConversationId)
	if err != nil {
		return nil, err
	}

	message = &domain.Message{
		ConversationId: conversation.ConversationId,
		UserId:         message.UserId,
		MtId:			0,
		MessageText: message.MessageText,
		MessageAudio: message.MessageAudio,
		MessageType: enum.MessageTypes.User,
	}

	_, err = ms.messageRepo.SaveMessage(ctx, message)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	messages, err := ms.messageRepo.GetAllMessagesByConversationId(ctx, message.ConversationId)
	if err != nil {
		slog.Error("Error message service", "error", err)
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}


	prompt, err := ms.promptSv.GetPrompt(ctx, conversation.PromptId)
	if err != nil {
		return nil, err
	}

	modelTransaction, err :=  ms.mtSv.RequestToModel(ctx, prompt.Prompt, prompt.Prompt2, messages)
	if err != nil {
		return nil, err
	}

	generatedMessage, err := ms.mtSv.GetResponseMessage(ctx, modelTransaction)
	if err != nil {
		return nil, err
	}

	// conversation.PromptId := 123
	messageAudio := "123"

	message = &domain.Message{
		ConversationId: conversation.ConversationId,
		UserId:         message.UserId,
		MtId:      		modelTransaction.MtId,
		MessageText: generatedMessage,
		MessageAudio: messageAudio,
		MessageType: enum.MessageTypes.Ai,
	}

	message, err = ms.messageRepo.SaveMessage(ctx, message)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return message, nil
}

func (ms *MessageService) GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error) {

	conversation, err := ms.conversRepo.GetConversationById(ctx, conversationId)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return conversation, nil
}