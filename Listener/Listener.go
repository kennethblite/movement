package Listener 

import ( 
	"movement/Entity"
	"azul3d.org/keyboard.v1"
)

type KeyListener interface{
	ProcessKey(keyboard.TypedEvent)
	Register(Entity.Actor)
	Unregister(Entity.Actor)
}

//This stores a list of what each actor wants to do when the event happens
type KeyboardListener struct{
	Listenermap map[string]Entity.Actor
}

func NewKeyListener()*KeyboardListener{
	temp := make(map[string]Entity.Actor)
	return &KeyboardListener{temp}
}
//find the map of the actions to perform for the mentioned key, then do all associated with that key.
func (k KeyboardListener) ProcessKey( key keyboard.TypedEvent){
	for _, v := range k.Listenermap{
		v.ProcessKey(key)
	}
}

func (k *KeyboardListener) Register(e Entity.Actor){
	k.Listenermap[e.Id()] = e
}


func (k *KeyboardListener) UnRegister(e Entity.Actor){
	delete(k.Listenermap, e.Id())
}
