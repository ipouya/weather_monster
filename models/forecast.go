package models

type Forecast struct {
	CityId    int64     `json:"city_id"`
	Max  	  int   	`json:"max"`
	Min       int 		`json:"min"`
	Sample    int 		`json:"sample"`
}