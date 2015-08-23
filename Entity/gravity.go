package Entity

import(
	"image"
)


type Gravity struct{
	*Simple
}

func(g Gravity) Interact(i Interactor){
	switch interactor := i.(type){
		case *PlayerSimple:
			interactor.Move(image.Point{0,2})
		default:
	}
}

func (g Gravity) Priority()int{
	return 75
}
