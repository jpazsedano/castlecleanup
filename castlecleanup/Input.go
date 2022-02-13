
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "fmt"
)

// Clases que implementen EventConsumer pueden recibir eventos de entrada.
type EventConsumer interface {
    // Recibe un evento y lo procesa. El valor de retorno indica si debe
    // seguir procesándose. Si devuelve un false, se para el proceso para el
    // resto de consumidores.
    ProcessEvent(e InputEvent) bool

    // ID asignado al escuchar. Puede ser -1 para indicar que aún no está registrado.
    GetConsumerId() int

    // Función que establece el valor del ID de consumidor. El valor recibido 
    // debe ser devuelto por GetConsumerId
    SetConsumerId(int)
}

// El evento de entrada tiene un tipo que lo identifica y un posible valor
// numérico. Básicamente indica que algo ha ocurrido. Opcionalmente puede
// tener uno o más valores numéricos (sticks, gatillos o clicks), pero eso
// ya está gestionado por cada implementación.
type InputEvent interface {
    GetType() int
}

type KeyEvent struct {
    eventId int
    PressDown bool // true si es KeyDown, false si es KeyUp
}

func (e KeyEvent*) GetType() string { e.eventId }

type ClickEvent struct {
    x int
    y int
}

func (e ClickEvent) GetType() string { Click }

// Los objetos que implementen esta interfaz deben poder capturar eventos
// de entrada desde la fuente que sea y devolver los InputEvent correspondientes.
type InputController interface {
    // Captura las entradas y envía los eventos. Debe ser llamado en Update.
    CaptureInput()
    // Esta función registra un consumer para recibir eventos de un determinado tipo.
    RegisterEventConsumer(consumer EventConsumer, event int) int
    // Esto elimina del registro de todos los eventos al elemento.
    UnregisterEventConsumer(consumerId int) bool
    // Esto otro elimina a un consumidor de un evento concreto.
    UnregisterFromEvent(consumerId int, event int) bool
}

// Claves de los eventos en 
const TilesetShowKey = 1
const Click = 2

// Este struct se encarga de guardar el mapeo de los controles, de manera
// sus métodos de captura de eventos pueden saber a qué evento de juego
// corresponde cada evento de entrada. También puede tener código para
// utilizar controladores si están disponibles.
type KeyboardInput struct {
    KeyMap map[int] ebiten.Key // Mapeo de evento a tecla que lo produce.
    Listeners map[int] []EventConsumer // Mapeo de evento a consumidor que lo quiere.
    // ¿Es necesario que esto sea concurrent-safe? Ahora mismo no lo es.
    LastId int
}

func MakeInputController() InputController {
    // TODO: Implementar la posibilidad de utilizar otros controladores
    controller := new(InputController)
    controller.KeyMap = make(map[int] ebiten.Key)
    controller.Listeners = make(map[int] []EventConsumer)
    controller.LastId = 0

    // Asignamos los eventos a las teclas.
    controller.KeyMap[TilesetShowKey] = ebiten.KeyF1

    return controller
}

func (c *KeyboardInput) CaptureInput() InputEvent[] {
    // Capturamos teclado
    for evID, v := range c.KeyMap {
        if inpututil.IsKeyJustPressed(v) {
            for _, l := range c.Listeners[evID] {
                l.ProcessEvent(KeyEvent{evID, true})
            }
        }
        if inpututil.IsKeyJustReleased(v) {
            for _, l := range c.Listeners[evID] {
                l.ProcessEvent(KeyEvent{evID, false})
            }
        }
    }
    // Capturamos clicks
    // TODO
}

func (c *KeyboardInput) RegisterEventConsumer(cons EventConsumer, event int) int 
    c.Listeners[event] = append(c.Listeners[event], cons)
    c.LastId++
    cons.SetConsumerId(c.LastId)
}

func (c *KeyboardInput) UnregisterEventConsumer(consumerId int) bool {
    // Para eliminar simplemente recorremos todos los eventos que tienen listener y eliminamos.
    for k, _ := range c.Listeners {
        c.UnregisterFromEvent(consumerId, k)
    }
}

// TODO: Test
func (c *KeyboardInput) UnregisterFromEvent(consumerId int, event int) bool {
    // Si la longitud de la lista de consumidores es 1, simplemente lo eliminamos.
    // O también si el elemento no existe.
    if _, ok = c.Listeners[event] ; !ok || len(c.Listeners[event]) <= 1 {
        c.Listeners[event] = []EventConsumer{}
    } else { // Si hay más de un item, hay que encontrar al customer que hay que eliminar.
        for i, v := range c.Listeners[event] {
            // Cuando encontramos al consumer a eliminar, lo hacemos.
            if v.GetConsumerId() == consumerId {
                c.Listeners[event][i] = c.Listeners[event][len(c.Listeners[event]) -1]
                c.Listeners[event] = c.Listeners[event][:len(c.Listeners[event]) -1]
                break
            }
        }
    }
}
