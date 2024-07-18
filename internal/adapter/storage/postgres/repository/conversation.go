package repository

import (
	"context"
	"fmt"

	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type ConversationRepository struct {
	db *postgres.DB
}

func NewConversationRepository(db *postgres.DB) *ConversationRepository {
	return &ConversationRepository{
		db,
	}
}

func (cr *ConversationRepository) CreateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {
	query := cr.db.QueryBuilder.Insert("conversation").
		Columns("user_id", "prompt_id").
		Values(conversation.UserId, conversation.PromptId).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&conversation.ConversationId,
		&conversation.UserId,
		&conversation.PromptId,
		&conversation.ConversationStart,
		&conversation.ConversationEnd,
	)

	// Debugging: Check if there was an error during Scan
	if err != nil {
		fmt.Println("Error during Scan:", err)
		return nil, err
	}

	return conversation, nil
}

func (cr *ConversationRepository) GetConversationById(ctx context.Context, conversationId uint64) (*domain.Conversation, error) {
	var conversation domain.Conversation

	query := cr.db.QueryBuilder.Select("*").
		From("conversation").
		Where(sq.Eq{"convers_id": conversationId}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&conversation.ConversationId,
		&conversation.UserId,
		&conversation.PromptId,
		&conversation.ConversationStart,
		&conversation.ConversationEnd,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &conversation, nil
}
