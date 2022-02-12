
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Clases que implementen EventConsumer pueden recibir eventos de entrada.
type EventConsumer interface {
    // Recibe un evento y lo procesa. El valor de retorno indica si debe
    // seguir procesándose. Si devuelve un false, se para el proceso para el
    // resto de consumidores.
    ProcessEvent(e InputEvent) bool
}

// El evento de entrada tiene un tipo que lo identifica y un posible valor
// numérico. Básicamente indica que algo ha ocurrido. Opcionalmente puede
// tener uno o más valores numéricos (sticks, gatillos o clicks), pero eso
// ya está gestionado por cada implementación.
type InputEvent interface {
    GetType() string
}

// Los objetos que implementen esta interfaz deben poder capturar eventos
// de entrada desde la fuente que sea y devolver los InputEvent correspondientes.
type InputController interface {
    // Captura las entradas y envía los eventos. Debe ser llamado en Update.
    CaptureInput()
    // Esta función registra un consumer para recibir eventos de un determinado tipo.
    RegisterEventConsumer(consumer EventConsumer, event int) int
    // Esto elimina del registro de todos los eventos al elemento.
    UnregisterEventCustomer(customerId int) bool
    // Esto otro elimina a un consumidor de un evento concreto.
    UnregisterFromEvent(customerId int, event int) bool
}

// Claves de los eventos en 
const TilesetShowKey = 1
const Click = 2

// Este struct se encarga de guardar el mapeo de los controles, de manera
// sus métodos de captura de eventos pueden saber a qué evento de juego
// corresponde cada evento de entrada. También puede tener código para
// utilizar controladores si están disponibles.
type KeyboardInputController struct {
    DebugTilesetShowKey ebiten.Key
    listeners int[]
}

func MakeInputController() InputController {
    // TODO: Implementar la posibilidad de utilizar otros controladores
    controller := new(InputController)
    cotroller.DebugTilesetShowKey = ebiten.KeyF1

    return controller
}

func (c KeyboardInputController*) CaptureInput() InputEvent[] {
    // 
    if inpututil.IsKeyJustPressed(c.DebugTilesetShowKey) {

    }
    // 
    if inpututil.IsKeyJustReleased(c.DebugTilesetShowKey) {

    }
}
