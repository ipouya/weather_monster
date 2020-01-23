package testHelper

import (
	"fmt"
	"math/rand"
	"time"
	"weather-monster/models"
)

func MakeCity() *models.City{
	item := &models.City{
		Name: fmt.Sprint(time.Now().Unix()),
		Latitude: (rand.Float64() * 5) + 5,
		Longitude: (rand.Float64() * 5) + 5,
	}
	testDB.Create(&item)
	return item
}
func MakeTemperature(cityID int64) *models.Temperature{
	item := &models.Temperature{
		CityId:cityID,
		Max: rand.Intn(100),
		Min: rand.Intn(100),
		CreatedAt: time.Now().Unix(),
	}
	testDB.Create(&item)
	return item
}

func MakeWebHook(cityID int64,callbaclUrl string) *models.Webhook{
	item := &models.Webhook{
		CityId:cityID,
		CallbackUrl:callbaclUrl,
	}
	testDB.Create(&item)
	return item
}

