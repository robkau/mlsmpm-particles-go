package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_GridToWorldToGrid(t *testing.T) {
	x0 := 1.3
	y0 := 3.9

	x1, y1 := gridToWorld(x0, y0, 3.3)
	x2, y2 := worldToGrid(x1, y1, 3.3)

	require.Equal(t, x0, x2)
	require.Equal(t, y0, y2)
}
