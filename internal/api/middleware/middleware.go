package middleware

import (
	"frame/internal/api/request"
	"frame/internal/api/response"
	"frame/internal/lib/session"
	userModel "frame/internal/model/user"
	"frame/internal/service"
	"log/slog"
	"net/http"
)

func Authenticate(s *session.Session, logger *slog.Logger, userService service.UserService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			id, err := s.GetUserID(r)
			if err != nil {
				response.ServerError(w, r, logger, err)
				return
			}

			if id == 0 {
				r = request.SetUser(r, userModel.GuestUser)
				next.ServeHTTP(w, r)
				return
			}

			user, err := userService.Get(r.Context(), id)
			if err != nil {
				response.ServerError(w, r, logger, err)
				return
			}

			r = request.SetUser(r, user)

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func RequireAuthenticated(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user := request.GetUser(r)

			if user.IsGuest() {
				response.AuthenticationRequired(w, r, logger)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
