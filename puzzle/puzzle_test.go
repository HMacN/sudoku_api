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

type helper struct {
	Puzzle Puzzle
	Board  [9][9]uint8
}

func newHelper() helper {
	board := [9][9]uint8{}
	puzzle := New(board)
	return helper{puzzle, board}
}
