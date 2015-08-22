package Entity

import(
	"azul3d.org/keyboard.v1"
)

//Actor will be used for anything that depends on a key press
type Actor interface{
Entity
ProcessKey(keyboard.TypedEvent)
}
//Register and Unregister the event with the master Listener
