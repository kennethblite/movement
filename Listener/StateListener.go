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

//if anything is touching, the higher priority process is given
func (p PositionListener) Process(){
	for id, inter := range p.Interactormap{
		for id2, inter2 := range p.Interactormap{
			if id != id2{
				if inter.Priority() > inter2.Priority(){
					inter.Interact(inter2)
				}else{
					inter2.Interact(inter)
				}
			}
		}
	}
}

func (p *PositionListener) RegisterForState(i Entity.Interactor){
	p.Interactormap[i.Id()] = i
}


func (p *PositionListener) UnRegisterForState(i Entity.Interactor){
	delete(p.Interactormap, i.Id())
}
