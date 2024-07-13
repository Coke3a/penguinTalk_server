package repository

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type MessageRepository struct {
	db *postgres.DB
}

func NewMessageRepository(db *postgres.DB) *MessageRepository {
	return &MessageRepository{
		db,
	}
}

func (mr *MessageRepository) SaveMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	query := mr.db.QueryBuilder.Insert("messages").
		Columns("convers_id", "user_id", "mt_id", "msg_text", "msg_audio", "msg_type").
		Values(message.ConversationId, message.UserId, message.MtId, message.MessageText, message.MessageAudio, message.MessageType).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = mr.db.QueryRow(ctx, sql, args...).Scan(
		&message.MessageId,
		&message.ConversationId,
		&message.UserId,
		&message.MtId,
		&message.MessageText,
		&message.MessageAudio,
		&message.MessageType,
		&message.MessageDate,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return message, nil
}

func (mr *MessageRepository) GetAllMessagesByConversationId(ctx context.Context, conversationId uint64) ([]domain.Message, error) {
	var messages []domain.Message

	ordersQuery := mr.db.QueryBuilder.Select("*").
		From("messages").
		Where(sq.Eq{"convers_id": conversationId}). // Ensure this key matches your schema
		OrderBy("msg_date")

	// Convert squirrel query builder to SQL query string and arguments
	sql, args, err := ordersQuery.ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query using the connection pool or connection
	rows, err := mr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var message domain.Message
		err := rows.Scan(
			&message.MessageId,
			&message.ConversationId,
			&message.UserId,
			&message.MtId,
			&message.MessageText,
			&message.MessageAudio,
			&message.MessageType,
			&message.MessageDate,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}