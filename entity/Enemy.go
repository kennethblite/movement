package entity

import(
	"image"	
)

type Enemy interface{
Interactor
}

type EnemyPatrol struct{
*Simple
 start image.Point
 end image.Point
 towards_start bool
}

const(
	enemy_patrol_speed = 6
)

//There is not a beginning because it is the point that the enemy spawns at
func NewEnemyPatrol(s *Simple, end, start image.Point )*EnemyPatrol{
	return &EnemyPatrol{s, start, end, true} 
}

func(e *EnemyPatrol) Interact(i Interactor){
	switch interactor := i.(type){
		case *Player:
			rect, _ := interactor.DrawInfo()
			enemy_rect ,_ := e.DrawInfo()
			if enemy_rect.Overlaps(rect){
				min_point := rect.Min
				//to be changed with spawn location for future implementations
				interactor.Move(image.Pt(-min_point.X+100, -min_point.Y+100))
			}
	}
}

func(e *EnemyPatrol) ProcessInteract(){
	rect, _ := e.DrawInfo()
	if e.towards_start{
		if(e.start.In(rect)){
			e.towards_start = false
		}else{
			e.Move(find_xy_vel(e.shape.Min, e.start))
		}
	}else{
		if e.end.In(rect){
			e.towards_start = true
		}else{
			e.Move(find_xy_vel(e.shape.Max, e.end))
		}
	}
}

func(e *EnemyPatrol) Move(p image.Point){
	e.shape = e.shape.Add(p)
}

func find_xy_vel(current, target image.Point)(image.Point){
	delta_x := target.X - current.X
	delta_y := target.Y - current.Y
	slope := abs(delta_y / delta_x)
	delta_x = sign(delta_x)*(enemy_patrol_speed/(slope+1))
	delta_y = (enemy_patrol_speed - abs(delta_x))*sign(delta_y)
	return image.Pt(delta_x, delta_y)
}

func abs(x int) int{
	if x < 0 {
	return -x
	}
	return x
}

func sign(x int) int{
	if x < 0 {
	return -1
	}else{
	return 1
	}
}
