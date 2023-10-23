package app

import (
	loginHandler "frame/internal/api/handler/auth/login"
	logoutHandler "frame/internal/api/handler/auth/logout"
	healthcheckHandler "frame/internal/api/handler/healthcheck"
	userActiveHandler "frame/internal/api/handler/user/active"
	userBanHandler "frame/internal/api/handler/user/ban"
	userCreateHandler "frame/internal/api/handler/user/create"
	userDeleteHandler "frame/internal/api/handler/user/delete"
	userViewHandler "frame/internal/api/handler/user/view"
	"frame/internal/api/middleware"
	"frame/internal/api/response"
	"frame/internal/lib/validator"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log/slog"
	"net/http"
)

func (app *App) routes() http.Handler {
	router := chi.NewRouter()

	router.NotFound(func(logger *slog.Logger) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			response.NotFound(w, r, logger)
		}
	}(app.container.Logger()))

	router.MethodNotAllowed(func(logger *slog.Logger) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			response.MethodNotAllowed(w, r, logger)
		}
	}(app.container.Logger()))

	router.Use(chiMiddleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9000"}, // todo from env
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))
	router.Use(middleware.Authenticate(app.container.Session(), app.container.Logger(), app.container.UserService()))

	requiredAuthenticated := middleware.RequireAuthenticated(app.container.Logger())

	router.Get("/healthcheck", healthcheckHandler.New(app.container.Logger(), app.container.Config(), app.Version).Get)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", loginHandler.New(app.container.Logger(), validator.NewValidator(), app.container.UserService(), app.container.Session()).Post)
		r.With(requiredAuthenticated).Post("/logout", logoutHandler.New(app.container.Logger(), app.container.Session()).Post)
	})

	router.Route("/users", func(r chi.Router) {
		r.Use(requiredAuthenticated)

		r.Post("/", userCreateHandler.New(app.container.Logger(), validator.NewValidator(), app.container.UserService()).Post)
		r.Get("/{id}", userViewHandler.New(app.container.Logger(), app.container.UserService()).Get)
		r.Put("/{id}/active", userActiveHandler.New(app.container.Logger(), app.container.UserService()).Put)
		r.Put("/{id}/ban", userBanHandler.New(app.container.Logger(), app.container.UserService()).Put)
		r.Delete("/{id}", userDeleteHandler.New(app.container.Logger(), app.container.UserService()).Delete)
	})

	return router

	//router.Get("/healthcheck", redirect.New(log, storage))
	//
	//router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	//
	//router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler))
	//router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	//router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", app.showMovieHandler))
	//router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", app.updateMovieHandler))
	//router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", app.deleteMovieHandler))
	//
	//router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	//router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	//
	//router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	//
	//router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())
	//
	//return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
