package Entity

import(
	"azul3d.org/gfx.v1"
	"image"
	"azul3d.org/keyboard.v1"
	"time"
)
type PlayerSimple struct{
Simple
last_frame time.Time
controls map[keyboard.Key]keyboard.State
x_vel float64
y_vel float64
}

func NewPlayerSimple(id string,shape image.Rectangle, color gfx.Color)*PlayerSimple{
	temp := make(map[keyboard.Key]keyboard.State)
	return &PlayerSimple{Simple{id, shape, color},time.Now(), temp, 0, 0}
}

func (s *PlayerSimple) ProcessKey(key keyboard.StateEvent){
	switch key.Key{
		case keyboard.D,keyboard.A,keyboard.W,keyboard.S:
			controls[key.Key] = key.State 
	}
}

func (s *PlayerSimple)Move(p image.Point){
	s.shape = s.shape.Add(p)
}

func (s *PlayerSimple) Process(){
	
}
func (s *PlayerSimple) Interact(i Interactor){
//TODO: Unimplemented Interact method
}

func (s PlayerSimple) Priority()int{
	return 50
}
