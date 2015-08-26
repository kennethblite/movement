package level

import(
	"movement/entity"
	"image"
)

type Level interface{
	getEntities()[]entity.Entity
	getInteractors()[]entity.Interactor
	getStartPoint()image.Point
	Load(string)
}


type level struct{
	interactor
}

func GetCurrentLevel() Level{
	return Current_Level
}
