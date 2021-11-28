package main

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/robkau/mlsmpm-particles/pkg/mpm"
)

// RenderCursor draws cursor onto img (expected to already be in window space)
func RenderCursor(pos mgl64.Vec2, img *ebiten.Image) {
	pos[1] = float64(width) - pos[1] // world y to ebiten y
	drawEbitenLogo(img, pos[0], pos[1])

	//op := &ebiten.DrawTrianglesOptions{
	//	FillRule: ebiten.EvenOdd,
	//}
	//vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	//for i := range vs {
	//	vs[i].SrcX = 1
	//	vs[i].SrcY = 1
	//	vs[i].ColorR = 0x33 / float32(0xff)
	//	vs[i].ColorG = 0xcc / float32(0xff)
	//	vs[i].ColorB = 0x66 / float32(0xff)
	//}
	//img.DrawTriangles(vs, is, emptySubImage, op)
}

func RenderVectors(ps *mpm.Particles, img *ebiten.Image) {
	for _, p := range ps.Ps {
		pos := p.Pos()
		xW, yW := gridToWorld(pos.X(), pos.Y(), scaleFactor)
		yW = float64(width) - yW // world y to ebiten y
		drawParticle(img, xW, yW, p.Vel().Len())
	}

	//op := &ebiten.DrawTrianglesOptions{
	//	FillRule: ebiten.EvenOdd,
	//}
	//vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	//for i := range vs {
	//	vs[i].SrcX = 1
	//	vs[i].SrcY = 1
	//	vs[i].ColorR = 0x33 / float32(0xff)
	//	vs[i].ColorG = 0xcc / float32(0xff)
	//	vs[i].ColorB = 0x66 / float32(0xff)
	//}
	//img.DrawTriangles(vs, is, emptySubImage, op)
}
func drawEbitenLogo(screen *ebiten.Image, x, y float64) {
	const unit = 2

	var path vector.Path
	xf, yf := float32(x), float32(y)

	path.MoveTo(xf, yf+4*unit)
	path.LineTo(xf, yf+6*unit)
	path.LineTo(xf+2*unit, yf+6*unit)
	path.LineTo(xf+2*unit, yf+5*unit)
	path.LineTo(xf+3*unit, yf+5*unit)
	path.LineTo(xf+3*unit, yf+4*unit)
	path.LineTo(xf+4*unit, yf+4*unit)
	path.LineTo(xf+4*unit, yf+2*unit)
	path.LineTo(xf+6*unit, yf+2*unit)
	path.LineTo(xf+6*unit, yf+1*unit)
	path.LineTo(xf+5*unit, yf+1*unit)
	path.LineTo(xf+5*unit, yf)
	path.LineTo(xf+4*unit, yf)
	path.LineTo(xf+4*unit, yf+2*unit)
	path.LineTo(xf+2*unit, yf+2*unit)
	path.LineTo(xf+2*unit, yf+3*unit)
	path.LineTo(xf+unit, yf+3*unit)
	path.LineTo(xf+unit, yf+4*unit)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

func drawParticle(screen *ebiten.Image, x, y, vmag float64) {
	const unit = 1.25

	var path vector.Path
	xf, yf := float32(x), float32(y)

	path.MoveTo(xf-unit, yf+unit)
	path.LineTo(xf+unit, yf+unit)
	path.LineTo(xf+unit, yf-unit)
	path.LineTo(xf-unit, yf-unit)
	path.LineTo(xf-unit, yf+unit)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(vmag) * 0xdb / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x20 / float32(0xff)
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}
