package app

import (
	"net/http"
	"sudoku_api/services/server"

	"go.uber.org/fx"
)

type App struct{}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (app *App) Run() error {
	fx.New(
		fx.Provide(
			server.NewServer,
			server.NewServeMux,
		),
		fx.Invoke(func(server *http.Server) {}),
	).Run()

	return nil
}
