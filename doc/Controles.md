# Gestión de controles

En el transcurso del juego tendremos varios elementos que responderán
a la entrada del usuario:

- Entidades del juego. Normalmente será una única entidad, el personaje
del jugador. Sin embargo puede darse el caso de que varios personajes
obedezcan a las acciones del jugador.
- Acciones del nivel. Estamos hablando de acciones como poner el juego en
pausa, abrir el menú o realizar guardados o cargas rápidas. Cosas que no
afectan a ninguna entidad si no directamente a la escena en curso.

En ambos casos la entrada puede venir o bien del teclado o bien de un
controlador, por lo que es interesante tener un código que abstraiga
de la fuente de la acción y que capture únicamente acciones.

Como el código de la librería del juego es de tipo "pull", en el que son
los objetos los que tienen que comprobar si la entrada está activa, es
lógico que la capa de abstracción también sea tipo "pull" y no "push", por
lo que toma forma de un módulo al que en lugar de preguntarle si un botón
está pulsado o acaba de ser pulsado, se le pregunta si una acción ha
sido realizada o acaba de ser realizada.

## Control de la IA y de red

Como el control de la IA (y hipotéticamente también de red) transfiere las
acciones a las entidades (y posiblemente también al nivel) por la misma
vía utilizada para los controles del jugador, es lógico que estos implementen
la misma interfaz y que por lo tanto sean también "pull" y no "push".

Los eventos de acción de IA se producen durante el Update, que según el estado
del juego la IA decidirá qué acciones tomar. Los eventos de red, llegan por red
cuando menos te lo esperas. Son intrínsecamente "push". Por ello, la clase
que implemente la interfaz de controlador, debe mantener un buffer y proporcionar
las acciones cuando las entidades o el nivel las pida.
