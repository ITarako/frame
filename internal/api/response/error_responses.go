package response

import (
	"fmt"
	"log/slog"
	"net/http"
)

func logError(r *http.Request, logger *slog.Logger, err error) {
	logger.Error(err.Error(),
		"request_method", r.Method,
		"request_url", r.URL.String(),
	)
}

func errorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, status int, message any) {
	env := Envelope{"error": message}

	err := WriteJSON(w, status, env, nil)
	if err != nil {
		logError(r, logger, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func FailedValidation(w http.ResponseWriter, r *http.Request, logger *slog.Logger, fields map[string]string) {
	env := Envelope{"fields": fields}

	err := WriteJSON(w, http.StatusUnprocessableEntity, env, nil)
	if err != nil {
		logError(r, logger, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	logError(r, logger, err)

	message := "the server encountered a problem and could not complete your request"
	errorResponse(w, r, logger, http.StatusInternalServerError, message)
}

func NotFound(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "the requested resource could not be found"
	errorResponse(w, r, logger, http.StatusNotFound, message)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, logger, http.StatusMethodNotAllowed, message)
}

func BadRequest(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	errorResponse(w, r, logger, http.StatusBadRequest, err.Error())
}

func EditConflict(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "unable to update the record due to an edit conflict, please try again"
	errorResponse(w, r, logger, http.StatusConflict, message)
}

func RateLimitExceeded(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "rate limit exceeded"
	errorResponse(w, r, logger, http.StatusTooManyRequests, message)
}

func InvalidCredentials(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "invalid credentials"
	errorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func InvalidAuthenticationToken(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	errorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func AuthenticationRequired(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "you must be authenticated to access this resource"
	errorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func InactiveAccount(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "you user account must be activated to access this resource"
	errorResponse(w, r, logger, http.StatusForbidden, message)
}

func NotPermitted(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "you user account doesn't have the necessary permissions to access this resource"
	errorResponse(w, r, logger, http.StatusForbidden, message)
}
