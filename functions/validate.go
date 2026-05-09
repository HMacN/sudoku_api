package functions

import "sudoku_api/models"

func IsValid(puzzle models.Puzzle) bool {
	for i := 0; i < 9; i++ {
		if !isValidSet(getRow(puzzle, i)) {
			return false
		}
		if !isValidSet(getColumn(puzzle, i)) {
			return false
		}
	}

	for i := 0; i <= 6; i = i + 3 {
		for j := 0; j <= 6; j = j + 3 {
			if !isValidSet(getBox(puzzle, i, j)) {
				return false
			}
		}
	}

	return true
}

func getColumn(puzzle models.Puzzle, index int) [9]uint8 {
	return puzzle[index]
}

func getRow(puzzle models.Puzzle, index int) [9]uint8 {
	col := [9]uint8{}
	for i := 0; i < 9; i++ {
		col[i] = puzzle[i][index]
	}
	return col
}

func getBox(puzzle models.Puzzle, topLeftCellRow int, topLeftCellCol int) [9]uint8 {
	index := 0
	box := [9]uint8{}
	for i := topLeftCellRow; i < 3; i++ {
		for j := topLeftCellCol; j < 3; j++ {
			box[index] = puzzle[i][j]
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
