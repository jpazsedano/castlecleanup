
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
    layers [][][]int; // Tener en cuenta que las coordeadas son [y][x]
    tileSize int
    tileXNum int
}

func MakeEmptyTilemap(editMode bool, tiles *ebiten.Image, nLayers int, sizeX int, sizeY int, tileSize int, tileXNum int) {
    // Inicializamos unos tiles vacíos.
    layers := make([][][]int, sizeY)
    for i := 0; i < nLayers; i++ {
        layers[i] = make([][]int, sizeY)
        for j := j < sizeY; j++ {
            layers[i][j] = make([]int, sizeX)
        }
    }
    return Tilemap{editMode, tiles, layers, tileSize, tileXNum}
}

// Carga un tilemap de fichero, que no es mas que el serializado de las capas.
func (t *Tilemap) LoadTilemap(tilemapImage image.Image, tilemapData io.Reader) {

}

func (t *Tilemap) SaveTilemap(writer io.Writer) {

}

// Esta función dibuja una capa del tilemap en la pantalla (o imagen) recibida.
func (t *Tilemap) DrawLayer(screen *ebiten.Image, layer int) {
    // Pillamos la capa, y por cada tile que haya que dibujar, lo obtenemos,
    // lo transformamos y lo copiamos a la pantalla.
    for ir, row := range t.layers[layer] {
        for it, tile := range t.layers[layer][ir] {
            // Seleccionamos el tile. Calculamos las coordenadas según la posicón.
            xTile := (tile % t.layers.tileXNum) * t.tileSize
            yTile := (tile / t.layers.tileXNum) * t.tileSize
            tileImg := t.tiles.SubImage(image.Rect(xTile, yTile, xTile+t.tileSize, yTile+t.tileSize))
            // Dibujamos el tile en la pantalla
            op := &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(ir*t.tileSize), float64(it*t.tileSize))
            // Dibujamos en pantalla la imagen casteada con la traslación.
            screen.DrawImage(tileImg.(*ebiten.Image), op)
        }
    }
}

// El tilemap puede recibir eventos de entrada. No los captura él mismo
// (como ningún objeto), si no que los recibe del objeto InputController.
func (t *Tilemap) ProcessEvent(e InputEvent) bool {
    if t.editMode {

    }  else { // Si no hay modo edición no se hace nada.
        return true
    }
}
