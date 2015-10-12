package level

import (
	"encoding/json"
	"image"
	"io/ioutil"
	"movement/entity"
)

var Current_Level Level

type Level interface {
	Entities() []entity.Entity
	Interactors() []entity.Interactor
	StartPoint() image.Point
	EndPoint() image.Point
}

type level struct {
	entities               []entity.Entity
	interactors            []entity.Interactor
	start_point, end_point image.Point
}

func newLevel() *level {
	a := make([]entity.Entity)
	b := make([]entity.Interactor)
	start := image.Pt(100, 100)
	end := image.Pt(600, 600)
	return &level{a, b, start, end}
}
func (l level) Entities() []entity.Entity {
	return l.entities
}

func (l level) Interactors() []entity.Interactor {
	return l.interactors
}

func (l level) StartPoint() image.Point {
	return l.start_point
}

func (l level) EndPoint() image.Point {
	return l.end_point
}

func Load(name string) {
	base_level := newLevel()
	entity_stream, err := ioutil.ReadFile("rsc/" + name + "/entities.txt")
	if err != nil {
		fmt.Println("Failed to load level")
	}
		Current_level = base_level

	//ideally, this would marshal out the level Json files, and populate it into a new Current level. While Unregistering everything
}

func GetCurrentLevel() Level {
	return Current_Level
}
