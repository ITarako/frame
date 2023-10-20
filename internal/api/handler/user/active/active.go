package userActiveHandler

import (
	"errors"
	"frame/internal/api/request"
	"frame/internal/api/response"
	"frame/internal/model"
	userModel "frame/internal/model/user"
	"frame/internal/service"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger      *slog.Logger
	userService service.UserService
}

func New(
	logger *slog.Logger,
	userService service.UserService,
) *Handler {
	return &Handler{
		logger:      logger,
		userService: userService,
	}
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	id, err := request.ReadID32Param(r)
	if err != nil {
		response.NotFound(w, r, h.logger)
		return
	}

	user, err := h.userService.GetWithProfile(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			response.NotFound(w, r, h.logger)
		default:
			response.ServerError(w, r, h.logger, err)
		}

		return
	}

	user.Status = userModel.StatusActive

	err = h.userService.UpdateStatus(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			response.NotFound(w, r, h.logger)
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
