package userModel

import "time"

type User struct {
	ID        int32
	Email     string
	Password  password
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   Profile
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
