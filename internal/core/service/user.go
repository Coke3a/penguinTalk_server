package service

import (
	"context"

	"github.com/Coke3a/TalkPenguin/internal/core/domain"
	"github.com/Coke3a/TalkPenguin/internal/core/port"
	"github.com/Coke3a/TalkPenguin/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
	// cache port.CacheRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
		// cache,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.PassWord)
	if err != nil {
		return nil, domain.ErrInternal
	}

	user.PassWord = hashedPassword

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return user, nil
}

func (us *UserService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	var user *domain.User

	user, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return user, nil
}

func (us *UserService) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	var users []domain.User

	users, err := us.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return users, nil
}

func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := us.repo.GetUserByID(ctx, user.UserId)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	emptyData := user.UserName == "" &&
		user.Email == "" &&
		user.PassWord == "" &&
		user.UserRank == 0 &&
		user.IncorrectLogin == 0

	sameData := existingUser.UserName == user.UserName &&
		existingUser.Email == user.Email &&
		existingUser.PassWord == user.PassWord &&
		existingUser.UserRank == user.UserRank &&
		existingUser.IncorrectLogin == user.IncorrectLogin
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	var hashedPassword string

	if user.PassWord != "" {
		hashedPassword, err = util.HashPassword(user.PassWord)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	user.PassWord = hashedPassword

	_, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return user, nil
}

func (us *UserService) DeleteUser(ctx context.Context, id uint64) error {
	_, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	return us.repo.DeleteUser(ctx, id)
}
