package gameutils

import (
    "github.com/hajimehoshi/ebiten/v2"
    "image"
)

// Interfaz básica de Entidad
type BaseEntity interface {
    Update() error
    Draw(screen *ebiten.Image)
    CheckPosition(x int, y int) bool
    GetPosition() (float64, float64)
    SetPosition(x float64, y float64)
    Move(x float64, y float64)
    GetImage() *ebiten.Image
}

// Implementación de Sprite para ser heredada por subclases y que compartan
// código básico de dibujado en pantalla, pero no se mueve ni interactúa.
type Sprite struct {
    // Referencia para compratir recurso entre entidades similares.
    image *ebiten.Image
    X float64
    Y float64
}

func (e *Sprite) Draw(screen *ebiten.Image) {
    // TODO: Ñapa temporal. En el futuro se debe utilizar una Animation
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(e.X, e.Y)
    screen.DrawImage(e.image, op)
}

// Como la entidad es estática, al actualizar no hace nada
func (e *Sprite) Update() error { return nil }

// Esta función comprueba si una posición está dentro de la imagen de la entidad.
// Si x e y es la posición del ratón, comprueba si el ratón está sobre la entidad.
func (e *Sprite) CheckPosition(x int, y int) bool {
    imgWidth, imgHeight := e.image.Size()
    fx := float64(x)
    fy := float64(y)
    return fx >= e.X && fx <= (e.X + float64(imgWidth)) && fy >= e.Y && (fy <= e.Y + float64(imgHeight))
}

func (e *Sprite) GetPosition() (float64, float64) {
    return e.X, e.Y
}

func (e *Sprite) SetPosition(x float64, y float64) {
    e.X = x
    e.Y = y
}
func (e *Sprite) Move(x float64, y float64) {
    e.X += x
    e.Y += y
}
func (e *Sprite) GetImage() *ebiten.Image {
    return e.image
}


// Construye un SolidSprite con su hitbox a partir de un Sprite y las
// coordenadas de su rectángulo.
func MakeSolidSprite(img *ebiten.Image, x, y float64, x0, y0, x1, y1 int) *SolidSprite {
    sprite := Sprite{img, x, y}
    hBox := image.Rect(x0 + int(x), y0 + int(y), x1 + int(x), y1 + int(y))

    return &SolidSprite{sprite, hBox}
}

// Una entidad sólida es una entidad que sigue sin moverse pero
// puede interactuar por contacto con otros elementos
type SolidSprite struct {
    Sprite
    Hbox image.Rectangle
}

// Establece la nueva posición del sprite
func (s *SolidSprite) SetPosition(x float64, y float64) {
    boxWidth := s.Hbox.Dx()
    boxHeight := s.Hbox.Dy()
    // El desplazamiento de la caja con respecto a la posición del sprite.
    boxXdesp := s.Hbox.Min.X - int(s.Sprite.X)
    boxYdesp := s.Hbox.Min.Y - int(s.Sprite.Y)
    // Movemos el sprite
    s.Sprite.SetPosition(x, y)
    // Y movemos la caja.
    s.Hbox.Min.X = int(x) + boxXdesp
    s.Hbox.Min.Y = int(y) + boxYdesp
    s.Hbox.Max.X = int(x) + boxXdesp + boxWidth
    s.Hbox.Max.Y = int(y) + boxYdesp + boxHeight
}

// Desplaza el sprite.
func (s *SolidSprite) Move(x float64, y float64) {
    currentX := s.Sprite.X
    currentY := s.Sprite.Y
    s.SetPosition(currentX + x, currentY + y)
}


// Construye un personaje con las funciones básicas.
func MakeCharacter(img *ebiten.Image, x, y float64, x0, y0, x1, y1 int, health int,
    moveSpeed, airFactor, jumpForce float64, canGrab, canShoot bool) *Character {
    sprite := MakeSolidSprite(img, x, y, x0, y0, x1, y1)

    
    return &Character{*sprite, health, health, moveSpeed, airFactor,
        jumpForce, canGrab, canShoot, []Action{}}
}

// Character representa cualquier personaje que pueda moverse, saltar y
// atacar.
type Character struct {
    SolidSprite

    health int
    maxHealth int
    moveSpeed float64 // Velocidad de movimiento
    airFactor float64 // Reducción de velocidad al estar en el aire
    jumpForce float64 // "fuerza" del salto. Define la altura
    canGrab bool // Define si puede agarrar cajas o bombas
    canShoot bool // Define si puede disparar cañones.

    actionQueue []Action // Acciones que el personaje debe hacer en el próximo frame (si puede)
}

func (c *Character) Update() error {
    // Llamamos a la función Update de la superclase ya que será necesario en
    // un futuro para las animaciones
    c.SolidSprite.Update()

    // Comprobamos acciones. De momento sólo movimiento horizontal, ya que aún no
    // hay físicas ni animaciones.
    for _, action := range c.actionQueue {
        // TODO: Completar con el resto de acciones (salto, ataque, agarrar y disparar)
        switch action.action {
        case MOVE_LEFT:
            c.Move(-c.moveSpeed, 0)
        case MOVE_RIGHT:
            c.Move(c.moveSpeed, 0)
        }
    }
    // Vaciamos la cola de acciones.
    c.actionQueue = []Action{}
    // Completar con la comprobación de colisión con otras entidades y si hacen daño.
    // La detección de muerte debe ser realizada por las estructuras contenedoras.
    return nil
}

func MakePlayer(img *ebiten.Image, x, y float64, x0, y0, x1, y1, health int,
        moveSpeed, airFactor, jumpForce float64, iManager *InputManager) *Player {
    character := MakeCharacter(img, x, y, x0, y0, x1, y1, health, moveSpeed, airFactor, jumpForce, false, false)
    // Queremos pasar el valor, no el puntero. ¿Se hace así?
    return &Player{*character, iManager}
}

type Player struct {
    Character

    iManager *InputManager
}

func (p *Player) Update() error {
    // TODO: Añadir salto y ataque.
    // Capturamos eventos y actualizamos el personaje.
    if a := p.iManager.IsActionPressed(MOVE_LEFT); a != nil {
        p.Character.actionQueue = append(p.Character.actionQueue, *a)
    }
    if a := p.iManager.IsActionPressed(MOVE_RIGHT); a != nil {
        p.Character.actionQueue = append(p.Character.actionQueue, *a)
    }

    return p.Character.Update()
}
