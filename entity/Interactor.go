package entity

import ()

//Used for anything that contains interactive logic
//Interact creates a queue of actions, that are resolved by ProcessInteract
type Interactor interface {
	Entity
	Interact(Interactor)
	ProcessInteract()
}
