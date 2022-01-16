package mpm

import "errors"

const dt = 0.1
const gravity = -0.3
const restDensity = 4
const dynamicViscosity = 0.1
const eosStiffness = 10
const eosPower = 4
const boundaryFrictionDamping = 0.001

type Grid struct {
	cells []Cell
	wh    int
}

func NewGrid(wh int) (*Grid, error) {
	if wh <= 0 {
		return nil, errors.New("width and height must be positive")
	}
	g := &Grid{
		wh:    wh,
		cells: make([]Cell, wh*wh),
	}
	return g, nil
}

func (g *Grid) Reset() {
	for i, _ := range g.cells {
		g.cells[i] = NewCell()
	}
}

func (g *Grid) Get(r, c int) Cell {
	return g.cells[r*g.wh+c]
}

func (g *Grid) GetAt(index int) Cell {
	return g.cells[index]
}

func (g *Grid) Set(r int, c int, cell Cell) {
	g.cells[r*g.wh+c] = cell
}

func (g *Grid) SetAt(index int, cell Cell) {
	g.cells[index] = cell
}

func (g *Grid) Update() {
	for i, c := range g.cells {
		if c.mass > 0 {
			// convert momentum to velocity, apply gravity
			c.v = c.v.Mul(1 / c.mass)
			c.v[1] += dt * gravity

			// boundary conditions
			x := i / g.wh
			y := i % g.wh
			if x < 2 {
				// can only stay in place or go right
				if c.v[0] < 0 {
					c.v[0] = 0
				}
				c.v[1] *= 1 - boundaryFrictionDamping
			}
			if x > g.wh-3 {
				// can only stay in place or go left
				if c.v[0] > 0 {
					c.v[0] = 0
				}
				c.v[1] *= 1 - boundaryFrictionDamping
			}
			if y < 2 {
				// can only stay in place or go up
				if c.v[1] < 0 {
					c.v[1] = 0
				}
				c.v[0] *= 1 - boundaryFrictionDamping
			}
			if y > g.wh-3 {
				// can only stay in place or go down
				if c.v[1] > 0 {
					c.v[1] = 0
				}
				c.v[0] *= 1 - boundaryFrictionDamping
			}
		}

		g.cells[i] = c
	}
}
