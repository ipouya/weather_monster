package repository

import (
	"github.com/jinzhu/gorm"
	"weather-monster/webhook"
	"weather-monster/models"
)


type sqlWebhookRepository struct {
	DB *gorm.DB
}

func NewSqlWebhookRepository(DB *gorm.DB) webhook.Repository {
	return &sqlWebhookRepository{DB}
}

func (r *sqlWebhookRepository) Get(model *models.Webhook) (res *models.Webhook, err error) {
	var item models.Webhook
	if err := r.DB.Where(model).First(&item).Error; err != nil {
		return nil, err
	}else {
		return &item , nil
	}
}
func (r *sqlWebhookRepository) GetList(model *models.Webhook) ([]models.Webhook, error) {
	var items []models.Webhook
	if err := r.DB.Where(model).Find(&items).Error; err != nil {
		return nil, err
	}else {
		return items , nil
	}
}
func (r *sqlWebhookRepository) Create(model *models.Webhook) (*models.Webhook, error) {
	if err := r.DB.Create(model).Error; err != nil {
		return nil, err
	}else {
		return model , nil
	}
}

func (r *sqlWebhookRepository) Delete(model *models.Webhook) (*models.Webhook, error) {
	existedData,err := r.Get(model)
	if err != nil{
		return nil,err
	}
	if err := r.DB.Where(model).Delete(model).Error; err != nil {
		return nil, err
	}else {
		return existedData , nil
	}
}
