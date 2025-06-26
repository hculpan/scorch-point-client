package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	world        *World
	pixels       []byte
	screenWidth  int
	screenHeight int
}

func NewGame(screenWidth, screenHeight int) *Game {
	return &Game{
		world:        NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10)),
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) Update() error {
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, g.screenWidth*g.screenHeight*4)
	}
	g.world.Draw(g.pixels)
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
