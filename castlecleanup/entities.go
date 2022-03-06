
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "errors"
    "fmt"
    "image"
    _ "image/png"
    "bytes"

    "github.com/jpazsedano/castlecleanup/inputmanager"
)

const (
    ENTITY_BOX int = iota
    ENTITY_BOMB
)

const NUM_ENTITIES = 2

var ENTITY_ASSIGNATIONS = map[int]string{
    ENTITY_BOX: BOX_IDLE,
    ENTITY_BOMB: BOMB_OFF,
}

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

// EntityManager se encarga de lanzar las llamadas a Update y Draw de las
// entidades y de añadir y quitar elementos de la misma según se solicite,
// gestionando la carga de imágenes para que sólo se cargue una vez cada recurso.
type EntityManager struct {
    entities map[int]BaseEntity
    resources map[string]*ebiten.Image

    // Representa la lista de imágenes para la edición de las entidades.
    availableEntities map[int]*ebiten.Image
    entitySelected int

    lastid int
}

// Esta función está pensada para crear una imagen de edición
// diferente del recurso original, sin embargo ahora mismo sólo
// devuelve el mismo valor.
func createEditImage(src *ebiten.Image) *ebiten.Image {
    return src
}

func MakeEntityManager() (*EntityManager, error) {
    em := &EntityManager{}
    em.resources = make(map[string]*ebiten.Image)
    em.entities = make(map[int]BaseEntity)
    em.lastid = 0
    // Como NewImageFromImage es un poco lento, cargamos todos los sprites aquí.
    for resName, resBytes := range SPRITE_RESOURCES {
        img, _, err := image.Decode(bytes.NewReader(resBytes))
        if err != nil {
            return nil, err
        }
        em.resources[resName] = ebiten.NewImageFromImage(img)
    }

    em.availableEntities = make(map[int]*ebiten.Image)
    // Rellenamos availableEntities con los sprites recién cargados
    for id, resource := range ENTITY_ASSIGNATIONS {
        em.availableEntities[id] = createEditImage(em.resources[resource])
    }

    return em, nil
}

// Funciones de gestión

// Esta función crea una entidad base con recurso de imagen
// compartido para ser utilizada en diferentes subclases
func (em *EntityManager) CreateEntity(x int, y int, resource string) (BaseEntity, error) {
    var res *ebiten.Image
    var ok bool
    res, ok = em.resources[resource]
    if !ok {
        // Si no está cargado desde MakeEntityManager, entonces es un error
        return nil, errors.New(fmt.Sprintf("Resource %s not found", resource))
    }
    
    return &Sprite{res, float64(x), float64(y)}, nil
}

// Registra una entidad en el manager. Normalmente es una subclase que ha utilizado
// la clase generada por CreateEntity como composite.
func (em *EntityManager) SpawnEntity(entity BaseEntity) int {
    em.lastid++
    em.entities[em.lastid] = entity

    return em.lastid
}

// Elimina una entidad del gestor de entidades a partir de su ID
func (em *EntityManager) DeleteEntity(entityID int) bool {
    _, ok := em.entities[entityID]
    // Si la entidad existe, se elimina.
    if ok {
        delete(em.entities, entityID)
    }
    return ok
}

// Funciones de edición

// Esta función devuelve la lista de entidades disponibles.
func (em *EntityManager) getEntityList() []int {
    keys := make([]int, len(em.availableEntities))
    i := 0
    for key, _ := range em.availableEntities {
        keys[i] = key
    }

    return keys
}

// Devuelve el identificador de la entidad ubicada en las coordenadas dadas.
// Las coordenadas deben ser coordenadas de mundo, no de pantalla.
// Devuelve -1 si no se ha encontrado ninguna entidad en dichas coordenadas.
func (em *EntityManager) GetEntityIDAt(x int, y int) int {
    for id, entity := range em.entities {
        if entity.CheckPosition(x, y) {
            return id
        }
    }
    // No se ha encontrado ninguna entidad.
    return -1
}

// Esta función de edición selecciona cambia la selección de la entidad a hacer aparecer.
func (em *EntityManager) ScrollEntity(dir int) {
    // Normalizamos la dirección
    var dScroll int = 0
    if dir < 0 {
        dScroll = -1
    } else if dir > 0 {
        dScroll = 1
    }

    em.entitySelected += dScroll
    // Si es menor que 0, volvemos al final
    if em.entitySelected < 0 {
        em.entitySelected = len(em.availableEntities) - 1
    }
    // Nos aseguramos 
    em.entitySelected = em.entitySelected % len(em.availableEntities)
}

// Hace spawn de un elemento según su tipo
func (em *EntityManager) SpawnByType(x int, y int, e_type int) error {
    // Hardcodeamos la manera de inicializar cada entidad.
    switch e_type {
    case ENTITY_BOX:
        // La caja será un SolidSprite cuando esté implementado, por el momento sólo es una entidad normal
        entity, err := em.CreateEntity(x, y, BOX_IDLE)
        if err != nil {
            return err
        }
        imgW, imgH := entity.GetImage().Size()
        entity.Move(float64(-imgW/2), float64(-imgH/2))
        em.SpawnEntity(entity)
    case ENTITY_BOMB:
        entity, err := em.CreateEntity(x, y, BOMB_OFF)
        if err != nil {
            return err
        }
        imgW, imgH := entity.GetImage().Size()
        entity.Move(float64(-imgW/2), float64(-imgH/2))
        em.SpawnEntity(entity)
    }
    return nil
}

// Devuelve el tipo de entidad seleccionada (que se puede cambiar con ScrollEntity)
func (em *EntityManager) GetSelectedEntityType() int {
    return em.entitySelected
}

// Devuelve una imagen representativa de la entidad seleccionada.
func (em *EntityManager) GetSelectedEntityImage() *ebiten.Image {
    selected := em.GetSelectedEntityType()
    return em.availableEntities[selected]
}

// Funciones de game loop

func (em *EntityManager) Update() error {
    for _, e := range em.entities {
        if err := e.Update(); err != nil {
            return err
        }
    }
    return nil
}
// FIXME: Error en el interior.
func (em *EntityManager) Draw(screen *ebiten.Image) {
    // FIXME: Esto no funciona, range no respeta ningún orden concreto y
    // si dos entidades se solapan, parpadean. Hay que usar algo que respete orden.
    for _, e := range em.entities {
        e.Draw(screen)
    }
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
    // Implementación básica, simplemente dibujamos la imagen en pantalla.
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
    EnqueueAction(action inputmanager.Action) bool

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

    actionQueue []inputmanager.Action
}

func (c *Character) EnqueueAction(action inputmanager.Action) bool {
    c.actionQueue = append(c.actionQueue, action)
    return true // TODO: Devolver valor según el resultado.
}

func (c *Character) GetActions() []string {
    return c.possibleActions
}
