package repository

import (
	"context"
	"database/sql"
	model "frame/internal/model/user"
)

type UserRepository interface {
	Insert(context.Context, *sql.Tx, *model.User) error
	Get(context.Context, int32) (*model.User, error)
	GetWithProfile(context.Context, int32) (*model.User, error)
	GetForAuth(context.Context, string) (*model.User, error)
	UpdateStatus(context.Context, *model.User) error
}

type UserProfileRepository interface {
	InsertForUser(context.Context, *sql.Tx, *model.User) error
}
