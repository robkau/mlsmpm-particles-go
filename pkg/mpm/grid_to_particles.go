package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func weightedVelocityAndCellDistToTerm(weightedVelocity mgl64.Vec2, cellDist mgl64.Vec2) mgl64.Mat2 {
	return mgl64.Mat2FromCols(weightedVelocity.Mul(cellDist.X()), weightedVelocity.Mul(cellDist.Y()))
}

func GridToParticles(ps *Particles, g *Grid, cx, cy, mouseRadius float64) {

	for i, p := range ps.Ps {
		// reset particle velocity. we calculate it from scratch each step using the grid
		p.v[0] = 0
		p.v[1] = 0

		cellX := int(p.p[0])
		cellY := int(p.p[1])
		cellDiff := mgl64.Vec2{
			p.p[0] - float64(cellX) - 0.5,
			p.p[1] - float64(cellY) - 0.5,
		}

		// quadratic interpolation weights
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

		// constructing affine per-particle momentum matrix from APIC / MLS-MPM.
		// see APIC paper (https://web.archive.org/web/20190427165435/https://www.math.ucla.edu/~jteran/papers/JSSTS15.pdf), page 6
		// below equation 11 for clarification. this is calculating C = B * (D^-1) for APIC equation 8,
		// where B is calculated in the inner loop at (D^-1) = 4 is a constant when using quadratic interpolation functions
		b := mgl64.Mat2{}

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
				cellAtIdx := cellPosX*g.wh + cellPosY

				weightedVelocity := g.cells[cellAtIdx].v.Mul(weight)
				term := weightedVelocityAndCellDistToTerm(weightedVelocity, cellDist)

				b = b.Add(term)
				p.v = p.v.Add(weightedVelocity)
			}
		}

		p.c = b.Mul(4)

		// advect particles
		p.p = p.p.Add(p.v.Mul(dt))

		// safety clamp to ensure particles don't exit simulation domain
		p.p[0] = math.Max(p.p[0], 1)
		p.p[0] = math.Min(p.p[0], float64(g.wh-2))
		p.p[1] = math.Max(p.p[1], 1)
		p.p[1] = math.Min(p.p[1], float64(g.wh-2))

		// cursor effects
		dist := mgl64.Vec2{
			p.p[0] - cx, // x distance
			p.p[1] - cy, // y distance
		}
		if dist.Dot(dist) < mouseRadius*mouseRadius {
			normFactor := dist.Len() / mouseRadius
			normFactor = math.Pow(math.Sqrt(normFactor), 8)
			force := dist.Normalize().Mul(normFactor / 2)
			p.v = p.v.Add(force)
		}

		// boundaries
		xn := p.p.Add(p.v)
		wallMin := float64(3)
		wallMax := float64(g.wh - 4)
		if xn[0] < wallMin {
			p.v[0] += wallMin - xn[0]
			p.v[1] *= 1 - boundaryFrictionDamping
		}
		if xn[0] > wallMax {
			p.v[0] += wallMax - xn[0]
			p.v[1] *= 1 - boundaryFrictionDamping
		}
		if xn[1] < wallMin {
			p.v[1] += wallMin - xn[1]
			p.v[0] *= 1 - boundaryFrictionDamping
		}
		if xn[1] > wallMax {
			p.v[1] += wallMax - xn[1]
			p.v[0] *= 1 - boundaryFrictionDamping
		}

		ps.Ps[i] = p
	}
}
