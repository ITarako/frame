package logoutHandler

import (
	"frame/internal/api/response"
	"frame/internal/lib/session"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger  *slog.Logger
	session *session.Session
}

func New(
	logger *slog.Logger,
	session *session.Session,
) *Handler {
	return &Handler{
		logger:  logger,
		session: session,
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	err := h.session.Logout(w, r)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}

	envelope := response.Envelope{
		"result": true,
	}

	err = response.WriteJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}
}
