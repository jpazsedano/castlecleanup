package gameutils

// TODO: Para evitar colisiones de nombres esto quizás debería estar en su propio paquete.
import (
    _ "embed"
)

//go:embed assets/Terrain_32x32.png
var tileRawImage []byte

//go:embed assets/08-Box/Idle.png
var boxIdle []byte
//go:embed assets/09-Bomb/bomb-off.png
var bombOff []byte
//go:embed assets/01-King_Human/Fall.png
var kingFall []byte

const (
    CASTLE_TILEMAP = "castle"
)

var AM_RESOURCES = map[string][]byte{
    CASTLE_TILEMAP: tileRawImage,
}

const (
    BOX_IDLE = "box-idle"
    BOMB_OFF = "bomb-off"
    KING_FALL = "king-fall"
)

// Recursos para sprites.
var SPRITE_RESOURCES = map[string][]byte{
    BOX_IDLE: boxIdle,
    BOMB_OFF: bombOff,
    KING_FALL: kingFall,
}
