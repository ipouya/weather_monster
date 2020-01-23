package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"weather-monster/models"
	"weather-monster/webhook"
)

type webhookUsecase struct {
	webhookRepo    webhook.Repository
}

func NewWebhookUsecase(w webhook.Repository) webhook.Usecase {
	return &webhookUsecase{
		webhookRepo:    w,
	}
}

func (u *webhookUsecase) GetList(model *models.Webhook) ([]models.Webhook, error) {

	res, err := u.webhookRepo.GetList(model)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *webhookUsecase) Create(model *models.Webhook) (*models.Webhook, error) {

	res, err := u.webhookRepo.Create(model)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *webhookUsecase) DeleteByID(id int64) (*models.Webhook, error) {

	res, err := u.webhookRepo.Delete(&models.Webhook{ID: id})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *webhookUsecase) PushData(cityId int64,temperature *models.Temperature) {
	if hooks, err := u.GetList(&models.Webhook{CityId: cityId});err != nil{
		fmt.Println(err.Error())
	}else{
		for _, wh := range hooks {
			callWebhook(&wh,temperature)
		}
	}
}

func callWebhook(wg *models.Webhook,temperature *models.Temperature ) {
	json,err := json.Marshal(temperature)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := http.Post(wg.CallbackUrl,"application/json",bytes.NewBuffer(json))
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}

