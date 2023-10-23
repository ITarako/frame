package user

import (
	"context"
	"database/sql"
	"frame/internal/database"
	"frame/internal/model"
	userModel "frame/internal/model/user"
	"frame/internal/repository"
)

type service struct {
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	trm                   database.TransactionManager
}

func NewService(
	userRepository repository.UserRepository,
	userProfileRepository repository.UserProfileRepository,
	trm database.TransactionManager,
) *service {
	return &service{
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		trm:                   trm,
	}
}

func (s *service) Create(ctx context.Context, user *userModel.User, plaintextPassword string) error {
	err := user.SetPassword(plaintextPassword)
	if err != nil {
		return err
	}

	return s.trm.Do(ctx, func(ctx context.Context, tx *sql.Tx) error {
		err := s.userRepository.Insert(ctx, tx, user)
		if err != nil {
			return err
		}

		return s.userProfileRepository.InsertForUser(ctx, tx, user)
	})
}

func (s *service) Get(ctx context.Context, id int32) (*userModel.User, error) {
	return s.userRepository.Get(ctx, id)
}

func (s *service) GetWithProfile(ctx context.Context, id int32) (*userModel.User, error) {
	return s.userRepository.GetWithProfile(ctx, id)
}

func (s *service) UpdateStatus(ctx context.Context, user *userModel.User) error {
	return s.userRepository.UpdateStatus(ctx, user)
}

func (s *service) Authenticate(ctx context.Context, email string, password string) (*userModel.User, error) {
	user, err := s.userRepository.GetForAuth(ctx, email)
	if err != nil {
		return nil, err
	}

	match, err := user.MatchPassword(password)
	if err != nil {
		return nil, err
	}

	if !match {
		return nil, model.ErrPasswordMismatch
	}

	return user, err
}
