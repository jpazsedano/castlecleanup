# Entidades y controladores

Una entidad en el motor es cualquier elemento que aparezca de manera independiente
en la pantalla y tenga posición, rutina de dibujado y lógica de actualización.

Las entidades son manejadas por un gestor de entidades, el cual es actualizado por
la escena. El gestor de entidades a su vez llama al método `Update` y `Draw` de las
entidades. Cada entidad implementa su lógica y su rutina de dibujado.

## Controladores

Los controladores pueden ser externos o internos.

- Controladores externos: Son una capa que contiene a la entidad a controlar. El
gestor de entidades llama a su método `Update` y el controlador externo es el que
llama al método `Update` de la entidad controlada. O no, si quiere tomar el control
y evitar su lógica interna. Útil para cinemáticas. Se podría decir que es 'push'
- Controlador interno: Un controlador interno es llamado por la propia entidad 
controlada, para obtener las acciones que debe realizar. Los detalles sobre cómo
realizarlas son gestionadas por la implementación de la entidad.

## Rutinas de dibujado y animaciones

La rutina de dibujo es implementada por una estructura `Animation` la cual
tiene un método de actualización (para hacer correr la animación) y otro de obtención
del frame actual. Este último método puede aplicar una subdivisión para una animación
por cambio de frame, o puede aplicar una transformación (rotación, escala, traslación, etc.)

Estas transformaciones y la división de frame son métodos diferentes, ya que en el motor
ebiten se hacen en dos pasos diferentes.

Todas las animaciones tienen un principio y un fin, pudiendo ser cíclicas. También
tienen un método para indicar si han llegado a ese fin en este frame.
