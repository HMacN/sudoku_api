package puzzle

type Puzzle struct {
	board [9][9]uint8
}

func New(board [9][9]uint8) Puzzle {
	return Puzzle{board}
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

	if !isValidSet(p.getBox(0, 0)) {
		return false
	}
	if !isValidSet(p.getBox(3, 0)) {
		return false
	}
	if !isValidSet(p.getBox(6, 0)) {
		return false
	}

	for i := 0; i <= 6; i = i + 3 {
		for j := 0; j <= 6; j = j + 3 {
			if !isValidSet(p.getBox(i, j)) {
				return false
			}
		}
	}

	return true
}

func (p *Puzzle) getColumn(index int) [9]uint8 {
	return p.board[index]
}

func (p *Puzzle) getRow(index int) [9]uint8 {
	col := [9]uint8{}
	for i := 0; i < 9; i++ {
		col[i] = p.board[i][index]
	}
	return col
}

func (p *Puzzle) getBox(topLeftCellRow int, topLeftCellCol int) [9]uint8 {
	index := 0
	box := [9]uint8{}
	for i := topLeftCellRow; i < 3; i++ {
		for j := topLeftCellCol; j < 3; j++ {
			box[index] = p.board[i][j]
			index++
		}
	}
	return box
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
