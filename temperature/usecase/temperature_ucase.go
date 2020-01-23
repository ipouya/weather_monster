package usecase

import (
	"fmt"
	"time"
	"weather-monster/temperature"
	"weather-monster/models"
	"weather-monster/webhook"
)

type temperatureUsecase struct {
	temperatureRepo    temperature.Repository
	webhookUsecase    	   webhook.Usecase
}

func NewTemperatureUsecase(t temperature.Repository,w webhook.Usecase) temperature.Usecase {
	return &temperatureUsecase{
		temperatureRepo:    t,
		webhookUsecase:		w,
	}
}

func (u *temperatureUsecase) Create(model *models.Temperature) (*models.Temperature, error) {
	model.CreatedAt = time.Now().Unix()
	res, err := u.temperatureRepo.Create(model)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	go u.webhookUsecase.PushData(model.CityId,res)

	return res, nil
}

func (u *temperatureUsecase) Forecast(cityId int64) (*models.Forecast, error) {

	res, err := u.temperatureRepo.Forecast(cityId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}