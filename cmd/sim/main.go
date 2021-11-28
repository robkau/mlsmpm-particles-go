package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/robkau/mlsmpm-particles/pkg/mpm"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"math/rand"
	"os"
)

// todo spawners shooting in particles - 4 directions to make a whirlpool
// todo particles that expire after X turns
// todo 1: multi phase simulation (per-particle properties, maybe different constitutive models too)
// todo 2: record to txt data files with particle positions
// todo 4: optimizations - cache, pool, concurrent processing

var width = 800 // window size (pixels)
var wh = 128    // simulation grid size (logical)
var scaleFactor = float64(width) / float64(wh)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func main() {
	emptyImage.Fill(color.White)

	var ps = mpm.NewParticles()
	ps.AddSquare(4, 14, int(float64(wh)/4))
	//ps.AddSquare(60, 14, int(float64(wh)/4), mpm.Fluid)
	//ps.AddSquare(4, 4, int(float64(wh)/2))
	//ps.AddRandomVelocity(0, 3)

	grid, err := mpm.NewGrid(wh)
	if err != nil {
		panic(fmt.Errorf("create grid: %w", err))
	}

	s := &state{
		ps:   ps,
		grid: grid,
		opts: opts{
			gifFrames:        0,
			stepsPerGifFrame: 1,
		},
	}

	if s.opts.gifFrames > 0 {
		s.gif = &gif.GIF{
			Image:     make([]*image.Paletted, s.opts.gifFrames),
			Delay:     make([]int, s.opts.gifFrames),
			LoopCount: -1,
		}
		outGif, err := os.Create("output.gif")
		if err != nil {
			log.Fatal(err)
		}
		s.outGif = outGif
		defer outGif.Close()
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

	opts opts

	gif        *gif.GIF
	outGif     *os.File
	atGifFrame int
}

type opts struct {
	gifFrames        int
	stepsPerGifFrame int
}

func (s *state) Update() error {
	s.frameCount++

	if s.frameCount%1 == 0 {
		p := mpm.NewParticleV(15, 115, mgl64.Vec2{3.5 + rand.Float64(), 0.5 + rand.Float64()})
		s.ps.AddParticle(p)
	}

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
	mpm.UpdateCells(s.ps, s.grid)

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

	if s.opts.gifFrames > 0 && s.atGifFrame < s.opts.gifFrames {
		// making a gif.
		if s.opts.stepsPerGifFrame <= 1 || s.frameCount%s.opts.stepsPerGifFrame == 0 {
			// saving the current simulation step as a gif frame.
			img := image.NewPaletted(screen.Bounds(), palette.Plan9)
			draw.Draw(img, img.Bounds(), screen, screen.Bounds().Min, draw.Over)
			s.gif.Image[s.atGifFrame] = img
			s.gif.Delay[s.atGifFrame] = 2
			s.atGifFrame++

			if s.atGifFrame >= s.opts.gifFrames {
				if err := gif.EncodeAll(s.outGif, s.gif); err != nil {
					panic(fmt.Sprintf("failed encode gif: %v", err))
				}
			}
		}
	}

	//RenderCursor(s.cursorW, screen)
	xW, yW := gridToWorld(s.cursorG[0], s.cursorG[1], scaleFactor)
	RenderCursor(mgl64.Vec2{xW, yW}, screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (s *state) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
