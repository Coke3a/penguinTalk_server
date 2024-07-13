package port

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
)

type PromptService interface {
	GetPrompt(ctx context.Context, promptId uint64) (*domain.Prompt, error)
}

type PromptRepository interface {
	GetPromptByID(ctx context.Context, promptId uint64) (*domain.Prompt, error)
}
