package service

import (
	"context"
	userModel "frame/internal/model/user"
)

type UserService interface {
	Create(context.Context, *userModel.User, string) error
	Get(context.Context, int32) (*userModel.User, error)
	GetWithProfile(context.Context, int32) (*userModel.User, error)
	UpdateStatus(context.Context, *userModel.User) error
	Authenticate(context.Context, string, string) (*userModel.User, error)
}
