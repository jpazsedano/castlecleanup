package main

import (
    "github.com/hajimehoshi/ebiten/v2"
//    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
)

const (
    screenWidth = 340
    screenHeight = 240
)

type Game struct {
    tiles Tilemap
}

func init() {

}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
    ebiten.SetWindowTitle("Castle Cleanup")
    game := &Game{}

    if err := ebiten.RunGame(game) ; err != nil {
        log.Fatal(err)
    }
}
