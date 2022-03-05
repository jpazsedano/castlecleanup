# Árbol de entidades

Para evitar la duplicación de funcionalidad, las entidades del videojuego
están compuestas de diferentes super-entidades que implementan las partes
compartidas de su funcionalidad. En este artículo se detallan las entidades
que va a haber en el videojuego y qué funcionalidad van a compartir.

## Lista de entidades

**Jugador**: El jugador es la única entidad que es controlada por el jugador
y puede interactuar con el escenario. Sus funcionalidades son.

- Mostrar una imagen en pantalla, que puede cambiar según su estado (cayendo, corriendo,
saltando, atacando, etc). Esta es una funcionaliad compartida.
- Responder a las acciones de un controlador. Esta es una funcionalidad compartida,
ya que los enemigos también responden a acciones de un controlador IA.
- Responde a físicas. Esto es una funcionalidad compartida.
- Detecta colisiones con otras entidades. Esta es una función compratida, sin
embargo la manera de responder a estas colisiones es, en parte compartida, en parte única.
    - Específicamente, la respuesta a colisiones con el escenario es compartida.
    Al igual que otros elementos del juego se detiene al colisionar contra las paredes y 
    suelos.
    - La respuesta a colisiones con enemigos, monedas y corazones es única.

**Caja**: Las cajas son elementos que están en el suelo, que pueden ser lanzados o
destruidos, y que si se les destruye, hacen aparecer su contenido.

- Mostrar una imagen en pantalla, que puede cambiar según su estado. Func. compartida.
- Pueden ser lanzadas. Al cogerlas las cajas hacen despawn, dejando de existir en el juego.
Al lanzarlas, hace spawn una caja nueva e idéntica con un vector de movimiento determinado.
- Responde a físicas. Esto es una funcionalidad compartida.
- Al ser destruido puede hacer aparecer loot. Esto es una funcionalidad compartida, ya
que los cerdos también pueden hacer aparecer loot.
- Detecta colisiones con otras entidades. Esto es una funcionalidad compartida pero su
respuesta a la colisión es única.

**Bomba**: Las bombas son elementos que están en el suelo y pueden ser lanzados, pero no
destruidos, y que cuando son destruidos causan daños a las entidades cercanas.

- Mostrar una imagen en pantalla, que varía en función del estado. Funcionalidad compartida.
- Pueden ser lanzadas. Al cogerlas las bombas hacen despawn y al lanzarlas vuelven a existir
en modo "encendidas".
- Responde a físicas. Esto es una funcionalidad compartida.
- Tras ser activada y pasar algún tiempo, la bomba hace explosión, cambiando su estado y
creando una gran hitbox que hace daño a todas las entidades que toca.

**Cerdo**: El cerdo es el enemigo básico del juego. Se mueven por el escenario y si ven
al jugador, le atacan. Pueden configurarse para lanzar objetos lanzables.

- Mostrar una imagen en pantalla, que puede cambiar su estado. Funcionalidad compartida.
- Responder a las acciones de un controlador. En este caso el controlador es una IA que 
decide que acciones realizar en función del estado del juego.

**Cañón**: El cañón es una entidad que se mantiene en un sitio, pegada al suelo y puede
ser activada por un evento. 

**Disparo de cañón**: Muy similar a las cajas cuando son lanzadas, solo que en lugar de
loot, hacen aparecer explosiones.
