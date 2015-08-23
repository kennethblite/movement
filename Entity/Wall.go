package Entity

import(
	"image"
)

//If no other neccessary method is added, consider deleting this

type Wall struct{
	*Simple
}

func(f Wall) Interact(i Interactor){
	switch interactor := i.(type){
	case *PlayerSimple:
		rect, _ := interactor.DrawInfo()
		wall_rect, _ := f.DrawInfo()
		if wall_rect.Overlaps(rect){
			interactor.Move(boundary_move(wall_rect, rect))
		}
	default:
	}
}

func(f Wall) Priority()int{
	return 100
}

//Used to determine which edge the squares are touching each other at.
func boundary_move(f image.Rectangle, other image.Rectangle)image.Point{
	point := f.Intersect(other).Size()
	if point.X >= point.Y{
		if f.Min.Y >= other.Min.Y {
			return image.Point{0, -point.Y}
		}else{
			return image.Point{0, point.Y}
		}
	}else {
		if f.Min.X >= other.Min.X {
			return image.Point{-point.X, 0}
		}else{
			return image.Point{point.X, 0}
		}
	}
}
