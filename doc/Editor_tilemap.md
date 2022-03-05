# Editor de tilemap

Para la edición de tilemap nivel vamos a usar el propio motor del
juego encargado de mover dichos niveles. Seguimos esta filosofía para poder
probar los cambios realizados sin necesidad de mantener dos softwares diferentes
(editor y reproductor).

# Interfaz del editor

Para el editor vamos a usar una interfaz muy sencilla y controlada
casi enteramente con el ratón (el cual no se usará durante el juego)

Primero, para editar un nivel hará falta lanzar el juego en modo editor, de
momento es el único modo disponible, pero más adelante deberá parametrizarse.

Al lanzarse el juego en modo editor, el bucle principal escuchará la puslación
de las telas de función. Esto estará hardcodeado porque la edición únicamente se hará
en PC. Si en un futuro se portea el juego a otra plataforma (como móviles),
no se permitirá la edición de niveles en los mismos.

Debe ser el bucle principal y no el gestor de tiles el que escuche estas pulsaciones
porque no sólo estamos editando el tileset de fondo. En una fase inicial sí,
es lo único que se editará, pero en el futuro se editarán todos los elementos
del nivel.

Y además el selector de elementos debe superponerse no sólo al fondo de tiles
si no también a los otros elementos del nivel. Los cuales no debe ser ni siquiera
actualizados mientras se esté en la pantalla de selección de elementos.

Para la captura de la pulsación de la tecla F1 el nivel va a saltarse al gestor
de entrada que utilizará el resto del juego. Esto se hace así primero por
sencillez, y segundo porque la tarea del gestor de entrada es hacer keymaps
que sean configurables y los elementos del juego sólo reciban eventos genéricos
que representen acciones y sean independientes de cómo se han generado (sea la
tecla que sea o incluso que sea un botón de un gamepad)
