package inputmanager

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Funciones de comprobación de acción.

// El struct de acción incluye un identificador para saber qué se está
// haciendo así como un modificador para saber con qué intensidad se hace.
type Action struct {
    action string
    modifier float64
}

// La estructura InputManager se encarga de almacenar 
type InputManager struct {
    // Mapeo de acciones realizadas con 
    keyMap map[int]string
    controllerMap map[int]string

    controllerMode int
}

// 
func MakeInputManager() *InputManager {
    keyMap := make(map[int]string)
    controllerMap := make(map[int]string)

    // TODO: Asignar los mapeos por defecto

    return &InputManager{keyMap, controllerMap}
} 

// Esta función devuelve si una acción está ahora mismo realizándose. Si el botón
// o stick 
func (im *InputManager) IsActionPressed(action string) *Action {
    return nil
}

func (im *InputManager) IsActionJustPressed(action string) *Action {
    return nil
}

// Cambia los mapeos de acciones al teclado.
func (im *InputManager) MapKey(action string, key int) bool {
    return false
}

// Cambia los mapeos de acciones al controlador
func (im *InputManager) MapButton(action string, button int) bool {
    return false
}

// Devuelve la posición del puntero en la pantalla, independientemetne del tipo
// de puntero que sea.
func (im *InputManager) PointerPosition() (int, int) {
    // TODO: ¿La función de ebiten tiene ya la funcionalidad que debe tener esta?
    return ebiten.CursorPosition()
}
