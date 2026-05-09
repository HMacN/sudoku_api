package functions

import (
	"sudoku_api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_IsValid_FindsColumnDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Puzzle[2][0] = 6
	h.Puzzle[2][2] = 6
	result := IsValid(h.Puzzle)

	// assert
	assert.False(t, result, "this puzzle should not be valid")
}

func TestValidate_IsValid_FindsRowDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Puzzle[2][2] = 6
	h.Puzzle[0][2] = 6
	result := IsValid(h.Puzzle)

	// assert
	assert.False(t, result, "this puzzle should not be valid")
}

func TestValidate_IsValid_FindsBoxDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Puzzle[1][1] = 6
	h.Puzzle[2][2] = 6
	result := IsValid(h.Puzzle)

	// assert
	assert.False(t, result, "this puzzle should not be valid")
}

func TestValidate_IsValid_ReturnsTrueForSolvedBoard(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	result := IsValid(h.SolvedPuzzle)

	// assert
	assert.True(t, result, "this puzzle should be valid")
}

func TestValidate_IsValid_ReturnsTrueForPartialBoard(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.SolvedPuzzle[0][2] = 0
	h.SolvedPuzzle[1][2] = 0
	h.SolvedPuzzle[0][1] = 0
	h.SolvedPuzzle[1][1] = 0
	h.SolvedPuzzle[2][0] = 0
	h.SolvedPuzzle[2][1] = 0
	result := IsValid(h.SolvedPuzzle)

	// assert
	assert.True(t, result, "this puzzle should be valid")
}

func TestValidate_IsValid_ReturnsTrueForEmptyBoard(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	result := IsValid(h.Puzzle)

	// assert
	assert.True(t, result, "this puzzle should be valid")
}

type helper struct {
	Puzzle       models.Puzzle
	SolvedPuzzle models.Puzzle
}

func newHelper() helper {
	board := [9][9]uint8{}
	solvedBoard := [9][9]uint8{}
	solvedBoard[0] = [9]uint8{1, 2, 3, 6, 7, 8, 9, 4, 5}
	solvedBoard[1] = [9]uint8{5, 8, 4, 2, 3, 9, 7, 6, 1}
	solvedBoard[2] = [9]uint8{9, 6, 7, 1, 4, 5, 3, 2, 8}
	solvedBoard[3] = [9]uint8{3, 7, 2, 4, 6, 1, 5, 8, 9}
	solvedBoard[4] = [9]uint8{6, 9, 1, 5, 8, 3, 2, 7, 4}
	solvedBoard[5] = [9]uint8{4, 5, 8, 7, 9, 2, 6, 1, 3}
	solvedBoard[6] = [9]uint8{8, 3, 6, 9, 2, 4, 1, 5, 7}
	solvedBoard[7] = [9]uint8{2, 1, 9, 8, 5, 7, 4, 3, 6}
	solvedBoard[8] = [9]uint8{7, 4, 5, 3, 1, 6, 8, 9, 2}
	return helper{board, solvedBoard}
}
