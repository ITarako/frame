package response

import (
	userModel "frame/internal/model/user"
	"time"
)

type User struct {
	ID          int32      `json:"id"`
	Email       string     `json:"email"`
	Status      UserStatus `json:"status"`
	Firstname   string     `json:"firstname"`
	Middlename  string     `json:"middlename"`
	Lastname    string     `json:"lastname"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type UserStatus struct {
	ID   int8   `json:"id"`
	Name string `json:"name"`
}

func FromUserModelToUserResponse(model *userModel.User) User {
	return User{
		ID:    model.ID,
		Email: model.Email,
		Status: UserStatus{
			ID:   int8(model.Status),
			Name: model.Status.String(),
		},
		Firstname:   model.Profile.Firstname,
		Middlename:  model.Profile.Middlename,
		Lastname:    model.Profile.Lastname,
		PhoneNumber: model.Profile.PhoneNumber,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
