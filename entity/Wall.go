package entity

import (
	"image"
	"math"
)

//If no other neccessary method is added, consider deleting this

type Wall struct {
	*Simple
}

func (f Wall) Interact(i Interactor) {
	switch interactor := i.(type) {
	case *Player:
		rect, _ := interactor.DrawInfo()
		wall_rect, _ := f.DrawInfo()
		if wall_rect.Overlaps(rect) {
			move_point, x_edge := boundary_move(wall_rect, rect)
			if x_edge {
				interactor.X_frame_force = math.MaxInt64
			} else {
				interactor.Y_frame_force = math.MaxInt64
			}
			//if the rectangle intersection goes through the entire block
			interactor.Move(move_point)
		}
	default:
	}
}

func (f Wall) Priority() int {
	return 100
}

//Used to determine which edge the squares are touching each other at.
//the first boolean is which edge(x or Y) while travelling
//the second boolean is whether or not the thing is passed through to the other side
//f in this case is the wall image.Rectangle

func boundary_move(f image.Rectangle, other image.Rectangle) (image.Point, bool) {
	point := f.Intersect(other).Size()
	if point.X >= point.Y {
		if f.Min.Y >= other.Min.Y {
			return image.Point{0, -point.Y}, false
		} else {
			return image.Point{0, point.Y}, false
		}
	} else {
		if f.Min.X >= other.Min.X {
			return image.Point{-point.X, 0}, true
		} else {
			return image.Point{point.X, 0}, true
		}
	}
}

func (f Wall) ProcessInteract() {}

//the wall does not move
