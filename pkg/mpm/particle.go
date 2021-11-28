package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Particle struct {
	p    mgl64.Vec2 // position 1x2
	v    mgl64.Vec2 // velocity 1x2
	f    mgl64.Mat2 // deformation gradient 2x2
	mass float64
	c    mgl64.Mat2 // affine momentum matrix 2x2
}

func NewParticle(x, y float64) Particle {
	return Particle{
		p: mgl64.Vec2{
			x,
			y,
		},
		f:    mgl64.Ident2(),
		mass: 1,
	}
}

func NewParticleV(x, y float64, v mgl64.Vec2) Particle {
	p := NewParticle(x, y)
	p.v = v
	return p
}

func (p *Particle) Pos() mgl64.Vec2 {
	return p.p
}

func (p *Particle) Vel() mgl64.Vec2 {
	return p.v
}
