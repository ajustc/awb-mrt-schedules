package station

import (
	"net/http"

	"github.com/ajustc/awb-mrt-schedules/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()
	station := router.Group("/station")

	station.GET("/", func(c *gin.Context) {
		GetAll(c, stationService)
	})
	station.GET("/:id", func(c *gin.Context) {
		GetByID(c, stationService)
	})
}

func GetAll(c *gin.Context, service Service) {
	data, err := service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  true,
		Message: "Success",
		Data:    data,
	})
}

func GetByID(c *gin.Context, service Service) {
	id := c.Param("id")
	data, err := service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Status:  true,
		Message: "Success get station by id",
		Data:    data,
	})
}
