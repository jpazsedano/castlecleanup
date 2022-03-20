package gameutils

/* En esta clase se implementa un controlador 
 */

// Las estructuras que implementen EntityController permiten hacer de
// capa de abstracción entre las entidades y el algoritmo que controla
// ya sea input o una IA.
type EntityController interface {
    Update()
    GetActions() []Action
}

// Esta clase implementa un controlador que captura las acciones del mando
// realizadas durante un frame y 
type InputEntityController struct {
    actionsDone []Action
}

func MakeInputEntityController() {
    
}
