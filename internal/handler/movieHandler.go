package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/looksaw/go_greenlight/internal/types"
)

// CreateMovieHandler
func (controller *Controller) CreateMovieHandler(c *gin.Context) {
	var req types.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrMovieResponse{
			Status:      http.StatusBadRequest,
			Err:         err,
			Description: "Your send a bad Request",
		})
	}
	//TODO 传入service
	res := controller.Service.CreateMovie(req)
	c.JSON(http.StatusCreated, res)
}

// ShowMovieHandler
func (controller *Controller) ShowMovieByIdHandler(c *gin.Context) {
	stdStr := c.Param("id")
	var id int
	var err error
	if id, err = strconv.Atoi(stdStr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrMovieResponse{
			Status:      http.StatusBadRequest,
			Err:         err,
			Description: "Your ID Query is wrong",
		})
	}
	res, err := controller.Service.ShowMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrMovieResponse{
			Status:      http.StatusInternalServerError,
			Err:         err,
			Description: "Your ID out of bound",
		})
	}
	c.JSON(http.StatusOK, res)
}

func (controller *Controller) ShowMovieAll(c *gin.Context) {
	res := controller.Service.ShowMovieAll()
	c.JSON(http.StatusOK, res)
}

func (controller *Controller) UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	errMsg := fmt.Sprintf("id %d is not in bound", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrMovieResponse{
			Status:      http.StatusBadRequest,
			Err:         err,
			Description: errMsg,
		})
	}
	var req types.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrMovieResponse{
			Status:      http.StatusBadRequest,
			Err:         err,
			Description: "Your send a bad Request",
		})
	}
	res, err := controller.Service.UpdateMovieById(id, req)
	c.JSON(http.StatusOK, res)
}

func (controller *Controller) DeleteID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	errMsg := fmt.Sprintf("id %d is not in bound", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrMovieResponse{
			Status:      http.StatusBadRequest,
			Err:         err,
			Description: errMsg,
		})
	}
	res := controller.Service.DeleteMovieById(id)
	c.JSON(http.StatusNoContent, res)

}
