package listener 

import ( 
	"movement/entity"
	"azul3d.org/keyboard.v1"
)

type KeyListener interface{
	ProcessKey(keyboard.TypedEvent)
	Register(entity.Actor)
	Unregister(entity.Actor)
}

//This stores a list of what each actor wants to do when the event happens
type KeyboardListener struct{
	Listenermap map[string]entity.Actor
}

func NewKeyListener()*KeyboardListener{
	temp := make(map[string]entity.Actor)
	return &KeyboardListener{temp}
}
//find the map of the actions to perform for the mentioned key, then do all associated with that key.
func (k KeyboardListener) ProcessKey( key keyboard.StateEvent){
	for _, v := range k.Listenermap{
		v.ProcessKey(key)
	}
}

func (k *KeyboardListener) Register(e entity.Actor){
	k.Listenermap[e.Id()] = e
}


func (k *KeyboardListener) UnRegister(e entity.Actor){
	delete(k.Listenermap, e.Id())
}
