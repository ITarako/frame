package userProfile

import (
	"context"
	"database/sql"
	userModel "frame/internal/model/user"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) InsertForUser(ctx context.Context, tx *sql.Tx, user *userModel.User) error {
	query := `
		INSERT INTO "quartz_user"."user_profile" (user_id, firstname, middlename, lastname, phone_number)
		VALUES ($1, $2, $3, $4)`
	args := []any{user.ID, user.Profile.Firstname, user.Profile.Middlename, user.Profile.Lastname, user.Profile.PhoneNumber}

	_, err := tx.ExecContext(ctx, query, args...)
	return err
}
