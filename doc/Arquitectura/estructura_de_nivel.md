# Estructura de nivel

En este documento se discuten las funcionalidades del nivel y cómo se va a
estructurar, así como los distintos componentes que deben formarlo y
la funcionalidad que éstos deben ofrecer para que el nivel funcione como debe.

También se discuten las fucionalidades que ofrecerá el nivel para que estos
componentes puedan realizar su función, así como las funcionalidades que el nivel
ofrecerá al nivel superior de componentes.

## Dónde encaja el nivel

El nivel es una implementación de Escena, el cual será utilizado por Game
y se encargará de toda la lógica y representación del entorno de juego principal.

El nivel y el bucle del juego se compone básicamente de un entorno, una serie
de entidades que habitan dicho entorno, unos eventos de entrada que realizan
modificaciones y una representación gráfica que tiene que llgar hasta la pantalla.

- El tilemap. El cual se encarga de gestionar la parte estática del escenario.
Esto incluye la definición de los distintos bloques con sus propiedades:
gráfico que lo representa y solidez<sup>A1</sup>. Como éste código es el que conoce
su estructura interna, también ofecerá métodos para realizar comprobaciones de
colisión con el mapa de solidez.
- Gestor de entrada (`inputmanager`). Definiendo unas acciones y unos
mapeos por defecto, este componente se encarga de ofrecer funciones para
poder capturar estas acciones sin preocuparse de dónde vienen, del mapeo
específico de teclas o incluso de si se usa un teclado o un mando.
- Gestor de entidades. Las entidades deben ser actualizadas y dibujadas en
cada frame. El gestor de entidades se encarga de hacer llegar el ciclo de
actualización y de dibujado a todas las entidades, y de hacer llegar información
sobre el resto del nivel a las mismas, para que puedan tomar decisiones sobre
su actualización. También organiza las entidades para que el nivel pueda sacar
un sentido de las mismas que afecte al estado del nivel en general. Ejemplo simple:
poder solicitar la entidad del jugador para saber cuánta salud le queda y si
está vivo o muerto, ya que si está muerto el juego se pierde y el nivel debe
pasar a su estado de fracaso.

```
    Game
    +-Nivel (escena)
      +-Tilemap
      +-Inputmanager
      +-EntityManager
```

## Funcionalidades que ofrece Nivel al componente superior

El componente superior es Game, el cual se encarga de gestionar los cambios
de escena y su ciclo de vida. Por lo tanto, las implementaciones de escena,
como Nivel, deben implementar métodos para ofrecer ese control.

- `Initialize`: Este método será llamado una única vez cuando la escena esté
cargando, ya sea al inicio del juego o al cambiar desde otra escena.
- `Update`: Este método será llamado en cada frame para actualizar el estado
interno de la escena en tiempo real.
- `Draw`: Este método será llamado en cada frame para obtener una representación
gráfica del estado de la escena.

No existe método `Finalize` ya que el propio lenguaje Go cuenta con recolector de basura.

El Nivel también debe llamar a `SwitchScene` cuando, por las circunstancias que
sean, toque cambiar de escena o cerrar el juego.

### Funcionalidades del componente superior requeridas

El único requisito que se le pide al nivel superior es llamar a las funciones
del ciclo de vida en los momentos correctos y proporcionar un método estandarizado
para el cambio de escena.

También se le requiere proporcionar en el dibujado una superficie en la que dibujar
y en la actualización un temporizador para saber cuánto tiempo ha transcurrido desde
la última llamada (delta).

Por último, al inicializar requiere del componente superior que se le indique si
el juego está en modo debug o no, para saber si el modo edición debe estar activo o
inactivo.

## Funcionalidades que ofrece Nivel para componentes inferiores

Algunos componentes inferiores necesitan que se les proporcionen algunas entradas para
poder funcionar.

- **Tilemap**: El tilemap no requiere mucha atennción durante el juego, ya que por 
naturaleza es estático, sin embargo necesitará que se le pasen los datos a cargar al
inicializar y, durante la edición que se le indique la selección de tile y la modificación
del entorno estático.
- **Inputmanager**: Principalmente durante la inicialización, ya que los controles se
escogen en el menú, antes de iniciar 

## Funcionalidades de Nivel para uso interno

En este caso hablamos principalmente de:

- Mantenimiento del estado de juego. Esto son las puntuaciones, número de enemigos
y salud del jugador, entre otros parámetros. Esto se hace con ayuda principalmente del
gestor de entidades.
- Cambio de escena. Cuando el nivel alcanza un cierto estado (de éxito o de fracaso), o
cuando el jugador solicita volver al menú o cerrar el juego, esto debe ser capturado
por el nivel.
- Modo edición. Otra de las tareas del Nivel, además de gestionar el funcionamiento
del mismo, es permitir editarlo. De esta manera podemos generar niveles y probarlos
en el momento, con una misma herramienta que utiliza siempre una misma implementación.
Por ello, el nivel debe gestionar el cambio a modo edición y la interfaz del mismo para
cambiar los elementos estáticos y dinámicos.

# Estado de implementación

Esto es básicamente un trabajo en progreso. En esta sección se describe la realidad
de las funcionalidades que están implementadas y las que aún no de la entidad Nivel.

# Apéndices

Información complementaria que puede ser de utilidad.

- A1: Estas son las únicas características que se implementarán en esta versión,
ya que son las únicas que utilizará este juego. En futuras versiones será
interesante poder añadir cosas como agua, lava y otras que luego el juego decidirá
cómo interpretar.
