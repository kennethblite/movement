package Entity

import(
	"azul3d.org/gfx.v1"
	"image"
	"azul3d.org/keyboard.v1"
	"time"
	"math"
)
type PlayerSimple struct{
Simple
last_frame time.Time
controls map[keyboard.Key]keyboard.State
x_vel int
y_vel int
X_frame_force int
Y_frame_force int
}
const(
	Move_force = 2
	Friction_force = 1
	Max_speed = 10
)

func NewPlayerSimple(id string,shape image.Rectangle, color gfx.Color)*PlayerSimple{
	temp := make(map[keyboard.Key]keyboard.State)
	return &PlayerSimple{Simple{id, shape, color},time.Now(), temp, 0, 0, 0, 0}
}

func (s *PlayerSimple) ProcessKey(key keyboard.StateEvent){
	switch key.Key{
		case keyboard.D,keyboard.A,keyboard.W,keyboard.S:
			s.controls[key.Key] = key.State 
	}
}

func (s *PlayerSimple)Move(p image.Point){
	s.shape = s.shape.Add(p)
}

func (s *PlayerSimple) ProcessInteract(){
	//assuming mass of 1kg force is the same as Accelleration
	if (s.X_frame_force == math.MaxInt64){
		s.x_vel = 0
	}else{
	if s.controls[keyboard.D] == keyboard.Down{
		s.X_frame_force += Move_force
	}
	if s.controls[keyboard.A] == keyboard.Down{
		s.X_frame_force -= Move_force
	}
	if s.controls[keyboard.A] == keyboard.Up && s.controls[keyboard.D] == keyboard.Up{
		if s.x_vel > 0{
		s.X_frame_force -= Friction_force
		}else if s.x_vel < 0 {
		s.X_frame_force += Friction_force
		}
	}
	s.x_vel +=  s.X_frame_force
	if s.x_vel >= Max_speed {
		s.x_vel = Max_speed
	}
	if s.x_vel < -Max_speed{
		s.x_vel = -Max_speed
	}
	}
	if (s.Y_frame_force == math.MaxInt64){
		s.y_vel = 0
	}else{
	if s.controls[keyboard.W] == keyboard.Down{
		s.Y_frame_force -= Move_force
	}
	if s.controls[keyboard.S] == keyboard.Down{
		s.Y_frame_force += Move_force
	}
	if s.controls[keyboard.W] == keyboard.Up && s.controls[keyboard.S] == keyboard.Up{
		if s.y_vel > 0{
		s.Y_frame_force -= Friction_force
		}else if s.y_vel < 0 {
		s.Y_frame_force += Friction_force
		}
	}
	s.y_vel +=  s.Y_frame_force
	if s.y_vel >= Max_speed {
		s.y_vel = Max_speed
	}
	if s.y_vel < -Max_speed{
		s.y_vel = -Max_speed
	}
	}
	Delta_Position :=  image.Point{s.x_vel,s.y_vel}
	s.Move(Delta_Position)
	s.X_frame_force = 0
	s.Y_frame_force = 0
}
func (s *PlayerSimple) Interact(i Interactor){
//TODO: Unimplemented Interact method
}

func (s PlayerSimple) Priority()int{
	return 50
}
