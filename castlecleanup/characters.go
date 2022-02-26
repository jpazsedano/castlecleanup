
package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "errors"
    "fmt"
    "image"
    _ "image/png"
    "bytes"
)

// Interfaz básica de Entidad
type BaseEntity interface {
    Update() error
    Draw(screen *ebiten.Image)
}

// EntityManager se encarga de lanzar las llamadas a Update y Draw de las
// entidades y de añadir y quitar elementos de la misma según se solicite,
// gestionando la carga de imágenes para que sólo se cargue una vez cada recurso.
type EntityManager struct {
    entities []BaseEntity
    resources map[string]ebiten.Image
}

func MakeEntityManager() (*EntityManager, error) {
    em := &EntityManager{}
    em.resources = make(map[string]ebiten.Image)
    // Como NewImageFromImage es un poco lento, cargamos todos los sprites aquí.
    for resName, resBytes := range SPRITE_RESOURCES {
        img, _, err := image.Decode(bytes.NewReader(resBytes))
        if err != nil {
            return nil, err
        }
        em.resources[resName] = ebiten.NewImageFromImage(img)
    }

    return em
}

// Funciones de gestión

// Esta función crea una entidad base con recurso de imagen
// compartido para ser utilizada en diferentes subclases
func (em *EntityManager) CreateEntity(x int, y int, resource string) (*Entity, error) {
    var res ebiten.Image
    var ok bool
    res, ok = em.resources[resource]
    if !ok {
        // Si no está cargado desde MakeEntityManager, entonces es un error
        return nil, errors.New(fmt.Sprintf("Resource %s not found", resource))
    }
    
    return &Entity{&res, x, y}
}

// Funciones de game loop

func (em *EntityManager) Update() error {
    for _, e := range em.entities {
        if err := e.Update(); err != nil {
            return err
        }
    }
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

// Una entidad sólida es una entidad que sigue sin moverse pero
// puede interactuar por contacto con otros elementos
type SolidEntity struct {
    Entity
    Hbox Hitbox
}


type Character struct {
    SolidEntity
}
