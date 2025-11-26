package presentation

import (
	"fmt"
	"net/http"
	"sudoku_api/layers/presentation/handlers"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type SudokuServer struct {
	mux           *http.ServeMux
	businessLayer sudokuService
}

// TODO: Implement this lot https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/
// TODO: Look into this: https://swagger.io/docs/#

func NewSudokuServer() *SudokuServer {
	mux := http.NewServeMux()
	mux.Handle("POST /v1/solve", &handlers.SolveHandler{})
	mux.Handle("POST /v1/solve/", &handlers.SolveHandler{})

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
