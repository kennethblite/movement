package Entity

import(
	"image"
	"azul3d.org/gfx.v1"
)
type Entity interface{
Draw()(image.Rectangle, gfx.Color)
Id()string
} 
