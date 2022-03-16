package gameutils

import (
    "github.com/hajimehoshi/ebiten/v2"
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
func MakeSolidSpite(sprite Sprite, x0, y0, x1, y1 int) *SolidSprite {
    x := int(sprite.X)
    y := int(sprite.Y)
    hBox := image.Rect(x0 + x, y0 + y, x1 + x, y1 + y)

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

// Las estructuras que implementen esta interfaz podrán recibir eventos de
// control.
type ControlableSprite interface {
    // Encola acciones para ser realizadas todas a la vez cuando se llame a
    // la función Update
    EnqueueAction(action Action) bool

    // Devuelve una lista de las acciones posibles.
    GetActions() []string
}


// Construye un personaje a partir de un SolidSprite
func MakeCharacter(sprite SolidSprite) {
    //= []string{"moveX", "moveY"}
}

type Character struct {
    SolidSprite

    // Lo definimos como estático porque 
    possibleActions []string

    actionQueue []Action
}

func (c *Character) EnqueueAction(action Action) bool {
    c.actionQueue = append(c.actionQueue, action)
    return true // TODO: Devolver valor según el resultado.
}

func (c *Character) GetActions() []string {
    return c.possibleActions
}

