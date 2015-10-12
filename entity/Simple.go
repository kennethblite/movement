package entity

import (
	"azul3d.org/gfx.v1"
	"image"
)

type Simple struct {
	id    string
	shape image.Rectangle
	color gfx.Color
}

func NewSimple(id string, shape image.Rectangle, color gfx.Color) *Simple {
	return &Simple{id, shape, color}
}

func (s Simple) DrawInfo() (image.Rectangle, gfx.Color) {
	return s.shape, s.color
}

func (s Simple) Id() string {
	return s.id
}
