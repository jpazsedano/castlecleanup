package main

import (
    "errors"
    "image"
    _ "image/png" // Importamos pero no usamos para que image pueda decodificar PNGs
    "bytes"
    "fmt"
    "os"
    "strings"
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: Eliminar y permitir escoger tamaño de nivel por parámetro, creando un nivel vacío
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

// Esto encapsula la lógica asociada a un nivel y la separa
// de la lógica de los menús, por ejemplo
type Level struct {
    tiles Tilemap
    enManager EntityManager
    inputController InputController
    debug bool
    editMode bool
}

// Esta función carga los assets necesarios y construye el nivel.
// En niveles grandes puede que necesite una pantalla de carga.
func MakeLevel(tilemapRes string, debug bool) (*Level, error) {
    tilemapBytes, ok := AM_RESOURCES[tilemapRes] ;
    if !ok {
        return nil, errors.New("Resource not found")
    }

    img, _, err := image.Decode(bytes.NewReader(tilemapBytes))
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Error loading image: %s", err))
    }

    tilemapImage := ebiten.NewImageFromImage(img)
    tilemap := Tilemap{tilemapImage, tileValues, 32, 19, -1}
    controller := MakeInputController()
    
    var level *Level = &Level{
        tiles: tilemap,
        inputController: controller,
        debug: debug,
        editMode: false,
    }

    return level, nil
}

func (l *Level) processEditInput() {
    // Esto captura los clicks y algunas pulsaciones de tecla si estamos en modo edición
    if !l.editMode { // Re-comprobación por si se me olvida en el código de más arriba.
        return
    }
    x, y := ebiten.CursorPosition()
    // Los clicks dependen de la interfaz que haya activada.
    if ebiten.IsKeyPressed(ebiten.KeyDigit1) {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            l.tiles.ClickTileSelection(x, y)
        }
    } else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        // Si no hay nada pulsado, pasarle el evento al modificador.
        l.tiles.ChangeTile(x, y)
    }
}

func (l *Level) Update() error {
    l.inputController.CaptureInput()

    // Activamos o desactivamos el modo edición
    if l.debug && inpututil.IsKeyJustPressed(ebiten.KeyF1) {
        l.editMode = !l.editMode
    }
    if l.debug && inpututil.IsKeyJustPressed(ebiten.KeyF2) {
        // Con F2 hacemos un guardado rápido del nivel.
        f, err := os.Create("level.dat")
        if err != nil {
            log.Println("Error saving level: ", err)
        } else {
            serErr := l.tiles.Serialize(f)
            if serErr != nil {
                log.Println("Error serializing level: ", serErr)
            } else {
                log.Println("Level saved into 'level.dat'")
            }
        }
    } else if l.debug && inpututil.IsKeyJustPressed(ebiten.KeyF3) {
        f, err := os.Open("level.dat")
        if err != nil {
            log.Println("Error loading level: ", err)
        } else {
            desErr := l.tiles.Deserialize(f)
            if desErr != nil{
                log.Println("Invalid level at level.dat", desErr)
            } else {
                log.Println("Level loaded from 'level.dat'")
            }
        }
    }

    // Si estamos en modo edición, procesamos la entrada de edición.
    if l.editMode {
        l.processEditInput()
    }

    return nil
}

func (l *Level) Draw(screen *ebiten.Image) {
    var debugMessages []string
    // Esto se hace lo primero.
    if l.debug {
        x, y := ebiten.CursorPosition()
        debugMessages = append(debugMessages, fmt.Sprintf("Cursor position: (%d, %d)", x, y))
    }
    if l.editMode {
        debugMessages = append(debugMessages, "Edit mode")
    }

    if l.editMode && ebiten.IsKeyPressed(ebiten.KeyDigit1) {
        // Tecla Tab para modificar el tilemap
        l.tiles.DrawTileSelection(screen)
    } else {
        l.tiles.DrawLayer(screen, 0)
    }

    if l.debug {
        ebitenutil.DebugPrint(screen, strings.Join(debugMessages, "\n"))
    }
}
