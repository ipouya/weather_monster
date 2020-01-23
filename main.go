package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"weather-monster/infrastructure"
	"weather-monster/models"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	_cityHttpDeliver "weather-monster/city/delivery/http"
	_cityRepo "weather-monster/city/repository"
	_cityUcase "weather-monster/city/usecase"

	_temperatureHttpDeliver "weather-monster/temperature/delivery/http"
	_temperatureRepo "weather-monster/temperature/repository"
	_temperatureUcase "weather-monster/temperature/usecase"

	_webhookHttpDeliver "weather-monster/webhook/delivery/http"
	_webhookRepo "weather-monster/webhook/repository"
	_webhookUcase "weather-monster/webhook/usecase"
)

var config = models.Config{}


func main() {

	if err := configor.Load(&config, "config.yml");err !=nil{
		panic(err)
	}
	fmt.Println(config)
	db := infrastructure.InitSqlDB(&config)
	defer db.Close()

	router := gin.Default()

	cityRepo := _cityRepo.NewSqlCityRepository(db)
	temperatureRepo := _temperatureRepo.NewSqlTemperatureRepository(db)
	webhookRepo := _webhookRepo.NewSqlWebhookRepository(db)

	cityUcase := _cityUcase.NewCityUsecase(cityRepo)
	_cityHttpDeliver.NewCityHandler(router, cityUcase)

	webhookUcase := _webhookUcase.NewWebhookUsecase(webhookRepo)
	_webhookHttpDeliver.NewWebhookHandler(router, webhookUcase)

	temperatureUcase := _temperatureUcase.NewTemperatureUsecase(temperatureRepo,webhookUcase)
	_temperatureHttpDeliver.NewTemperatureHandler(router, temperatureUcase)

	router.Run(config.Server.Address)
}
