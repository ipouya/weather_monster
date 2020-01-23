package models

type Temperature struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	CityId    int64     `json:"city_id" binding:"required" gorm:"not null"`
	Max  	  int   	`json:"max" binding:"required" gorm:"not null"`
	Min       int 		`json:"min" binding:"required" gorm:"not null"`
	CreatedAt int64 	`json:"timestamp"`
}