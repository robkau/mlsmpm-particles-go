package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/stretchr/testify/require"
	"math"
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
			mass: 1.06,
			c: mgl64.Mat2{
				-0.4838, 0.01124,
				-0.0248, 0.169,
			},
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
			require.True(t, FloatEqual(5.0, ps.Ps[0].p.X()))
			require.True(t, FloatEqual(5.0, ps.Ps[0].p.Y()))
			require.True(t, FloatEqual(0.0, ps.Ps[0].v.X()))
			require.True(t, FloatEqual(-1.0, ps.Ps[0].v.Y()))
			require.True(t, FloatEqual(1.06, ps.Ps[0].mass))
			require.Equal(t, mgl64.Mat2{-0.4838, 0.01124, -0.0248, 0.169}, ps.Ps[0].c)

			// local grid cells mass and velocity should be updated from particle.
			require.True(t, FloatEqual(0.265, grid.cells[44].mass))
			require.True(t, FloatEqual(0.265, grid.cells[45].mass))
			require.True(t, FloatEqual(0.0, grid.cells[46].mass))
			require.True(t, FloatEqual(0.265, grid.cells[54].mass))
			require.True(t, FloatEqual(0.265, grid.cells[55].mass))
			require.True(t, FloatEqual(0.0, grid.cells[56].mass))
			require.True(t, FloatEqual(0.0, grid.cells[64].mass))
			require.True(t, FloatEqual(0.0, grid.cells[65].mass))
			require.True(t, FloatEqual(0.0, grid.cells[66].mass))

			require.True(t, FloatEqual(0.0673895, grid.cells[44].v.X()))
			require.True(t, FloatEqual(0.0608175, grid.cells[45].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[46].v.X()))
			require.True(t, FloatEqual(-0.0608175, grid.cells[54].v.X()))
			require.True(t, FloatEqual(-0.0673895, grid.cells[55].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[56].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[64].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[65].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[66].v.X()))

			require.True(t, FloatEqual(-0.2888818, grid.cells[44].v.Y()))
			require.True(t, FloatEqual(-0.2440968, grid.cells[45].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[46].v.Y()))
			require.True(t, FloatEqual(-0.2859032, grid.cells[54].v.Y()))
			require.True(t, FloatEqual(-0.2411182, grid.cells[55].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[56].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[64].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[65].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[66].v.Y()))
		})
	}
}

func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestParticlesToGrid(t *testing.T) {
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
			mass: 1.06,
			c: mgl64.Mat2{
				-0.4838, 0.01124,
				-0.0248, 0.169,
			},
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
			// manually put some mass in the grid at particle location since previous steps not run
			grid.Set(5, 5, Cell{
				mass: 0.25,
			})

			ParticlesToGrid(ps, grid)

			// particle should not change.
			require.True(t, FloatEqual(5.0, ps.Ps[0].p.X()))
			require.True(t, FloatEqual(5.0, ps.Ps[0].p.Y()))
			require.True(t, FloatEqual(0.0, ps.Ps[0].v.X()))
			require.True(t, FloatEqual(-1.0, ps.Ps[0].v.Y()))
			require.True(t, FloatEqual(1.06, ps.Ps[0].mass))
			require.Equal(t, mgl64.Mat2{-0.4838, 0.01124, -0.0248, 0.169}, ps.Ps[0].c)

			// local grid cells momentum (as velocity) should be updated from particle.
			require.True(t, FloatEqual(0.042623872, grid.cells[44].v.X()))
			require.True(t, FloatEqual(0.044923648, grid.cells[45].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[46].v.X()))
			require.True(t, FloatEqual(-0.044923648, grid.cells[54].v.X()))
			require.True(t, FloatEqual(-0.042623872, grid.cells[55].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[56].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[64].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[65].v.X()))
			require.True(t, FloatEqual(0.0, grid.cells[66].v.X()))

			require.True(t, FloatEqual(0.097981312, grid.cells[44].v.Y()))
			require.True(t, FloatEqual(-0.100281088, grid.cells[45].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[46].v.Y()))
			require.True(t, FloatEqual(0.100281088, grid.cells[54].v.Y()))
			require.True(t, FloatEqual(-0.097981312, grid.cells[55].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[56].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[64].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[65].v.Y()))
			require.True(t, FloatEqual(0.0, grid.cells[66].v.Y()))
		})
	}
}
