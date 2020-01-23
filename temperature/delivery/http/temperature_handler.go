package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weather-monster/models"
	"weather-monster/temperature"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TemperatureHandler struct {
	TUsecase temperature.Usecase
}

func NewTemperatureHandler(r *gin.Engine, us temperature.Usecase) {
	handler := &TemperatureHandler{
		TUsecase: us,
	}

	r.POST("/temperatures", handler.Create)
	r.GET("/forecasts/:city_id", handler.ForecastByCityId)
}


func (a *TemperatureHandler) Create(c *gin.Context) {

	var json models.Temperature
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	obj, err := a.TUsecase.Create(&json)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}

func (a *TemperatureHandler) ForecastByCityId(c *gin.Context) {

	cityId, err := strconv.ParseInt(c.Params.ByName("city_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	obj, err := a.TUsecase.Forecast(cityId)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}
