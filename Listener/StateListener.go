package Listener


import(
	"movement/Entity"
)

type StateListener interface{
	Process()
	RegisterForState(Entity.Interactor)
	UnRegisterForState(Entity.Interactor)
}

func NewPositionListener()*PositionListener{
	temp:= make(map[string]Entity.Interactor)
	return &PositionListener{temp}
}

type PositionListener struct{
	Interactormap map[string]Entity.Interactor
}

//Create a queue of forces between each other, then Process the interactions for each object
func (p PositionListener) Process(){
//commenting out for speed testing
	for id, inter := range p.Interactormap{
		for id2, inter2 := range p.Interactormap{
			if id != id2{
					inter.Interact(inter2)
			}
		}
	}
	for _, interactors := range p.Interactormap{
			interactors.ProcessInteract()
	}
}

func (p *PositionListener) RegisterForState(i Entity.Interactor){
	p.Interactormap[i.Id()] = i
}


func (p *PositionListener) UnRegisterForState(i Entity.Interactor){
	delete(p.Interactormap, i.Id())
}
