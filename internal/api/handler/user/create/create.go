package userCreateHandler

import (
	"errors"
	"frame/internal/api/request"
	"frame/internal/api/response"
	"frame/internal/lib/validator"
	"frame/internal/model"
	userModel "frame/internal/model/user"
	"frame/internal/service"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger      *slog.Logger
	validator   *validator.Validator
	userService service.UserService
}

func New(
	logger *slog.Logger,
	validator *validator.Validator,
	userService service.UserService,
) *Handler {
	return &Handler{
		logger:      logger,
		validator:   validator,
		userService: userService,
	}
}

type input struct {
	Name        string `json:"name" validate:"required,max=255"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=72"`
	Firstname   string `json:"firstname" validate:"max=255"`
	Middlename  string `json:"middlename" validate:"max=255"`
	Lastname    string `json:"lastname" validate:"max=255"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,e164"`
}

func (i *input) toUserModel() *userModel.User {
	return &userModel.User{
		Email: i.Email,
		Profile: userModel.Profile{
			Firstname:   i.Firstname,
			Middlename:  i.Middlename,
			Lastname:    i.Lastname,
			PhoneNumber: i.PhoneNumber,
		},
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	in := &input{}
	err := request.ReadJSON(w, r, in)
	if err != nil {
		response.BadRequest(w, r, h.logger, err)
		return
	}

	if h.validator.ValidateStruct(in); !h.validator.Valid() {
		response.FailedValidation(w, r, h.logger, h.validator.Errors)
		return
	}

	user := in.toUserModel()

	err = h.userService.Create(r.Context(), user, in.Password)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrDuplicateEmail):
			h.validator.AddError("email", "a user with this email address already exists")
			response.FailedValidation(w, r, h.logger, h.validator.Errors)
		default:
			response.ServerError(w, r, h.logger, err)
		}

		return
	}

	envelope := response.Envelope{
		"user": response.FromUserModelToUserResponse(user),
	}

	err = response.WriteJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}
}
