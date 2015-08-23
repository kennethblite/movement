package Entity

import(
	"image"
	"azul3d.org/gfx.v1"
)
//Entity is anything that needs to be drawn, including backgrounds
type Entity interface{
//DrawInfo is also used in Bounds Calculations for interactors.
DrawInfo()(image.Rectangle, gfx.Color)
Id()string
} 
