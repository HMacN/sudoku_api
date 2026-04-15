package app

import (
	"net/http"
	"sudoku_api/services/logging"
	"sudoku_api/services/server"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type App struct{}

func NewApp() (*App, error) {
	return &App{}, nil
}

func (app *App) Run() error {
	logger, fxLogger := logging.NewLogger()
	fakeLoggingConstructor := func() logging.LogWrapper { return logger }
	fx.New(
		fx.WithLogger(func() fxevent.Logger { return fxLogger }),
		fx.Provide(
			server.NewServer,
			server.NewServeMux,
			fakeLoggingConstructor,
		),
		fx.Invoke(func(server *http.Server) {}),
	).Run()

	return nil
}
