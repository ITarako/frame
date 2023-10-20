package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *App) initServer(_ context.Context) error {

	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.container.Config().Server.Port),
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(app.container.Logger().Handler(), slog.LevelError),
		IdleTimeout:  time.Duration(app.container.Config().Server.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(app.container.Config().Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.container.Config().Server.WriteTimeout) * time.Second,
	}

	return nil
}

func (app *App) runServer() error {
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.container.Logger().Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		err := app.server.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		app.container.Logger().Info("completing background tasks", "addr", app.server.Addr)

		app.backgroundWG.Wait()
		shutdownError <- nil
	}()

	app.container.Logger().Info("starting server",
		"addr", app.server.Addr,
		"env", app.container.Config().Project.Env,
	)

	err := app.server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.container.Logger().Info("stopped server", "addr", app.server.Addr)

	return nil
}
