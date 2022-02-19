package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "log"
    _ "embed"
    "image"
    _ "image/png" // Importamos pero no usamos para que image pueda decodificar PNGs
    "bytes"
    "fmt"
    "strings"
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
    tiles Tilemap
    inputController InputController
    debug bool
    editMode bool
}

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte
// TODO: Hace falta un gestor de recursos de algún tipo.
var tilemapImage *ebiten.Image

// Usamos init() para cargar los assets.
func init() {
    img, _, err := image.Decode(bytes.NewReader(tileRawImage))
    if err != nil {
        log.Fatal(fmt.Sprintf("Error loading image: %s", err))
    }

    tilemapImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
    g.inputController.CaptureInput()

    if g.debug && inpututil.IsKeyJustPressed(ebiten.KeyF1) {
        g.editMode = !g.editMode
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    var debugMessages []string
    // Esto se hace lo primero.
    if g.debug {
        x, y := ebiten.CursorPosition()
        debugMessages = append(debugMessages, fmt.Sprintf("Cursor position: (%d, %d)", x, y))
    }
    if g.editMode {
        debugMessages = append(debugMessages, "Edit mode")
    }

    if g.editMode && ebiten.IsKeyPressed(ebiten.KeyTab) {
        // Tecla Tab para modificar el tilemap
        g.tiles.DrawTileSelection(screen)
    } else {
        g.tiles.DrawLayer(screen, 0)
    }

    if g.debug {
        ebitenutil.DebugPrint(screen, strings.Join(debugMessages, "\n"))
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth*2, screenHeight*2
}

func main() {
    ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
    ebiten.SetWindowTitle("Castle Cleanup")
    tilemap := Tilemap{tilemapImage, tileValues, 32, 19, -1}

    controller := MakeInputController()

    // TODO: Hacer un constructor.
    game := &Game{tiles: tilemap, inputController: controller, debug: true, editMode: false}

    if err := ebiten.RunGame(game) ; err != nil {
        log.Fatal(err)
    }
}
