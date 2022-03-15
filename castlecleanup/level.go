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

    "github.com/jpazsedano/castlecleanup/gameutils"
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

type EditSelection int

const (
    EDIT_MAP EditSelection = iota
    EDIT_ENTITIES
)

// Esto encapsula la lógica asociada a un nivel y la separa
// de la lógica de los menús, por ejemplo
type Level struct {
    tiles Tilemap
    enManager EntityManager
    inputController InputController
    debug bool
    editMode bool

    editSelection EditSelection
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
    entityManager, _ := gameutils.MakeEntityManager()
    
    var level *Level = &Level{
        tiles: tilemap,
        inputController: controller,
        debug: debug,
        editMode: false,
        enManager: *entityManager,
    }

    return level, nil
}

func (l *Level) switchEditMode() {
    switch {
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit1):
        log.Println("'Edit Map' mode selected")
        l.editSelection = EDIT_MAP
    case inpututil.IsKeyJustPressed(ebiten.KeyDigit2):
        log.Println("'Edit Entities' mode selected")
        l.editSelection = EDIT_ENTITIES
    }
}

// Esta función se encarga de gestionar toda la lógica perteneciente al editor.
// capturando los eventos y cambiando el estado interno del nivel (de las propiedades
// de edición) y de llamar a los métodos de edición de los subobjetos.
func (l *Level) processEditInput() {
    // Esto captura los clicks y algunas pulsaciones de tecla si estamos en modo edición
    if !l.editMode { // Re-comprobación por si se me olvida en el código de más arriba.
        return
    }
    // TODO: Cuando haya scroll habrá que transformar estas coordenadas.
    x, y := ebiten.CursorPosition()

    // Primero gestionamos el modo
    l.switchEditMode()

    switch l.editSelection {
    case EDIT_MAP:
        // Los clicks dependen de la interfaz que haya activada.
        if ebiten.IsKeyPressed(ebiten.KeyDigit1) {
            if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
                l.tiles.ClickTileSelection(x, y)
            }
        } else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
            // Si no hay nada pulsado, pasarle el evento al modificador.
            l.tiles.ChangeTile(x, y)
        }
    case EDIT_ENTITIES:
        // Si se mueve la rueda, se cambia la entidad seleccionada.
        _, wy := ebiten.Wheel()
        if wy != 0 {
            l.enManager.ScrollEntity(int(wy))
        }

        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            selectedType := l.enManager.GetSelectedEntityType()
            l.enManager.SpawnByType(x, y, selectedType)
        }
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
            entityID := l.enManager.GetEntityIDAt(x, y)
            l.enManager.DeleteEntity(entityID)
        }
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
        // Tecla 1 para mostrar la selección de tiles
        l.tiles.DrawTileSelection(screen)
    } else {
        l.tiles.DrawLayer(screen, 0)
        l.enManager.Draw(screen)

        // Tras dibujar los tiles y las entidades, mostramos la interfaz.
        // Interfaz de edición.
        if l.editMode && l.editSelection == EDIT_ENTITIES {
            // Si estamos en modo edición de entidades, mostramos en una esquina
            // la entidad seleccionada.
            img := l.enManager.GetSelectedEntityImage()
            op := &ebiten.DrawImageOptions{}
            // Buscamos la coordanda inferior derecha.
            screenW, screenH := screen.Size()
            imgW, imgH := img.Size()
            imgX := screenW - imgW
            imgY := screenH - imgH
            op.GeoM.Translate(float64(imgX), float64(imgY))
            screen.DrawImage(img, op)
        }
    }

    if l.debug {
        ebitenutil.DebugPrint(screen, strings.Join(debugMessages, "\n"))
    }
}
