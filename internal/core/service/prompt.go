package service

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/Coke3a/TalkPenguin/internal/core/port"
)

type PromptService struct {
	repo port.PromptRepository
	// cache port.CacheRepository
}

func NewPromptService(repo port.PromptRepository) *PromptService {
	return &PromptService{
		repo,
		// cache,
	}
}

func (ps *PromptService) GetPrompt(ctx context.Context, promptId uint64) (*domain.Prompt, error) {

	var prompt *domain.Prompt
	prompt, err := ps.repo.GetPromptByID(ctx, promptId)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return prompt, nil
}
