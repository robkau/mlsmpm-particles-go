package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Particle struct {
	p    mgl64.Vec2 // position 1x2
	v    mgl64.Vec2 // velocity 1x2
	mass float64
	c    mgl64.Mat2 // affine momentum matrix 2x2
}

func NewParticle(x, y float64) Particle {
	return Particle{
		p: mgl64.Vec2{
			x,
			y,
		},
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

func (p *Particle) Equals(o Particle) bool {
	return Vec2Equal(p.p, o.p) &&
		Vec2Equal(p.v, o.v) &&
		Mat2Equal(p.c, o.c) &&
		p.mass == o.mass
}

func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func Mat2Equal(a, b mgl64.Mat2) bool {
	return Vec2Equal(a.Row(0), b.Row(0)) &&
		Vec2Equal(a.Row(1), b.Row(1))
}

func Vec2Equal(a, b mgl64.Vec2) bool {
	return FloatEqual(a.X(), b.X()) &&
		FloatEqual(a.Y(), b.Y())
}
