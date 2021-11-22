package mpm

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func ApplyCursorEffect(cx, cy, mouseRadius float64, ps *Particles) {
	for i, p := range ps.Ps {
		dist := mgl64.Vec2{
			p.p[0] - cx, // x distance
			p.p[1] - cy, // y distance
		}
		if dist.Dot(dist) < mouseRadius*mouseRadius {
			normFactor := dist.Len() / mouseRadius
			normFactor = math.Pow(math.Sqrt(normFactor), 8)
			force := dist.Normalize().Mul(normFactor / 2)
			fmt.Printf("adding force: %v\n", force)
			p.v = p.v.Add(force)
		}
		ps.Ps[i] = p
	}
}
