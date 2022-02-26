package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "log"
)

const (
    screenWidth = 340
    screenHeight = 240
)

// Consumer de prueba que captura teclas y lo muestra en pantalla.
type LoggerKeyboardConsumer struct {
    consumerId int
}

func (c *LoggerKeyboardConsumer) GetConsumerId() int{
    return c.consumerId
}

func (c *LoggerKeyboardConsumer) SetConsumerId(id int) {
    c.consumerId = id
}

func (c *LoggerKeyboardConsumer) ProcessEvent(e InputEvent) bool {
    if e.GetType() == TilesetShowKey {
        var iEv KeyEvent = e.(KeyEvent)
        if iEv.PressDown {
            log.Println("TilesetShowKey keyDown")
        } else {
            log.Println("TilesetShowKey keyUp")
        }
    }
    return true
}

type Game struct {
    scene Scene
}

type Scene interface {
    Update() error
    Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
    return g.scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
    g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth*2, screenHeight*2
}

func main() {
    ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
    ebiten.SetWindowTitle("Castle Cleanup")
    
    // TODO: Parametrize
    scene, err := MakeLevel(CASTLE_TILEMAP, true)
    if err != nil {
        log.Fatal(err)
    }

    // TODO: Hacer un constructor.
    game := &Game{scene: scene}

    if err = ebiten.RunGame(game) ; err != nil {
        log.Fatal(err)
    }
}
