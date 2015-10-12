package entity

import (
	"image"
	"movement/intmath"
)

type Enemy interface {
	Interactor
}

type EnemyPatrol struct {
	*Simple
	start         image.Point
	end           image.Point
	towards_start bool
}

const (
	enemy_patrol_speed = 6
)

//There is not a beginning because it is the point that the enemy spawns at
func NewEnemyPatrol(s *Simple, end, start image.Point) *EnemyPatrol {
	return &EnemyPatrol{s, start, end, true}
}

func (e *EnemyPatrol) Interact(i Interactor) {
	switch interactor := i.(type) {
	case *Player:
		rect, _ := interactor.DrawInfo()
		enemy_rect, _ := e.DrawInfo()
		if enemy_rect.Overlaps(rect) {
			min_point := rect.Min
			//to be changed with spawn location for future implementations
			interactor.Move(image.Pt(-min_point.X+100, -min_point.Y+100))
		}
	}
}

func (e *EnemyPatrol) ProcessInteract() {
	rect, _ := e.DrawInfo()
	if e.towards_start {
		if e.start.In(rect) {
			e.towards_start = false
		} else {
			e.Move(intmath.Find_xy_vel(e.shape.Min, e.start, enemy_patrol_speed))
		}
	} else {
		if e.end.In(rect) {
			e.towards_start = true
		} else {
			e.Move(intmath.Find_xy_vel(e.shape.Max, e.end, enemy_patrol_speed))
		}
	}
}

func (e *EnemyPatrol) Move(p image.Point) {
	e.shape = e.shape.Add(p)
}
