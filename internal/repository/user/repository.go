package user

import (
	"context"
	"database/sql"
	"errors"
	"frame/internal/model"
	userModel "frame/internal/model/user"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Insert(ctx context.Context, tx *sql.Tx, user *userModel.User) error {
	query := `
		INSERT INTO "quartz_user"."user" (email, password_hash, status)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`
	args := []any{user.Email, user.Password.GetHash(), userModel.StatusNoActive}

	err := tx.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return model.ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (r *repository) Get(ctx context.Context, id int32) (*userModel.User, error) {
	if id < 1 {
		return nil, model.ErrRecordNotFound
	}

	query := `
		SELECT id, email, status, created_at, updated_at
		FROM "quartz_user"."user"
		WHERE id = $1`

	var user userModel.User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, model.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *repository) GetWithProfile(ctx context.Context, id int32) (*userModel.User, error) {
	if id < 1 {
		return nil, model.ErrRecordNotFound
	}

	query := `
		SELECT id, email, status, created_at, updated_at, firstname, middlename, lastname, phone_number
		FROM "quartz_user"."user"
		LEFT JOIN quartz_user.user_profile up on "user".id = up.user_id
		WHERE id = $1`

	var user userModel.User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Profile.Firstname,
		&user.Profile.Middlename,
		&user.Profile.Lastname,
		&user.Profile.PhoneNumber,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, model.ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *repository) UpdateStatus(ctx context.Context, user *userModel.User) error {
	query := `
		UPDATE "quartz_user"."user"
		SET status = $1
		WHERE id = $2`

	result, err := r.db.ExecContext(ctx, query, user.Status, user.ID)
	if err != nil {
		return nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return model.ErrRecordNotFound
	}

	return nil
}
