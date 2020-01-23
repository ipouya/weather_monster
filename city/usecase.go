package city

import (
	"weather-monster/models"
)


type Usecase interface {

	GetByID(id int64) (*models.City, error)
	Create(model *models.City) (*models.City, error)
	Patch(model *models.City) (*models.City, error)
	DeleteByID(id int64) (*models.City, error)
}
