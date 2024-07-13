package port

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

type ModelTransactionService interface {
	RequestToModel(ctx context.Context, prompt string, prompt2 string, messages []domain.Message) (*domain.ModelTransaction, error)
	GetResponseMessage(ctx context.Context, modelTransaction *domain.ModelTransaction) (string, error)
}

type ModelTransactionRepository interface {
	CreateModelTransaction(ctx context.Context, requestPrompt []byte, responseData []byte) (*domain.ModelTransaction, error)
}