package repository

import (
	"github.com/jinzhu/gorm"
	"weather-monster/temperature"
	"weather-monster/models"
)


type sqlTemperatureRepository struct {
	DB *gorm.DB
}

func NewSqlTemperatureRepository(DB *gorm.DB) temperature.Repository {
	return &sqlTemperatureRepository{DB}
}

func (r *sqlTemperatureRepository) Create(model *models.Temperature) (*models.Temperature, error) {
	if err := r.DB.Create(model).Error; err != nil {
		return nil, err
	}else {
		return model , nil
	}
}
func (r *sqlTemperatureRepository) Forecast(cityId int64) (*models.Forecast, error) {
	var forecast =  models.Forecast{CityId:cityId}

	if err := r.DB.Raw(`SELECT FLOOR(avg(max)) as max,
       					FLOOR(avg(min)) as min,
 						count(id) as sample
                  FROM temperatures
				  WHERE city_id = ? and TO_TIMESTAMP(created_at) >= NOW() - interval '1 day'`, cityId).Scan(&forecast).Error; err != nil {
		return nil, err
	}else{
		return &forecast,nil
	}
}

