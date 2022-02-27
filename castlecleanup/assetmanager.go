package main

// TODO: Para evitar colisiones de nombres esto quizás debería estar en su propio paquete.
import (
    _ "embed"
)

const (
    CASTLE_TILEMAP = "castle"
)

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte

//go:embed assets/08-Box/Idle.png
var boxIdle []byte

var AM_RESOURCES = map[string][]byte{
    CASTLE_TILEMAP: tileRawImage,
}

const (
    BOX_IDLE = "box-idle"
)

// Recursos para sprites.
var SPRITE_RESOURCES = map[string][]byte{
    BOX_IDLE: boxIdle,
}
