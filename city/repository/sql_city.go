package repository

import (
	"github.com/jinzhu/gorm"
	"weather-monster/city"
	"weather-monster/models"
)


type sqlCityRepository struct {
	DB *gorm.DB
}

func NewSqlCityRepository(DB *gorm.DB) city.Repository {
	return &sqlCityRepository{DB}
}

func (r *sqlCityRepository) Get(model *models.City) (res *models.City, err error) {
	var item models.City
	if err := r.DB.Where(model).First(&item).Error; err != nil {
		return nil, err
	}else {
		return &item , nil
	}
}

func (r *sqlCityRepository) Create(model *models.City) (*models.City, error) {
	if err := r.DB.Create(model).Error; err != nil {
		return nil, err
	}else {
		return model , nil
	}
}
func (r *sqlCityRepository) Patch(model *models.City) (*models.City, error) {
	_,err := r.Get(&models.City{ID:model.ID})
	if err != nil{
		return nil,err
	}
	if err := r.DB.Model(&models.City{}).Where(&models.City{ID:model.ID}).Update(model).Error; err != nil {
		return nil, err
	}else {
		return model , nil
	}
}
func (r *sqlCityRepository) Delete(model *models.City) (*models.City, error) {
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
