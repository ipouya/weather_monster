package models

type City struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required" gorm:"not null;unique_index"`
	Latitude  float64   `json:"latitude"`
	Longitude float64 	`json:"longitude"`
}