package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_New(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Board[2][3] = 1
	result := New(h.Board)

	// assert
	assert.Equal(t, h.Board, result.board, "this puzzle should contain the starting board")
}

func TestPuzzle_IsValid_FindsColumnDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Board[2][0] = 6
	h.Board[2][2] = 6
	h.Puzzle = New(h.Board)

	// assert
	assert.False(t, h.Puzzle.IsValid(), "this puzzle should not be valid")
}

func TestPuzzle_IsValid_FindsRowDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Board[2][2] = 6
	h.Board[0][2] = 6
	h.Puzzle = New(h.Board)

	// assert
	assert.False(t, h.Puzzle.IsValid(), "this puzzle should not be valid")
}

func TestPuzzle_IsValid_FindsBoxDuplicates(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Board[1][1] = 6
	h.Board[2][2] = 6
	h.Puzzle = New(h.Board)

	// assert
	assert.False(t, h.Puzzle.IsValid(), "this puzzle should not be valid")
}

func TestPuzzle_IsValid_ReturnsTrueForSolvedBoard(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Puzzle = New(h.SolvedBoard)

	// assert
	assert.True(t, h.Puzzle.IsValid(), "this puzzle should be valid")
}

func TestPuzzle_IsValid_ReturnsTrueForEmptyBoard(t *testing.T) {
	// arrange
	h := newHelper()

	// act
	h.Puzzle = New(h.Board)

	// assert
	assert.True(t, h.Puzzle.IsValid(), "this puzzle should be valid")
}

type helper struct {
	Puzzle      Puzzle
	Board       [9][9]uint8
	SolvedBoard [9][9]uint8
}

func newHelper() helper {
	board := [9][9]uint8{}
	puzzle := New(board)
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
	return helper{puzzle, board, solvedBoard}
}
