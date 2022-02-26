package main

import (
	_ "embed"
)

const CASTLE_TILEMAP = "castle"

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte

var Resources = map[string][]byte{
	"castle": tileRawImage,
}
