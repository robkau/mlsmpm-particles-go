package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"testing"
)
import "github.com/stretchr/testify/require"

func Test_GridGetSetReset(t *testing.T) {
	g, err := NewGrid(10)
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

func TestGrid_Update(t *testing.T) {
	g, err := NewGrid(10)
	require.NoError(t, err)

	// add border cell with mass and velocity
	g.Set(3, 0, Cell{
		v:    mgl64.Vec2{2.2, -2.4},
		mass: 1.17171717,
	})
	// add middle cell with mass and velocity
	g.Set(5, 5, Cell{
		v:    mgl64.Vec2{3.7333, -1.111},
		mass: 3.333,
	})

	// apply grid update
	g.Update()

	// border cell should have updated velocity and -y velocity cancelled
	require.True(t, FloatEqual(1.8757, g.Get(3, 0).v.X()))
	require.Equal(t, 0.0, g.Get(3, 0).v.Y())
	require.Equal(t, 1.17171717, g.Get(3, 0).mass)

	// middle cell should have updated velocity
	require.True(t, FloatEqual(1.120102, g.Get(5, 5).v.X()))
	require.True(t, FloatEqual(-0.36333, g.Get(5, 5).v.Y()))
	require.Equal(t, 3.333, g.Get(5, 5).mass)
}
