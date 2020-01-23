package webhook

import (
	"weather-monster/models"
)


type Usecase interface {

	Create(model *models.Webhook) (*models.Webhook, error)
	DeleteByID(id int64) (*models.Webhook, error)
	GetList(model *models.Webhook) ([]models.Webhook, error)
	PushData(cityId int64,temperature *models.Temperature)
}
