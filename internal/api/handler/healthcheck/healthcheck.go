package healthcheckHandler

import (
	"frame/internal/api/response"
	"frame/internal/config"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger  *slog.Logger
	config  *config.Config
	version string
}

func New(
	logger *slog.Logger,
	config *config.Config,
	version string,
) *Handler {
	return &Handler{
		logger:  logger,
		config:  config,
		version: version,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	envelope := response.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": h.config.Project.Env,
			"version":     h.version,
		},
	}

	err := response.WriteJSON(w, http.StatusOK, envelope, nil)
	if err != nil {
		response.ServerError(w, r, h.logger, err)
	}
}
