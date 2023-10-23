package loginHandler

import (
	"errors"
	"frame/internal/api/request"
	"frame/internal/api/response"
	"frame/internal/lib/session"
	"frame/internal/lib/validator"
	"frame/internal/model"
	"frame/internal/service"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger      *slog.Logger
	validator   *validator.Validator
	userService service.UserService
	session     *session.Session
}

func New(
	logger *slog.Logger,
	validator *validator.Validator,
	userService service.UserService,
	session *session.Session,
) *Handler {
	return &Handler{
		logger:      logger,
		validator:   validator,
		userService: userService,
		session:     session,
	}
}

type input struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
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

	user, err := h.userService.Authenticate(r.Context(), in.Email, in.Password)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			response.InvalidCredentials(w, r, h.logger)
		case errors.Is(err, model.ErrPasswordMismatch):
			response.InvalidCredentials(w, r, h.logger)
		default:
			response.ServerError(w, r, h.logger, err)
		}

		return
	}

	err = h.session.SetUserID(w, r, user.ID)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}

	envelope := response.Envelope{
		"user": response.FromUserModelToUserResponse(user),
	}

	err = response.WriteJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}
}
