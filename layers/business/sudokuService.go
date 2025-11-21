package business

type SudokuService struct {
	dataService dataService
}

func NewSudokuService() *SudokuService {
	return &SudokuService{}
}

func (h SudokuService) SetDataService(dataService dataService) {
	h.dataService = dataService
}

type dataService interface{}
