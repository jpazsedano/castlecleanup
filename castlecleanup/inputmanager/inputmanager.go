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

// Definimos el modo de controlador, que puede ser por mando o por teclado.
type CONTROLLER_MODE int
const (
    MODE_KEYBOARD CONTROLLER_MODE = iota
    MODE_CONTROLLER
)

type AXIS_ID int
const (
    MAIN_AXIS_H AXIS_ID = iota
    MAIN_AXIS_V
    SECONDARY_AXIS_H
    SECONDARY_AXIS_V
    LT_AXIS
    RT_AXIS
)

const (
    MOVE_LEFT string = "move-left"
    MOVE_RIGHT string = "move-right"
    JUMP string = "jump"
    ATTACK string = "attack"
)

const STICK_DEAD_ZONE = 0.3

func AvailableActions() []string {
    return []string{MOVE_LEFT, MOVE_RIGHT, JUMP, ATTACK}
}

// La estructura InputManager se encarga de almacenar 
type InputManager struct {
    // Mapeo de acciones realizadas con teclado.
    keyMap map[string]ebiten.Key
    // Mapeo de acciones realizadas con controlador.
    controllerMap map[string]ebiten.GamepadButton

    // Modo de controlador. Indica si la última acción se ha
    // realizado a través de teclado o de mando.
    controllerMode CONTROLLER_MODE
    // Definición de ejes.
    axisMap map[AXIS_ID]int

    // Flag que indica si el stick analógico debe considerarse
    // equivalente del pad direccional.
    analogEqualsDpad bool

    // Almacena las acciones realizadas 
    actionsDone map[string]bool
}

// Simplemente crea un inputmanager con parámetros por defecto.
func MakeDefaultInputManager() *InputManager {
    return MakeInputManager(true)
}

// Crea un InputManager con mapeos por defecto.
func MakeInputManager(analogEqualsDpad bool) *InputManager {
    keyMap := make(map[string]ebiten.Key)
    controllerMap := make(map[string]ebiten.GamepadButton)
    axisMap := make(map[AXIS_ID]int)
    actionsDone := make(map[string]bool)

    // TODO: Quizás los controles por defecto no deberían definirse aquí
    // Controles por defecto de teclado.
    keyMap[MOVE_LEFT] = ebiten.KeyArrowLeft
    keyMap[MOVE_RIGHT] = ebiten.KeyArrowRight
    keyMap[JUMP] = ebiten.KeyZ
    keyMap[ATTACK] = ebiten.KeyX
    // Controles por defecto de mando.
    controllerMap[MOVE_LEFT] = ebiten.GamepadButton11
    controllerMap[MOVE_RIGHT] = ebiten.GamepadButton12
    controllerMap[JUMP] = ebiten.GamepadButton0
    controllerMap[ATTACK] = ebiten.GamepadButton2
    axisMap[MAIN_AXIS_H] = 0
    axisMap[MAIN_AXIS_V] = 1
    axisMap[SECONDARY_AXIS_H] = 3
    axisMap[SECONDARY_AXIS_V] = 4
    axisMap[LT_AXIS] = 2
    axisMap[RT_AXIS] = 5

    actionsDone[MOVE_LEFT] = false
    actionsDone[MOVE_RIGHT] = false
    actionsDone[JUMP] = false
    actionsDone[ATTACK] = false

    return &InputManager{
        keyMap,
        controllerMap,
        MODE_KEYBOARD,
        axisMap,
        analogEqualsDpad,
        actionsDone,
    }
} 

// Esta función devuelve si una acción está ahora mismo realizándose. Devuelve
// nil si la acción on está realizada. Tiene en cuenta equivalencia de stick.
func (im *InputManager) IsActionPressed(action string) *Action {
    // Miramos la tecla respectiva y, si está pulsada, devolvemos el valor.
    key, ok := im.keyMap[action]
    // Si la acción no existe, por supuesto que no se está haciendo.
    if !ok {
        return nil
    }

    // Si el botón está pulsándose, se realiza.
    if ebiten.IsKeyPressed(key) {
        return &Action{action, 1.0} // Por el momento no hay modificador de la acción.
    } else {
        return nil
    }
}

// Esta función devuelve una acción si está ahora mismo realizándose.
// Devuelve nil si no. Tiene en cuenta equivalencia de stick.
func (im *InputManager) IsActionJustPressed(action string) *Action {
    key, ok := im.keyMap[action]

    if !ok {
        return nil
    }

    if inpututil.IsKeyJustPressed(key) {
        return &Action{action, 1.0}
    } else {
        return nil
    }
}

// Cambia los mapeos de acciones al teclado.
func (im *InputManager) MapKey(action string, key ebiten.Key) bool {
    switch action {
    // Sólo hacemos caso si es una de las acciones posibles.
    case MOVE_LEFT, MOVE_RIGHT, JUMP, ATTACK:
        im.keyMap[action] = key
        return true
    }
    return false
}

// Cambia los mapeos de acciones al controlador
func (im *InputManager) MapButton(action string, button ebiten.GamepadButton) bool {
    return false
}

func (im *InputManager) MapAxis(axis AXIS_ID, padAxis int) bool {
    return false
}

// Devuelve la posición del puntero en la pantalla, independientemetne del tipo
// de puntero que sea.
func (im *InputManager) PointerPosition() (int, int) {
    // TODO: ¿La función de ebiten tiene ya la funcionalidad que debe tener esta?
    // Es decir, el abstraer cursor y pulsación en pantalla. Y ¿Queremos esa funcionalidad?
    return ebiten.CursorPosition()
}
