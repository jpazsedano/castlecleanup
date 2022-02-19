
package main

import (
    "image"
    "io"
    "github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
    tiles *ebiten.Image;
    layers [][][]int; // Tener en cuenta que las coordeadas son [y][x]
    tileSize int
    tileXNum int
}

// OJO: TileXNum no es el número de tiles en X en pantalla, si no en el recurso.
func MakeEmptyTilemap(tiles *ebiten.Image, nLayers int, sizeX int, sizeY int, tileSize int, tileXNum int) Tilemap {
    // Inicializamos unos tiles vacíos.
    layers := make([][][]int, sizeY)
    for i := 0; i < nLayers; i++ {
        layers[i] = make([][]int, sizeY)
        for j := 0; j < sizeY; j++ {
            layers[i][j] = make([]int, sizeX)
        }
    }
    return Tilemap{tiles, layers, tileSize, tileXNum}
}

// Carga un tilemap de fichero, que no es mas que el serializado de las capas.
func (t *Tilemap) LoadTilemap(tilemapImage image.Image, tilemapData io.Reader) {

}

func (t *Tilemap) SaveTilemap(writer io.Writer) {

}

// Esta función dibuja una capa del tilemap en la pantalla (o imagen) recibida.
// Dejamos en manos del código de escena escoger las capas para poder poner unas debajo
// de los sprites y otras encima. O poder aplicar efectos a capas específicas.
func (t *Tilemap) DrawLayer(screen *ebiten.Image, layer int) {
    // Pillamos la capa, y por cada tile que haya que dibujar, lo obtenemos,
    // lo transformamos y lo copiamos a la pantalla.
    for ir, row := range t.layers[layer] {
        for it, tile := range row {
            // Seleccionamos el tile. Calculamos las coordenadas según la posicón.
            xTile := (tile % t.tileXNum) * t.tileSize
            yTile := (tile / t.tileXNum) * t.tileSize
            tileImg := t.tiles.SubImage(image.Rect(xTile, yTile, xTile+t.tileSize, yTile+t.tileSize))
            // Dibujamos el tile en la pantalla
            op := &ebiten.DrawImageOptions{}
            op.GeoM.Translate(float64(it*t.tileSize), float64(ir*t.tileSize))
            // Dibujamos en pantalla la imagen casteada con la traslación.
            screen.DrawImage(tileImg.(*ebiten.Image), op)
        }
    }
}
