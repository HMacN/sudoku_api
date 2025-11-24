package app

import (
	"sudoku_api/layers/business"
	"sudoku_api/layers/persistence"
	"sudoku_api/layers/presentation"

	"github.com/pkg/errors"
)

type App struct {
	presentationLayer *presentation.SudokuServer
	businessLayer     *business.SudokuService
	persistenceLayer  *persistence.DataService
}

func NewApp() (*App, error) {
	persLayer := persistence.NewDataService()

	busLayer := business.NewSudokuService()
	busLayer.SetDataService(persLayer)

	presLayer := presentation.NewSudokuServer()
	presLayer.SetSudokuService(busLayer)

	return &App{
		presentationLayer: presLayer,
		businessLayer:     busLayer,
		persistenceLayer:  persLayer,
	}, nil
}

func (app *App) Run() error {
	err := app.presentationLayer.Run()
	if err != nil {
		return errors.Wrap(err, "encountered error while running the presentation layer")
	}

	return nil
}
