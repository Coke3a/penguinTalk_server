package repository

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type ModelTransactionRepository struct {
	db *postgres.DB
}

func NewModelTransactionRepository(db *postgres.DB) *ModelTransactionRepository {
	return &ModelTransactionRepository{
		db,
	}
}

func (mr *ModelTransactionRepository) CreateModelTransaction(ctx context.Context, requestPrompt []byte, responseData []byte) (*domain.ModelTransaction, error) {
	var modelTransaction domain.ModelTransaction

	query := mr.db.QueryBuilder.Insert("model_transaction").
		Columns("request_prompt", "response_data").
		Values(requestPrompt, responseData).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = mr.db.QueryRow(ctx, sql, args...).Scan(
		&modelTransaction.MtId,
		&modelTransaction.RequestPrompt,
		&modelTransaction.ResponseData,
		&modelTransaction.TransactionDate,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &modelTransaction, nil
}