package presentation

type SudokuHandler struct {
	businessLayer businessLayer
}

func NewSudokuHandler() *SudokuHandler {
	return &SudokuHandler{}
}

func (h SudokuHandler) SetBusinessLayer(businessLayer businessLayer) {
	h.businessLayer = businessLayer
}

func (h SudokuHandler) RunServer() error {
	return nil
}

type businessLayer interface{}
