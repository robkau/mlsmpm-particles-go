package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestGridToParticles(t *testing.T) {
	type args struct {
		particles []Particle
		gridSize  int
	}
	tests := []struct {
		name string
		args args
	}{
		{"two particles adjacent cells", args{[]Particle{
			{
				p:    mgl64.Vec2{5.2, 5.3},
				v:    mgl64.Vec2{3.30, 3},
				mass: 1.06,
				c: mgl64.Mat2{
					-0.4838, 0.01124,
					-0.0248, 0.169,
				},
			},
			{
				p:    mgl64.Vec2{6.6, 5.9},
				v:    mgl64.Vec2{1.2, -1},
				mass: 1.23,
				c: mgl64.Mat2{
					-0.4838, 0.01124,
					-0.0248, 0.169,
				},
			},
		}, 10}},
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
				v:    mgl64.Vec2{1, 1},
				mass: 0.25,
			})
			grid.Set(6, 5, Cell{
				v:    mgl64.Vec2{2, 2},
				mass: 0.25,
			})

			GridToParticles(ps, grid, 0, 0, 0)

			// check particles
			require.True(t, ps.Ps[0].Equals(
				Particle{
					p:    mgl64.Vec2{5.2497, 5.3496},
					v:    mgl64.Vec2{0.497, 0.497},
					mass: 1.06,
					c:    mgl64.Mat2{0.71, 0.71, 0.3976, 0.3976},
				},
			))
			require.True(t, ps.Ps[1].Equals(
				Particle{
					p:    mgl64.Vec2{6.692, 5.992},
					v:    mgl64.Vec2{-0.69135, 0.00704},
					mass: 1.23,
					c:    mgl64.Mat2{-0.557, -0.557, -1.47264, -1.47264},
				},
			))
		})
	}
}

func Test_weightedVelocityAndCellDistToTerm(t *testing.T) {
	type args struct {
		weightedVelocity mgl64.Vec2
		cellDist         mgl64.Vec2
	}
	tests := []struct {
		name    string
		args    args
		want    mgl64.Mat2
		wantDet float64
	}{
		{"cellDist X 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{0, 1}}, mgl64.Mat2{0, 0, 1, 1}, 0},
		{"cellDist Y 0", args{mgl64.Vec2{1, 1}, mgl64.Vec2{1, 0}}, mgl64.Mat2{1, 1, 0, 0}, 0},
		{"both", args{mgl64.Vec2{0.22, 0.77}, mgl64.Vec2{2, -1}}, mgl64.Mat2{0.22 * 2, 0.77 * 2, 0.22 * -1, 0.77 * -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := weightedVelocityAndCellDistToTerm(tt.args.weightedVelocity, tt.args.cellDist)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("weightedVelocityAndCellDistToTerm() = %v, want %v", got, tt.want)
			}
			gotDet := got.Det()
			if gotDet != tt.wantDet {
				t.Errorf("weightedVelocityAndCellDistToTerm().Det() = %v, want %v", gotDet, tt.wantDet)
			}

		})
	}
}
