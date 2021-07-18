package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLit(t *testing.T) {
	s := New(2)

	assert.Equal(t, 1, s.Lit(0, 0, 1))
	assert.Equal(t, 4, s.Lit(0, 0, 4))
	assert.Equal(t, 64, s.Lit(3, 3, 4))
}
