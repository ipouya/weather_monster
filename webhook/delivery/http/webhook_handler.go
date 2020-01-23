package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weather-monster/models"
	"weather-monster/webhook"
)

type ResponseError struct {
	Message string `json:"message"`
}

type WebhookHandler struct {
	CUsecase webhook.Usecase
}

func NewWebhookHandler(r *gin.Engine, us webhook.Usecase) {
	handler := &WebhookHandler{
		CUsecase: us,
	}

	r.POST("/webhooks", handler.Create)
	r.DELETE("/webhooks/:id", handler.Delete)
}

func (a *WebhookHandler) Create(c *gin.Context) {

	var json models.Webhook
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

func (a *WebhookHandler) Delete(c *gin.Context) {

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

