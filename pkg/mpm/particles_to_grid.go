package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Lamé parameters for stress-strain relationship
const elasticLambda = 0.1
const elasticMu = 2000

// ParticlesToGrid transfers data from particles to grid
func ParticlesToGrid(ps *Particles, g *Grid) {

	for _, p := range ps.Ps {
		stress := mgl64.Mat2{}
		j := p.f.Det()
		volume := p.volume0 * j

		f_T := p.f.Transpose()
		f_inv_T := f_T.Inv()
		f_minus_f_inv_T := p.f.Sub(f_inv_T)

		p_term_0 := f_minus_f_inv_T.Mul(elasticMu)
		p_term_1 := f_inv_T.Mul(math.Log(j) * elasticLambda)
		p_combined := p_term_0.Add(p_term_1)

		stress = p_combined.Mul2(f_T).Mul(1.0 / j)

		eq16Term0 := stress.Mul(-volume * 4 * dt)

		cellX := int(p.p[0])
		cellY := int(p.p[1])
		cellDiff := mgl64.Vec2{
			p.p[0] - float64(cellX) - 0.5,
			p.p[1] - float64(cellY) - 0.5,
		}

		weights := make([]mgl64.Vec2, 3)
		weights[0] = mgl64.Vec2{
			0.5 * math.Pow(0.5-cellDiff[0], 2),
			0.5 * math.Pow(0.5-cellDiff[1], 2),
		}
		weights[1] = mgl64.Vec2{
			0.75 - math.Pow(cellDiff[0], 2),
			0.75 - math.Pow(cellDiff[1], 2),
		}
		weights[2] = mgl64.Vec2{
			0.5 * math.Pow(0.5+cellDiff[0], 2),
			0.5 * math.Pow(0.5+cellDiff[1], 2),
		}

		// for all surrounding 9 cells
		for gx := 0; gx < 3; gx++ {
			for gy := 0; gy < 3; gy++ {
				weight := weights[gx][0] * weights[gy][1]
				cellPosX := cellX + gx - 1
				cellPosY := cellY + gy - 1

				cellDist := mgl64.Vec2{
					float64(cellPosX) - p.p[0] + 0.5,
					float64(cellPosY) - p.p[1] + 0.5,
				}

				q := p.c.Mul2x1(cellDist)

				// scatter mass and momentum to the grid
				cellAtIdx := cellPosX*g.wh + cellPosY
				cell := g.GetAt(cellAtIdx)
				weightedMass := weight * p.mass
				cell.mass += weightedMass
				cell.v = cell.v.Add(p.v.Add(q).Mul(weightedMass))

				momentum := eq16Term0.Mul(weight).Mul2x1(cellDist)
				cell.v = cell.v.Add(momentum)

				g.SetAt(cellAtIdx, cell)
			}
		}
	}
}

func ComputeParticleVolumes(ps *Particles, grid *Grid) {
	for i, p := range ps.Ps {
		// quadratic interpolation weights
		cellX := int(p.p[0])
		cellY := int(p.p[1])
		cellDiff := mgl64.Vec2{
			p.p[0] - float64(cellX) - 0.5,
			p.p[1] - float64(cellY) - 0.5,
		}

		weights := make([]mgl64.Vec2, 3)
		weights[0] = mgl64.Vec2{
			0.5 * math.Pow(0.5-cellDiff[0], 2),
			0.5 * math.Pow(0.5-cellDiff[1], 2),
		}
		weights[1] = mgl64.Vec2{
			0.75 - math.Pow(cellDiff[0], 2),
			0.75 - math.Pow(cellDiff[1], 2),
		}
		weights[2] = mgl64.Vec2{
			0.5 * math.Pow(0.5+cellDiff[0], 2),
			0.5 * math.Pow(0.5+cellDiff[1], 2),
		}

		// for all surrounding 9 cells
		var density float64 = 0

		for gx := 0; gx < 3; gx++ {
			for gy := 0; gy < 3; gy++ {
				weight := weights[gx][0] * weights[gy][1]
				// map 2D to 1D index in grid
				cellIndex := (cellX+(gx-1))*grid.wh + (cellY + gy - 1)
				density += grid.GetAt(cellIndex).mass * weight
			}
		}
		// per-particle volume estimate has now been computed
		p.volume0 = p.mass / density
		ps.Ps[i] = p
	}
}
