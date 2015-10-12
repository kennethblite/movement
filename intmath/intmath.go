package intmath

import (
	"image"
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Absfloat(x float64) float64{
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func Find_xy_vel(current, target image.Point, speed int) image.Point {
	delta_x := target.X - current.X
	delta_y := target.Y - current.Y
	length := float64(speed)/math.Sqrt(float64(delta_x*delta_x + delta_y*delta_y))
	dx := int(length*float64(delta_x))
	dy := int(length*float64(delta_y))
	/*if delta_x == 0 {
		return image.Pt(0,speed*Sign(delta_y))
	}
	if delta_y == 0{
		return image.Pt(speed*Sign(delta_x),0)
	}*/
	/*
	slope := Absfloat(float64(delta_y) / float64(delta_x))
	delta_x = int(float64(Sign(delta_x)) * (float64(speed) / (slope + 1)))
	delta_y = int((float64(speed) - float64(Abs(delta_x))) * float64(Sign(delta_y)))
	*/
	return image.Pt(dx, dy)
}

func Pos(x int) bool {
	if x < 0 {
		return false
	}
	return true
}
