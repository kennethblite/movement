package Entity

import()

//Used for anything that contains interactive logic
//Priority returns which Interact has precedence over the other when comparing two 
type Interactor interface{
	Entity
	Interact(Interactor)
	Process()
}
