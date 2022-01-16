package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateCells(t *testing.T) {
	type args struct {
		particles []Particle
		gridSize  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"single particle freefall", args{[]Particle{{
			p:    mgl64.Vec2{5, 5},
			v:    mgl64.Vec2{0, -1},
			mass: 1,
			c:    mgl64.Mat2{},
		}}, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := NewParticles()
			for _, p := range tt.args.particles {
				ps.AddParticle(p)
			}
			grid, err := NewGrid(tt.args.gridSize)
			require.NoError(t, err)

			UpdateCells(ps, grid)

			// particle should not change.
			require.Equal(t, 5.0, ps.Ps[0].p.X())
			require.Equal(t, 5.0, ps.Ps[0].p.Y())
			require.Equal(t, 0.0, ps.Ps[0].v.X())
			require.Equal(t, -1.0, ps.Ps[0].v.Y())

			// local grid cells should be updated from particle.
			require.Equal(t, 0.25, grid.cells[44].mass)
			require.Equal(t, 0.25, grid.cells[45].mass)
			require.Equal(t, 0.0, grid.cells[46].mass)
			require.Equal(t, 0.25, grid.cells[54].mass)
			require.Equal(t, 0.25, grid.cells[55].mass)
			require.Equal(t, 0.0, grid.cells[56].mass)
			require.Equal(t, 0.0, grid.cells[64].mass)
			require.Equal(t, 0.0, grid.cells[65].mass)
			require.Equal(t, 0.0, grid.cells[66].mass)

			require.Equal(t, 0.0, grid.cells[44].v.X())
			require.Equal(t, 0.0, grid.cells[45].v.X())
			require.Equal(t, 0.0, grid.cells[46].v.X())
			require.Equal(t, 0.0, grid.cells[54].v.X())
			require.Equal(t, 0.0, grid.cells[55].v.X())
			require.Equal(t, 0.0, grid.cells[56].v.X())
			require.Equal(t, 0.0, grid.cells[64].v.X())
			require.Equal(t, 0.0, grid.cells[65].v.X())
			require.Equal(t, 0.0, grid.cells[66].v.X())

			require.Equal(t, -0.25, grid.cells[44].v.Y())
			require.Equal(t, -0.25, grid.cells[45].v.Y())
			require.Equal(t, 0.0, grid.cells[46].v.Y())
			require.Equal(t, -0.25, grid.cells[54].v.Y())
			require.Equal(t, -0.25, grid.cells[55].v.Y())
			require.Equal(t, 0.0, grid.cells[56].v.Y())
			require.Equal(t, 0.0, grid.cells[64].v.Y())
			require.Equal(t, 0.0, grid.cells[65].v.Y())
			require.Equal(t, 0.0, grid.cells[66].v.Y())
		})
	}
}
