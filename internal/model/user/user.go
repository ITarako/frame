package userModel

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var GuestUser = &User{}

type User struct {
	ID           int32
	Email        string
	PasswordHash []byte
	Status       Status
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Profile      Profile
}

type Profile struct {
	UserID      int32
	Firstname   string
	Middlename  string
	Lastname    string
	PhoneNumber string
}

type Status int8

const (
	StatusNoActive Status = 5
	StatusActive   Status = 10
	StatusBanned   Status = 15
	StatusDeleted  Status = 20
)

func (s Status) String() string {
	var str string

	switch s {
	case StatusNoActive:
		str = "Не активен"
	case StatusActive:
		str = "Активен"
	case StatusBanned:
		str = "Заблокирован"
	case StatusDeleted:
		str = "Удален"
	}

	return str
}

func (u *User) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	u.PasswordHash = hash

	return nil
}

func (u *User) MatchPassword(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (u *User) IsGuest() bool {
	return u == GuestUser
}
