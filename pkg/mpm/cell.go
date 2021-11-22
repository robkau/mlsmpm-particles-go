package mpm

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Cell struct {
	v    mgl64.Vec2 // velocity 1x2
	mass float64
}

func NewCell() Cell {
	return Cell{}
}
