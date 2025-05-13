package router

import (
	"github.com/gin-gonic/gin"
	"github.com/looksaw/go_greenlight/internal/handler"
	"github.com/looksaw/go_greenlight/internal/repository"
	"github.com/looksaw/go_greenlight/internal/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	service := service.NewInMemService(&repository.InMemRepository{})
	ctrl := handler.NewController(service)
	api := r.Group("/v1/api")
	{
		api.GET("/health", ctrl.HealthCheckerHandler)
		api.POST("/createMovie", ctrl.CreateMovieHandler)
		api.GET("/showMovie", ctrl.ShowMovieAll)
		api.GET("/showMovie/:id", ctrl.ShowMovieByIdHandler)
	}
	return r
}
