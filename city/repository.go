package city

import (
	"weather-monster/models"
)

type Repository interface {

	Create(model *models.City) (*models.City, error)
	Get(model *models.City) (*models.City, error)
	Patch(model *models.City) (*models.City, error)
	Delete(model *models.City) (*models.City, error)
}
