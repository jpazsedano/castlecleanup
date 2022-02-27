
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "errors"
    "fmt"
    "image"
    _ "image/png"
    "bytes"
)

type EntityType int64

const (
    ENTITY_BOX EntityType = iota
)

var ENTITY_ASSIGNATIONS = map[EntityType]string{
    ENTITY_BOX: BOX_IDLE,
}

// Interfaz básica de Entidad
type BaseEntity interface {
    Update() error
    Draw(screen *ebiten.Image)
    CheckPosition(x int, y int) bool
}

// EntityManager se encarga de lanzar las llamadas a Update y Draw de las
// entidades y de añadir y quitar elementos de la misma según se solicite,
// gestionando la carga de imágenes para que sólo se cargue una vez cada recurso.
type EntityManager struct {
    entities map[int]BaseEntity
    resources map[string]*ebiten.Image

    // Representa la lista de imágenes para la edición de las entidades.
    availableEntities map[EntityType]*ebiten.Image

    lastid int
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

    em.availableEntities = make(map[EntityType]*ebiten.Image)
    // Rellenamos availableEntities con los sprites recién cargados
    for id, resource := range ENTITY_ASSIGNATIONS {
        em.availableEntities[id] = em.resources[resource]
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
    
    return &Entity{res, x, y}, nil
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
func (em *EntityManager) getEntityList() []EntityType {
    keys := make([]EntityType, len(em.availableEntities))
    i := 0
    for key, _ := range em.availableEntities {
        keys[i] = key
    }

    return keys
}

// Devuelve el identificador de la entidad ubicada en las coordenadas dadas.
// Las coordenadas deben ser coordenadas de mundo, no de pantalla.
// Devuelve -1 si no se ha encontrado ninguna entidad en las coordenadas dadas.
func (em *EntityManager) getEntityIDAt(x int, y int) int {
    for id, entity := range em.entities {
        if entity.CheckPosition(x, y) {
            return id
        }
    }
    // No se ha encontrado ninguna entidad.
    return -1
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
func (em *EntityManager) Draw(screen *ebiten.Image) {
    for _, e := range em.entities {
        e.Draw(screen)
    }
}

// Implementación de Entity para ser heredada por subclases y que compartan
// código básico de dibujado en pantalla, pero no se mueve ni interactúa.
type Entity struct {
    // Referencia para compratir recurso entre entidades similares.
    image *ebiten.Image
    X int
    Y int
}

func (e *Entity) Draw(screen *ebiten.Image) {

}

// Como la entidad es estática, al actualizar no hace nada
func (e *Entity) Update() error { return nil }

// Esta función comprueba si una posición está dentro de la imagen de la entidad.
// Si x e y es la posición del ratón, comprueba si el ratón está sobre la entidad.
func (e *Entity) CheckPosition(x int, y int) bool {
    imgWidth, imgHeight := e.image.Size()
    return x >= e.X && x <= (e.X + imgWidth) && y >= e.Y && (y <= e.Y + imgHeight)
}

// Una entidad sólida es una entidad que sigue sin moverse pero
// puede interactuar por contacto con otros elementos
type SolidEntity struct {
    Entity
    Hbox Hitbox
}


type Character struct {
    SolidEntity
}
