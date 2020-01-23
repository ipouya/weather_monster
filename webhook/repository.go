package webhook

import (
	"weather-monster/models"
)

type Repository interface {

	Create(model *models.Webhook) (*models.Webhook, error)
	Get(model *models.Webhook) (*models.Webhook, error)
	GetList(model *models.Webhook) ([]models.Webhook, error)
	Delete(model *models.Webhook) (*models.Webhook, error)
}
