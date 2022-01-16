package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// ParticlesToGrid transfers momentum from particles to grid
func ParticlesToGrid(ps *Particles, g *Grid) {
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

		// check surrounding 9 cells to get volume from density
		var density float64 = 0
		for gx := 0; gx < 3; gx++ {
			for gy := 0; gy < 3; gy++ {
				weight := weights[gx][0] * weights[gy][1]
				cellPosX := cellX + gx - 1
				cellPosY := cellY + gy - 1
				cellAtIdx := cellPosX*g.wh + cellPosY
				density += g.GetAt(cellAtIdx).mass * weight
			}
		}
		volume := p.mass / density

		// fluid constitutive model
		pressure := math.Max(-0.1, eosStiffness*(math.Pow(density/restDensity, eosPower)-1))
		stress := mgl64.Mat2{
			-pressure, 0,
			0, -pressure,
		}
		dudv := p.c
		strain := dudv

		trace := strain.Col(1).X() + strain.Col(0).Y()
		strain.Set(0, 1, trace)
		strain.Set(1, 0, trace)

		viscosityTerm := strain.Mul(dynamicViscosity)
		stress = stress.Add(viscosityTerm)

		eq16Term0 := stress.Mul(-volume * 4 * dt)

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
				cell := g.GetAt(cellAtIdx)
				momentum := eq16Term0.Mul(weight).Mul2x1(cellDist)
				cell.v = cell.v.Add(momentum)
				g.SetAt(cellAtIdx, cell)
			}
		}
	}
}

// UpdateCells transfers mass and velocity from particles to grid
func UpdateCells(ps *Particles, grid *Grid) {
	for _, p := range ps.Ps {
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

		for gx := 0; gx < 3; gx++ {
			for gy := 0; gy < 3; gy++ {
				weight := weights[gx][0] * weights[gy][1]

				cellPosX := cellX + gx - 1
				cellPosY := cellY + gy - 1
				cellAtIdx := cellPosX*grid.wh + cellPosY

				cellDist := mgl64.Vec2{
					float64(cellPosX) - p.p[0] + 0.5,
					float64(cellPosY) - p.p[1] + 0.5,
				}

				q := p.c.Mul2x1(cellDist)
				massContrib := weight * p.mass

				// mass and momentum update
				cell := grid.GetAt(cellAtIdx)
				cell.mass += massContrib
				cell.v = cell.v.Add(p.v.Add(q).Mul(massContrib))
				grid.SetAt(cellAtIdx, cell)
			}
		}
	}
}
