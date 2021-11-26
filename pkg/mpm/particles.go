package mpm

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/robkau/coordinate_supplier"
	"math/rand"
)

type Particles struct {
	Ps []Particle
}

func NewParticles() *Particles {
	return &Particles{
		Ps: make([]Particle, 0),
	}
}

func (p *Particles) AddParticle(pc Particle) {
	p.Ps = append(p.Ps, pc)
}

func (p *Particles) AddSquare(x0, y0 float64, wh int) {
	cs, err := coordinate_supplier.NewCoordinateSupplier(coordinate_supplier.CoordinateSupplierOptions{
		Width:  wh,
		Height: wh,
		Order:  coordinate_supplier.Asc,
	})
	if err != nil {
		panic(fmt.Errorf("create coordinates: %w", err))
	}

	for x, y, done := cs.Next(); !done; x, y, done = cs.Next() {
		p.AddParticle(
			NewParticle(
				x0+float64(x),
				y0+float64(y),
			))
	}
}

func (p *Particles) AddRandomVelocity(upToX, upToY float64) {
	for i, pr := range p.Ps {
		pr.v = pr.v.Add(mgl64.Vec2{rand.Float64() * upToX, rand.Float64() * upToY})
		p.Ps[i] = pr
	}
}
