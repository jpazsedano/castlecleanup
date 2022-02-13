package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
)

const (
    screenWidth = 340
    screenHeight = 240
)

type Game struct {
    tiles Tilemap
}

var game Game

func init() {

}

func (g *Game) Update() {

}

func (g *Game) Draw(screen *ebiten.Image) {

}

func main() {
    ebiten.setWindowSize(screenWidth*2, screenHeight*2)
    ebiten.setWindowTitle("Castle Cleanup")

    if err := ebiten.RunGame(game) ; err != nil {
        log.Fatal(err)
    }
}
