
package main

import (
    "image"
    "io"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tilemap struct {
    editMode bool;
    tiles *ebiten.Image;
    layers [][] int;
}

// Carga un tilemap de fichero, que no es mas que el serializado de las capas.
func (t *Tilemap) LoadTilemap(tilemapImage image.Image, tilemapData io.Reader) {

}

func (t *Tilemap) SaveTilemap(writer io.Writer) {

}

// Esta función dibuja una capa del tilemap en la pantalla (o imagen) recibida.
func (t *Tilemap) DrawLayer(screen *ebiten.Image, layer int) {

}

// El tilemap puede recibir eventos de entrada. No los captura él mismo
// (como ningún objeto), si no que los recibe del objeto InputController.
func (t *Tilemap) ProcessEvent(e InputEvent) bool {
    if t.editMode {

    }  else { // Si no hay modo edición no se hace nada.
        return true
    }
}
