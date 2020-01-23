package models

type Webhook struct {
	ID        	int64     `json:"id" gorm:"primary_key"`
	CityId    	int64     `json:"city_id" binding:"required" gorm:"not null"`
	CallbackUrl string 	  `json:"callback_url" binding:"required" gorm:"not null"`
}