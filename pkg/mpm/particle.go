package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Particle struct {
	p       mgl64.Vec2 // position 1x2
	v       mgl64.Vec2 // velocity 1x2
	f       mgl64.Mat2 // deformation gradient 2x2
	mass    float64
	c       mgl64.Mat2 // affine momentum matrix 2x2
	volume0 float64    // initial volume estimate
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

func (p *Particle) Pos() (x, y float64) {
	return p.p[0], p.p[1]
}
