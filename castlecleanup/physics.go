
package main

import (
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Hitbox se encarga de comprobar colisiones con otras hitboxes
// ATENCIÓN: Parece que buena parte de esta funcionalidad ya está implementada
// en image.Rectangle
type Hitbox struct {
    X int
    Y int
    Width int
    Height int
}

// Colisión con otra hitbox
func (h *Hitbox) CheckCollision(h2 Hitbox) bool {
    return false
}

// Colisión con un elemento puntual
func (h *Hitbox) CheckPoint(x int, y int) bool {
    return false
}

// Esta clase se encarga básicamente de detectar las colisiones entre
// objetos y escenario.
type Physics struct {
    SolidityMap [][]int
    TileSize int
}

const AIR = 0
const SOLID = 1

func MakePhysics(mapSizeX int, mapSizeY int, tileSize int) *Physics {
    s_map := make([][]int, mapSizeY)
    for i := 0; i < mapSizeY; i++ {
        s_map[i] = make([]int, mapSizeX)
        for j := 0; j < mapSizeX; j++ {
            s_map[i][j] = AIR
        }
    }

    return &Physics{s_map, tileSize}
}

// Comprueba si una hitbox tiene colisión con algún elemento del mapa.
func (p *Physics) CollideMap(hitbox *Hitbox) bool {
    // Primero calculamos qué casillas del mapa de solidez está "pisando" la hitbox
    var stx, sty, edx, edy int
    stx = hitbox.X / p.TileSize
    sty = hitbox.Y / p.TileSize
    edx = (hitbox.X + hitbox.Width) / p.TileSize
    edy = (hitbox.Y + hitbox.Height) / p.TileSize

    // Recorremos el cuadrado de hitboxes para ver si alguna de ellas tiene colisión
    for y := sty; y < edy; y++ {
        for x := stx; x < edx; x++ {
            // TODO: Necesitamos una versión más avanzada que indique con qué elementos
            // se colisiona y no sólo que devuelva un booleano.
            if p.SolidityMap[y][x] == SOLID {
                return true
            }
        }
    }
    return false
}

// Esta función sólo es llamada si el modo debug está activado, por el que
// dibujará en pantalla los tiles que son sólidos.
func (p *Physics) DrawDebug(screen *ebiten.Image) {
    rectColor := color.RGBA{100, 100, 100, 127}
    for y := 0; y < len(p.SolidityMap); y++ {
        for x := 0; x < len(p.SolidityMap[y]); x++ {
            // TODO: Añadir soporte para varios tipos de "solidez"
            if p.SolidityMap[y][x] == SOLID {
                ebitenutil.DrawRect(screen, float64(x*p.TileSize), 
                        float64(y*p.TileSize), float64(p.TileSize),
                        float64(p.TileSize), rectColor)
            }
        }
    }
}
