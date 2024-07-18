package repository

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type PromptRepository struct {
	db *postgres.DB
}

func NewPromptRepository(db *postgres.DB) *PromptRepository {
	return &PromptRepository{
		db,
	}
}

func (pr *PromptRepository) GetPromptByID(ctx context.Context, promptId uint64) (*domain.Prompt, error) {
	var prompt domain.Prompt

	query := pr.db.QueryBuilder.Select("*").
		From("prompt").
		Where(sq.Eq{"prompt_id": promptId}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = pr.db.QueryRow(ctx, sql, args...).Scan(
		&prompt.PromptId,
		&prompt.ConversationTopicId,
		&prompt.PromptLangId,
		&prompt.Prompt,
		&prompt.Prompt2,
		&prompt.AiRole,
		&prompt.UserRole,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &prompt, nil
}
