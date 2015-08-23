package Entity

import(
	"azul3d.org/gfx.v1"
	"image"
	"azul3d.org/keyboard.v1"
)
type PlayerSimple struct{
Simple
}

func NewPlayerSimple(id string,shape image.Rectangle, color gfx.Color)*PlayerSimple{
	return &PlayerSimple{Simple{id, shape, color}}
}

func (s *PlayerSimple) ProcessKey(key keyboard.TypedEvent){
	switch key.Rune{
		case 'd':
			s.Move(image.Point{2,0})
		case 'a':
			s.Move(image.Point{-2,0})
		case 'w':
			s.Move(image.Point{0,-2})
		case 's':
			s.Move(image.Point{0,2})
		default:
	}
}

func (s *PlayerSimple)Move(p image.Point){
	s.shape = s.shape.Add(p)
}

func (s *PlayerSimple) Interact(i Interactor){
//TODO: Unimplemented Interact method
}

func (s PlayerSimple) Priority()int{
	return 50
}
