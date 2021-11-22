package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// ParticlesToGrid transfers data from particles to grid
func ParticlesToGrid(ps *Particles, g Grid) {

	for _, p := range ps.Ps {
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
				massContrib := weight * p.mass

				// converting 2D index to 1D
				cellAtIdx := cellPosX*g.wh + cellPosY
				cellAt := g.GetAt(cellAtIdx)
				cellAt.mass += massContrib
				cellAt.v = cellAt.v.Add(p.v.Add(q).Mul(massContrib))
				g.SetAt(cellAtIdx, cellAt)

				// // note: currently "cell.v" refers to MOMENTUM, not velocity!
				// // this gets corrected in the UpdateGrid step below.
			}
		}
	}
}
