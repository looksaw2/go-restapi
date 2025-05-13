package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/looksaw/go_greenlight/cmd/config"
	"github.com/looksaw/go_greenlight/internal/service"
)

type Controller struct {
	Service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

// health检查
func (controller *Controller) HealthCheckerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"description": "health OK",
		"status":      "aviable",
		"serviceName": config.GreenLightEnvelope.ServiceName,
		"enviroment":  config.GreenLightEnvelope.Data.Env,
	})
}
