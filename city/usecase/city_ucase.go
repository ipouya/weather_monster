package usecase

import (

	"weather-monster/city"
	"weather-monster/models"
)

type cityUsecase struct {
	cityRepo    city.Repository
}

func NewCityUsecase(r city.Repository) city.Usecase {
	return &cityUsecase{
		cityRepo:    r,
	}
}

func (u *cityUsecase) Create(model *models.City) (*models.City, error) {

	res, err := u.cityRepo.Create(model)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *cityUsecase) Patch(model *models.City) (*models.City, error) {

	res, err := u.cityRepo.Patch(model)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *cityUsecase) DeleteByID(id int64) (*models.City, error) {

	res, err := u.cityRepo.Delete(&models.City{ID: id})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *cityUsecase) GetByID(id int64) (*models.City, error) {

	res, err := u.cityRepo.Get(&models.City{ID: id})
	if err != nil {
		return nil, err
	}
	return res, nil
}

