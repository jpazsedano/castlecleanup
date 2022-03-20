# Entidades y escenario

En un videojuego las entidades (personajes, drops, etc) interactúan con
el escenario, principalmente mediante la restrinción de movimientos. Para
conseguirlo, es necesario detectar la colisión de los hitboxes de los
`SolidSprite` con el mapa de solidez del escenario.

## Teoría

Por un lado tenemos los hitboxes de los objetos, cada uno tiene la suya, y
se define como un rectángulo del cual se pueden saber las coordenadas de cada
una de sus esquinas.

Por otro lado, tenemos la definición del mapa de solidez del escenario, el cual
está definido como una matriz de enteros, en la que cada entero indica la solidez
o no solidez de un bloque. A partir de esto, y sabiendo que cada bloque tiene una
anchura fija, podemos saber en qué coordenadas del mundo cae cada bloque. O también
qué casillas está "pisando" un rectángulo. Y si una de ellas es sólida, sabemos que
hay colisión.

Sin embargo, no sólo necesitamos saber si hay colisión, también necesitamos saber
hacia qué direcciones el movimiento está restringido y, si hay solapación, cómo
resolverla. Es muy importante que la restricción de movimiento se detecte antes
que la solapación, para tener colisiones firmes y sin "tembleques".

### Solapación de rectángulos y casillas

Para la detección básica de colisiones, podemos tomar las coordenadas de cada una
de las esquinas y, mediante una simple división entera por el ancho de casilla, 
podemos saber en qué casilla cae cada una. De esta manera tenemos índices iniciales
y finales para recorrer la matriz de solidez y comprobar colisiones.

Para realizar la comprobación entre el escenario y una entidad, podemos generar ad-hoc
una serie de rectángulos para las casillas con las que se tiene solapamiento (y las
próximas, para restricción de movimiento) y usar el mismo algoritmo que para la colisión
entre entidades.

### Solapación de rectángulos de entidades.

Utilizamos la estructura `Rect` del paquete `image` de la librería estándar, el cual
ya contiene una rutina de comprobación de colisiones entre rectángulos.

Sin embargo, nos interesa saber también en qué dirección restringe el movimiento y cómo
resolver una solapación. En el escenario era muy sencillo porque el escenario es inmóvil
y en caso de conflicto siempre es la entidad la que debe ser desplazada.

En el caso de las entidades, podemos definir entidades 'pesadas' y 'ligeras'. Si una de
cada tipo colisiona, siempre será desplazada la entidad ligera para resolver el conflicto.
Sin embargo, si ambas son ligeras o ambas son pesadas, se desplazará siempre la
entidad de la cual se está haciendo la comprobación, de esta manera una entidad pesada no
puede mover otra entidad pesada, y una entidad ligera no puede mover nada.

> Nota: Para ahorrar tiempo, cuando se comprueba la colisión entre entidades, se comprueba
únicamente la colisión de la entidad que se ha desplazado contra la lista de las demás.
Cuando aquí se habla de 'la entidad de la cual se está haciendo la comprobación', se habla
de la entidad que se ha desplazado, no ninguna de la lista.

Para decidir en qué dirección resolver el conficto, se calculan qué vértices están dentro
del rectángulo comprobado. Si hay 2, es muy sencillo, se debe 'empujar' en dirección
contraria al borde que definen esos dos.

Si sólo hay 1, se calcula escogen los bordes que están más cerca del centro del rectángulo
a corregir y se calcula la dirección en la cual empujar según la distancia de este único
vértice a dichos lados.

En el caso de que los 4 vértices estén dentro del rectángulo a comprobar, la rutina calcula
una dirección de preferencia mediante un struct recibido como parámetro que contendrá un
método de cálculo. Se utiliza un struct y no una función ya que, por ejemplo, para el
cálculo de la dirección al colisionar con el escenario, necesita todo el mapa de solidez
para calcular la dirección de preferencia. Esta dirección puede ser diagonal.

> **AVISO**: Comprobar que, efectivamente, el algoritmo de la librería estándar se comporta
como se espera en el caso de que dos rectángulos se solapen sólo en el borde.

Por definición un rectángulo que restringe el movimiento no está solapado con el rectángulo
que se está desplazando. Y si lo está, ya se habrá ejecutado la rutina de corrección antes
que la rutina de detección de restricciones, así que podemos dar por hecho que 

## Implementación de colisiones con escenario



## Implementación de colisiones entre entidades


