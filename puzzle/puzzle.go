package puzzle

type Puzzle struct {
	board [9][9]uint8
}

func (p *Puzzle) New(board [9][9]uint8) {
	p.board = board
}

func (p *Puzzle) IsValid() bool {
	for i := 0; i < 9; i++ {
		if !isValidSet(p.getRow(i)) {
			return false
		}
		if !isValidSet(p.getColumn(i)) {
			return false
		}
	}

	// TODO: Also check validity of 3x3 "boxes" of cells.

	return true
}

func (p *Puzzle) getRow(index int) [9]uint8 {
	return p.board[index]
}

func (p *Puzzle) getColumn(index int) [9]uint8 {
	col := [9]uint8{}
	for i := 0; i < 9; i++ {
		col[i] = p.board[i][index]
	}
	return col
}

func isValidSet(givenSet [9]uint8) bool {
	for i := 1; i <= 9; i++ {
		if countInstances(givenSet, uint8(i)) > 1 {
			return false
		}
	}
	return true
}

func countInstances(givenSet [9]uint8, number uint8) int {
	count := 0
	for i := 0; i < len(givenSet); i++ {
		if givenSet[i] == number {
			count++
		}
	}
	return count
}
