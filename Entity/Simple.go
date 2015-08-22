package Entity

import(
	"azul3d.org/gfx.v1"
	"image"
	"azul3d.org/keyboard.v1"
)

type Simple struct {
id string
shape image.Rectangle
color gfx.Color
}

func NewSimple(id string,shape image.Rectangle, color gfx.Color)*Simple{
	return &Simple{id, shape, color}
}

func(s Simple) Draw()(image.Rectangle, gfx.Color){
return s.shape, s.color
}

func(s Simple) Id()string{
return s.id
}

type PlayerSimple struct{
Simple
}

func NewPlayerSimple(id string,shape image.Rectangle, color gfx.Color)*PlayerSimple{
	return &PlayerSimple{Simple{id, shape, color}}
}

func (s *PlayerSimple) ProcessKey(key keyboard.TypedEvent){
	switch key.Rune{
		case 'd':
			s.MoveRight()
		default:
	}
}

func (s *PlayerSimple)MoveRight(){
	s.shape = s.shape.Add(image.Point{2, 0})
}

