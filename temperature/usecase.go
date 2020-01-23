package temperature

import (
	"weather-monster/models"
)


type Usecase interface {

	Create(model *models.Temperature) (*models.Temperature, error)
	Forecast(cityId int64) (*models.Forecast, error)
}
