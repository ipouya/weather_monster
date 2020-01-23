package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weather-monster/city"
	"weather-monster/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CityHandler struct {
	CUsecase city.Usecase
}

func NewCityHandler(r *gin.Engine, us city.Usecase) {
	handler := &CityHandler{
		CUsecase: us,
	}

	r.POST("/cities", handler.Create)
	r.GET("/cities/:id", handler.GetByID)
	r.PATCH("/cities/:id", handler.Patch)
	r.DELETE("/cities/:id", handler.Delete)
}


func (a *CityHandler) Create(c *gin.Context) {

	var json models.City
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	obj, err := a.CUsecase.Create(&json)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}
func (a *CityHandler) Patch(c *gin.Context) {

	var json models.City
	if err := c.Bind(&json); err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	json.ID = id
	obj, err := a.CUsecase.Patch(&json)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}
func (a *CityHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	obj, err := a.CUsecase.DeleteByID(id)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}
func (a *CityHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	obj, err := a.CUsecase.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON( http.StatusBadRequest,ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, obj)
}

