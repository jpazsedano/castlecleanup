# Animaciones

En el juego se distinguen dos tipos de animaciones: las animaciones de frames y las
animaciones generadas.

Las animaciones de frames son sencillamente cambios de frames de un objeto de manera
cíclica. Se carga una imagen que contiene varios frames con algunos metadatos y, cuando
se recibe una llamada a `Update` se cambia el frame cuando es necesario.

Las animaciones generadas reciben únicamente un valor a animar, un valor de inicio
y un valor de fin, así como una función de transición opcional.
Según el tiempo que deba durar la animación y los valores inicial y final se calculan
una serie de valores intermedios, a través de los cuales se itera cuando se llama a
la función `Update`.
La ventaja de las animaciones generadas es que permiten actuar sobre distintos tipos de
valores: Posición, rotación, escala, color, transparencia, tinte, recorte de imagen...

Las animaciones generadas pueden actuar únicamente sobre la imagen o sobre imagen y hitbox.
Para ello, ambos deben implementar una interfaz que pueda ser usada por el animador y
al registrar la animación, se puede registrar sólo la imagen, imagen y hitbox y técnicamente
también sólo el hitbox, aunque eso no tiene mucho sentido.

> Sobre cómo implementar las funciones de ease, mirar en https://easings.net.
