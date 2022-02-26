package main

import (
    _ "embed"
)

const CASTLE_TILEMAP = "castle"

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte

//go:embed assets/08-Box/Idle.png
var boxIdle []byte

var AM_RESOURCES = map[string][]byte{
    "castle": tileRawImage,
}

// Recursos para sprites.
var SPRITE_RESOURCES = map[string][]byte{
    "box-idle": boxIdle,
}
