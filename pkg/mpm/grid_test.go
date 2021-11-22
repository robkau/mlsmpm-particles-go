package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"testing"
)
import "github.com/stretchr/testify/require"

func Test_GridGetSetReset(t *testing.T) {
	g, err := NewGrid(1)
	require.NoError(t, err)

	// it starts with empty cells
	require.Equal(t, NewCell(), g.Get(0, 0))

	// update a cell
	addedCell := Cell{
		v:    mgl64.Vec2{2, 2},
		mass: 3.333,
	}
	g.Set(0, 0, addedCell)

	// get the cell
	require.Equal(t, addedCell, g.Get(0, 0))

	// reset back to empty
	g.Reset()
	require.Equal(t, NewCell(), g.Get(0, 0))
}
