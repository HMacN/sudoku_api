package presentation

import (
	"context"
	"fmt"
	"net/http"
	"sudoku_api/layers/presentation/handlers"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// TODO: https://uber-go.github.io/fx/get-started/minimal.html

type SudokuServer struct {
	mux           *http.ServeMux
	businessLayer sudokuService
}

func NewSudokuServer(lc fx.Lifecycle) *SudokuServer {
	mux := http.NewServeMux()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			mux.Handle("GET /healthy", &handlers.Healthy{})
			mux.Handle("POST /v1/solve", &handlers.Solve{})
			mux.Handle("POST /v1/solve/", &handlers.Solve{})
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return &SudokuServer{
		mux: mux,
	}
}

func (h SudokuServer) SetSudokuService(service sudokuService) {
	h.businessLayer = service
}

func (h SudokuServer) Run() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), h.mux)
	if err != nil {
		return errors.Wrap(err, "encountered error while starting server")
	}
	return nil
}

type sudokuService interface{}
