package app

import (
	http2 "net/http"
	"sudoku_api/services/http"

	"go.uber.org/fx"
)

type App struct{}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (app *App) Run() error {
	return fx.New(
		fx.Provide(
			http.NewServer,
			http.NewServeMux,
		),
		fx.Invoke(func(server *http2.Server) {}),
	).Err()
}
