package testHelper

import (
	"github.com/jinzhu/gorm"
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

var config  models.Config
var testDB  *gorm.DB

func StartTestServer() *gin.Engine{

	config.Sql.Type = "postgres"
	config.Sql.Host = "localhost"
	config.Sql.User = "test"
	config.Sql.DbName = "test_db"
	config.Sql.Password = "12345"
	config.Sql.Port = 5432
	config.Sql.Sslmode = "disable"

	testDB = infrastructure.InitSqlDB(&config)

	router := gin.Default()

	cityRepo := _cityRepo.NewSqlCityRepository(testDB)
	temperatureRepo := _temperatureRepo.NewSqlTemperatureRepository(testDB)
	webhookRepo := _webhookRepo.NewSqlWebhookRepository(testDB)

	cityUcase := _cityUcase.NewCityUsecase(cityRepo)
	_cityHttpDeliver.NewCityHandler(router, cityUcase)

	webhookUcase := _webhookUcase.NewWebhookUsecase(webhookRepo)
	_webhookHttpDeliver.NewWebhookHandler(router, webhookUcase)

	temperatureUcase := _temperatureUcase.NewTemperatureUsecase(temperatureRepo,webhookUcase)
	_temperatureHttpDeliver.NewTemperatureHandler(router, temperatureUcase)

	return router
}
