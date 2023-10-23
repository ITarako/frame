package app

import (
	"context"
	"database/sql"
	"frame/internal/lib/vcs"
	"net/http"
	"sync"
)

type App struct {
	backgroundWG sync.WaitGroup
	container    *container
	server       *http.Server
	Version      string
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	app.Version = vcs.Version()

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	return app.runServer()
}

func (app *App) DB() *sql.DB {
	return app.container.DB()
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.initContainer,
		app.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initContainer(_ context.Context) error {
	app.container = newContainer()

	return nil
}
