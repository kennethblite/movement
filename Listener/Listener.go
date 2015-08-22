package Listener 

import ( 
	"movement/Entity"
	"azul3d.org/keyboard.v1"
	"reflect"
)

type KeyListener interface{
	ProcessKey(keyboard.TypedEvent, map[string]Entity.Entity)
	Register(keyboard.TypedEvent, string,func(*Entity.Actor)())
	Unregister(Entity.Actor)
}

//This stores a list of what each actor wants to do when the event happens
type KeyboardListener struct{
	Listenermap map[rune]map[string]func(e *Entity.Actor)()
}


//find the map of the actions to perform for the mentioned key, then do all associated with that key.
func (k KeyboardListener) ProcessKey( key keyboard.TypedEvent, rendering map[string]Entity.Entity){
	if actions, found := k.Listenermap[key.Rune]; found{
		for i, v := range actions{
			(reflect.TypeOf(rendering[i]))(rendering[i]).v()
		}
	}
}

func (k *KeyboardListener) Register(type_event keyboard.TypedEvent, id string, function func(e *Entity.Actor)()){
	k.Listenermap[type_event.Rune][id] = function
}


func (k *KeyboardListener) UnRegister(typed_event keyboard.TypedEvent, id string){
	delete(k.Listenermap[typed_event.Rune], id)
}
