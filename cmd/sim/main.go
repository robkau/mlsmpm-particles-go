package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ps.AddSquare(4, 4, int(float64(wh)/1.2))
	ps.AddRandomVelocity(0, 3)
	ps.AddSquare(4.5, 4.5, int(float64(wh)/1.2))
	ps.AddRandomVelocity(0, 5)

	grid, err := mpm.NewGrid(wh)
	if err != nil {
		panic(fmt.Errorf("create grid: %w", err))
	}

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

	// p2g 1
	mpm.ComputeParticleVolumes(s.ps, s.grid)

	// p2g 2
	mpm.ParticlesToGrid(s.ps, s.grid)

	// grid update
	s.grid.Update()

	// g2p
	mpm.GridToParticles(s.ps, s.grid, xG, yG, 10)

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
