package repository

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/adapter/storage/postgres"
	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := ur.db.QueryBuilder.Insert("users").
		Columns("username", "password", "email", "users_rank").
		Values(user.UserName, user.PassWord, user.Email, user.UserRank).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.UserId,
		&user.UserName,
		&user.PassWord,
		&user.Email,
		&user.UserRank,
		&user.LastLogin,
		&user.IncorrectLogin,
		&user.CreateDate,
	)

	// Debugging: Check if there was an error during Scan
	if err != nil {
		if errCode := ur.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	var user domain.User

	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where(sq.Eq{"user_id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.UserId,
		&user.UserName,
		&user.PassWord,
		&user.Email,
		&user.UserRank,
		&user.LastLogin,
		&user.IncorrectLogin,
		&user.CreateDate,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := ur.db.QueryBuilder.Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.UserId,
		&user.UserName,
		&user.PassWord,
		&user.Email,
		&user.UserRank,
		&user.LastLogin,
		&user.IncorrectLogin,
		&user.CreateDate,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	var users []domain.User

	ordersQuery := ur.db.QueryBuilder.Select("*").
		From("users").
		OrderBy("user_id").
		Limit(limit).
		Offset((skip - 1) * limit)

	// Convert squirrel query builder to SQL query string and arguments
	sql, args, err := ordersQuery.ToSql()
	if err != nil {
		return nil, err
	}

	// Execute the query using the connection pool or connection
	rows, err := ur.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.UserId,
			&user.UserName,
			&user.PassWord,
			&user.Email,
			&user.UserRank,
			&user.LastLogin,
			&user.IncorrectLogin,
			&user.CreateDate,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	query := ur.db.QueryBuilder.Update("users").
		Set("user_name", sq.Expr("COALESCE(?, user_name)", user.UserName)).
		Set("password", sq.Expr("COALESCE(?, password)", user.PassWord)).
		Set("email", sq.Expr("COALESCE(?, email)", user.Email)).
		Set("users_rank", sq.Expr("COALESCE(?, users_rank)", user.UserRank)).
		Set("incorrect_login", sq.Expr("COALESCE(?, incorrect_login)", user.IncorrectLogin)).
		Set("last_login", sq.Expr("COALESCE(?, last_login)", user.LastLogin)).
		Where(sq.Eq{"user_id": user.UserId}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = ur.db.QueryRow(ctx, sql, args...).Scan(
		&user.UserId,
		&user.UserName,
		&user.PassWord,
		&user.Email,
		&user.UserRank,
		&user.LastLogin,
		&user.IncorrectLogin,
		&user.CreateDate,
	)

	// Debugging: Check if there was an error during Scan
	if err != nil {
		if errCode := ur.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uint64) error {
	query := ur.db.QueryBuilder.Delete("users").
		Where(sq.Eq{"user_id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = ur.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
