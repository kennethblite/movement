package listener

import (
	"movement/entity"
)

type StateListener interface {
	Process()
	RegisterForState(entity.Interactor)
	UnRegisterForState(entity.Interactor)
}

func NewPositionListener() *PositionListener {
	temp := make(map[string]entity.Interactor)
	return &PositionListener{temp}
}

type PositionListener struct {
	Interactormap map[string]entity.Interactor
}

//Create a queue of forces between each other, then Process the interactions for each object
func (p PositionListener) Process() {
	//commenting out for speed testing
	for id, inter := range p.Interactormap {
		for id2, inter2 := range p.Interactormap {
			if id != id2 {
				inter.Interact(inter2)
			}
		}
	}
	for _, interactors := range p.Interactormap {
		interactors.ProcessInteract()
	}
}

func (p *PositionListener) RegisterForState(i entity.Interactor) {
	p.Interactormap[i.Id()] = i
}

func (p *PositionListener) UnRegisterForState(i entity.Interactor) {
	delete(p.Interactormap, i.Id())
}
