package main

import (
    "github.com/hajimehoshi/ebiten/v2"
//    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "log"
    _ "embed"
    "image"
    _ "image/png"
    "bytes"
    "fmt"
)

const (
    screenWidth = 340
    screenHeight = 240
)

var tileValues = [][][]int{
    // Capa 1
    {
        {40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
        {40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
        {40, 40, 26, 59, 59, 59, 59, 27, 40, 40},
        {40, 40, 41,134,135,135,136, 39, 40, 40},
        {40, 40, 41,153,154,154,155, 39, 40, 40},
        {40, 40, 41,153,154,154,155, 39, 40, 40},
        {40, 40, 41,172,173,173,174, 39, 40, 40},
        {40, 40, 45, 21, 21, 21, 21, 46, 40, 40},
        {40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
        {40, 40, 40, 40, 40, 40, 40, 40, 40, 40},
    },
}

type Game struct {
    tiles Tilemap
}

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte
// TODO: Esto no debería ser una variable global.
var tilemapImage *ebiten.Image

func init() {
    img, _, err := image.Decode(bytes.NewReader(tileRawImage))
    if err != nil {
        log.Fatal(fmt.Sprintf("Error loading image: %s", err))
    }

    tilemapImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.tiles.DrawLayer(screen, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth*2, screenHeight*2
}

func main() {
    ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
    ebiten.SetWindowTitle("Castle Cleanup")
    //tilemap := MakeEmptyTilemap(false, tilemapImage, 1, 10, 10, 32)
    tilemap := Tilemap{false, tilemapImage, tileValues, 32, 19}

    game := &Game{tiles: tilemap}

    if err := ebiten.RunGame(game) ; err != nil {
        log.Fatal(err)
    }
}
