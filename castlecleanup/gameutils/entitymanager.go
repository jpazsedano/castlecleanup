
package gameutils

import (
    "github.com/hajimehoshi/ebiten/v2"
    "errors"
    "fmt"
    "image"
    _ "image/png"
    "bytes"
)

const (
    ENTITY_BOX int = iota
    ENTITY_BOMB
    ENTITY_PLAYER
)

const NUM_ENTITIES = 3

var ENTITY_ASSIGNATIONS = map[int]string{
    ENTITY_BOX: BOX_IDLE,
    ENTITY_BOMB: BOMB_OFF,
    ENTITY_PLAYER: KING_FALL, // Por ahora ponemos la imagen del rey cayendo
}

// EntityManager se encarga de lanzar las llamadas a Update y Draw de las
// entidades y de añadir y quitar elementos de la misma según se solicite,
// gestionando la carga de imágenes para que sólo se cargue una vez cada recurso.
type EntityManager struct {
    entities map[int]BaseEntity
    resources map[string]*ebiten.Image

    // Representa la lista de imágenes para la edición de las entidades.
    availableEntities map[int]*ebiten.Image
    entitySelected int

    lastid int

    iManager *InputManager // Para las entidades que escuchan al input
}

// Esta función está pensada para crear una imagen de edición
// diferente del recurso original, sin embargo ahora mismo sólo
// devuelve el mismo valor.
func createEditImage(src *ebiten.Image) *ebiten.Image {
    return src
}

func MakeEntityManager() (*EntityManager, error) {
    em := &EntityManager{}
    em.resources = make(map[string]*ebiten.Image)
    em.entities = make(map[int]BaseEntity)
    em.lastid = 0
    // Como NewImageFromImage es un poco lento, cargamos todos los sprites aquí.
    for resName, resBytes := range SPRITE_RESOURCES {
        img, _, err := image.Decode(bytes.NewReader(resBytes))
        if err != nil {
            return nil, err
        }
        em.resources[resName] = ebiten.NewImageFromImage(img)
    }

    em.availableEntities = make(map[int]*ebiten.Image)
    // Rellenamos availableEntities con los sprites recién cargados
    for id, resource := range ENTITY_ASSIGNATIONS {
        em.availableEntities[id] = createEditImage(em.resources[resource])
    }
    em.iManager = MakeDefaultInputManager()

    return em, nil
}

// Funciones de gestión

// Esta función crea una entidad base con recurso de imagen
// compartido para ser utilizada en diferentes subclases
func (em *EntityManager) CreateSprite(x int, y int, resource string) (BaseEntity, error) {
    var res *ebiten.Image
    var ok bool
    res, ok = em.resources[resource]
    if !ok {
        // Si no está cargado desde MakeEntityManager, entonces es un error
        return nil, errors.New(fmt.Sprintf("Resource %s not found", resource))
    }
    
    return &Sprite{res, float64(x), float64(y)}, nil
}

// Registra una entidad en el manager. Normalmente es una subclase que ha utilizado
// la clase generada por CreateSprite como composite.
func (em *EntityManager) SpawnEntity(entity BaseEntity) int {
    em.lastid++
    em.entities[em.lastid] = entity

    return em.lastid
}

// Elimina una entidad del gestor de entidades a partir de su ID
func (em *EntityManager) DeleteEntity(entityID int) bool {
    _, ok := em.entities[entityID]
    // Si la entidad existe, se elimina.
    if ok {
        delete(em.entities, entityID)
    }
    return ok
}

// Funciones de edición

// Esta función devuelve la lista de entidades disponibles.
func (em *EntityManager) getEntityList() []int {
    keys := make([]int, len(em.availableEntities))
    i := 0
    for key, _ := range em.availableEntities {
        keys[i] = key
    }

    return keys
}

// Devuelve el identificador de la entidad ubicada en las coordenadas dadas.
// Las coordenadas deben ser coordenadas de mundo, no de pantalla.
// Devuelve -1 si no se ha encontrado ninguna entidad en dichas coordenadas.
func (em *EntityManager) GetEntityIDAt(x int, y int) int {
    for id, entity := range em.entities {
        if entity.CheckPosition(x, y) {
            return id
        }
    }
    // No se ha encontrado ninguna entidad.
    return -1
}

// Esta función de edición selecciona cambia la selección de la entidad a hacer aparecer.
func (em *EntityManager) ScrollEntity(dir int) {
    // Normalizamos la dirección
    var dScroll int = 0
    if dir < 0 {
        dScroll = -1
    } else if dir > 0 {
        dScroll = 1
    }

    em.entitySelected += dScroll
    // Si es menor que 0, volvemos al final
    if em.entitySelected < 0 {
        em.entitySelected = len(em.availableEntities) - 1
    }
    // Nos aseguramos 
    em.entitySelected = em.entitySelected % len(em.availableEntities)
}

// Hace spawn de un elemento según su tipo
func (em *EntityManager) SpawnByType(x int, y int, e_type int) error {
    // Hardcodeamos la manera de inicializar cada entidad.
    switch e_type {
    case ENTITY_BOX:
        // La caja será un SolidSprite cuando esté implementado, por el momento sólo es una entidad normal
        entity, err := em.CreateSprite(x, y, BOX_IDLE)
        if err != nil {
            return err
        }
        imgW, imgH := entity.GetImage().Size()
        // Center the image on the cursor
        entity.Move(float64(-imgW/2), float64(-imgH/2))
        em.SpawnEntity(entity)
    case ENTITY_BOMB:
        entity, err := em.CreateSprite(x, y, BOMB_OFF)
        if err != nil {
            return err
        }
        imgW, imgH := entity.GetImage().Size()
        // Center the image on the cursor.
        entity.Move(float64(-imgW/2), float64(-imgH/2))
        em.SpawnEntity(entity)
    case ENTITY_PLAYER:
        // De momento hardcodeamos el tamaño del hitbox al que sabemos que tiene el sprite.
        player := MakePlayer(em.resources[KING_FALL], float64(x-(78/2)), float64(y-(58/2)),
            0, 0, 78, 58, 100, 1.0, 0.8, 2.0, em.iManager)
        em.SpawnEntity(player)
    }
    return nil
}

// Devuelve el tipo de entidad seleccionada (que se puede cambiar con ScrollEntity)
func (em *EntityManager) GetSelectedEntityType() int {
    return em.entitySelected
}

// Devuelve una imagen representativa de la entidad seleccionada.
func (em *EntityManager) GetSelectedEntityImage() *ebiten.Image {
    selected := em.GetSelectedEntityType()
    return em.availableEntities[selected]
}

// Funciones de game loop

func (em *EntityManager) Update() error {
    for _, e := range em.entities {
        if err := e.Update(); err != nil {
            return err
        }
    }
    return nil
}
// FIXME: Error en el interior.
func (em *EntityManager) Draw(screen *ebiten.Image) {
    // FIXME: Esto no funciona, range no respeta ningún orden concreto y
    // si dos entidades se solapan, parpadean. Hay que usar algo que respete orden.
    for _, e := range em.entities {
        e.Draw(screen)
    }
}
