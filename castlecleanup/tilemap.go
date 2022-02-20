
package main

import (
    "image"
    "image/color"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "encoding/gob"
    "io"
)

type SerializableTilemap struct {
    Layers [][][]int
    TileSize int
    TileXNum int
}

type Tilemap struct {
    tiles *ebiten.Image;
    layers [][][]int; // Tener en cuenta que las coordeadas son [y][x]
    tileSize int
    tileXNum int
    editorSelectedTile int
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
    return Tilemap{tiles, layers, tileSize, tileXNum, -1}
}

func (t *Tilemap) Serialize(buffer io.Writer) error {
    enc := gob.NewEncoder(buffer)
    return enc.Encode(t.layers)
}

func (t *Tilemap) Deserialize(buffer io.Reader) error {
    dec := gob.NewDecoder(buffer)
    return dec.Decode(&t.layers)
}

func (t *Tilemap) DrawTileSelection(screen *ebiten.Image) {
    // TODO: Eliminar duplicados..
    op := &ebiten.DrawImageOptions{}
    screen.DrawImage(t.tiles, op)

    // Si hay un tile seleccionado, se dibuja un rectángulo a su alrededor.
    if t.editorSelectedTile != -1 {
        xRect := float64((t.editorSelectedTile % t.tileXNum) * t.tileSize)
        yRect := float64((t.editorSelectedTile / t.tileXNum) * t.tileSize)

        rectColor := color.RGBA{255, 0, 0, 127}
        ebitenutil.DrawRect(screen, xRect, yRect, float64(t.tileSize), float64(t.tileSize), rectColor)
    }
}

// Esta función debe seleccionar un tile del selector de tiles. No lo muestra, de
// eso se encarga DrawTileSelection.
func (t *Tilemap) ClickTileSelection(x int, y int) {
    // TODO: Comprobar que el cursor esté dentro de la imagen y si no lo está, deseleccionar.
    // x e y son coordenadas de pantalla. Es esta clase la que sabe el tamaño
    // de tile y por lo tanto la que debe traducirla a coordenada de tile.
    tileX := x / t.tileSize
    tileY := y / t.tileSize
    t.editorSelectedTile = tileY * t.tileXNum + tileX
}

func (t *Tilemap) UnselectTile() {
    t.editorSelectedTile = -1
}

// Esta función debe superponer sobre la imagen el tile que se va a colocar
// a modo de "pista", añadiendo una transparencia parcial.
func (t *Tilemap) ShowSelectedTileInMap(screen *ebiten.Image, x int, y int) {

}

// 
func (t *Tilemap) ChangeTile(x int, y int) bool {
    if t.editorSelectedTile != -1 {
        tilemapX := x / t.tileSize
        tilemapY := y / t.tileSize
        if tilemapY < 0 || tilemapY >= len(t.layers[0]) || tilemapX < 0 || tilemapX >= len(t.layers[0][0]) {
            return false
        }

        // TODO: Tener en cuenta la capa
        t.layers[0][tilemapY][tilemapX] = t.editorSelectedTile
        return true
    } else {
        return false
    }
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
