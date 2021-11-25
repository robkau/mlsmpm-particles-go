package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/robkau/mlsmpm-particles/pkg/mpm"
	"image"
	"image/color"
	"log"
)

var width = 600 // window size (pixels)
var wh = 64     // simulation grid size (logical)
var scaleFactor = float64(width) / float64(wh)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func main() {
	emptyImage.Fill(color.White)

	var ps = mpm.NewParticles()
	ps.AddSquare(12, 12, wh/4)
	ps.AddRandomVelocity(13, 15)
	ps.AddRandomVelocity(-13, 0)

	grid, err := mpm.NewGrid(wh)
	if err != nil {
		panic(fmt.Errorf("create grid: %w", err))
	}

	// precomputation of particle volumes
	mpm.ParticlesToGrid(ps, grid)
	mpm.ComputeParticleVolumes(ps, grid)

	s := &state{
		ps:   ps,
		grid: grid,
	}

	ebiten.SetWindowSize(width, width)
	ebiten.SetWindowTitle("mlsmpm-particles")
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}

// state struct implements ebiten.Game interface
type state struct {
	frameCount int
	ps         *mpm.Particles
	grid       *mpm.Grid

	cursorW mgl64.Vec2 // cursor position in window
	cursorG mgl64.Vec2 // cursor position in logical mpm simulation grid
}

func (s *state) Update() error {
	s.frameCount++

	// update cursor position(s)
	cx, cy := ebiten.CursorPosition()
	s.cursorW[0] = float64(cx)
	s.cursorW[1] = float64(width) - float64(cy) // ebiten y to world y
	xG, yG := worldToGrid(s.cursorW[0], s.cursorW[1], scaleFactor)
	s.cursorG[0] = xG
	s.cursorG[1] = yG

	// reset grid
	s.grid.Reset()

	// p2g
	mpm.ParticlesToGrid(s.ps, s.grid)

	// grid velocity update
	s.grid.UpdateVelocity()

	// g2p
	mpm.GridToParticles(s.ps, s.grid)

	// apply cursor interactions
	mpm.ApplyCursorEffect(xG, yG, 4, s.ps)

	// update particle deformations
	mpm.UpdateDeformationGradients(s.ps)

	return nil
}

func (s *state) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	//RenderSprites(s.ps, screen)
	RenderVectors(s.ps, screen)

	//RenderCursor(s.cursorW, screen)
	xW, yW := gridToWorld(s.cursorG[0], s.cursorG[1], scaleFactor)
	RenderCursor(mgl64.Vec2{xW, yW}, screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (s *state) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

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
		x, y := p.Pos()
		xW, yW := gridToWorld(x, y, scaleFactor)
		yW = float64(width) - yW // world y to ebiten y
		drawEbitenLogo(img, xW, yW)
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

	// TODO: Add curves
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
