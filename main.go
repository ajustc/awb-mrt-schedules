package main

import (
	"github.com/ajustc/awb-mrt-schedules/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitializeRoutes()
}

func InitializeRoutes() {
	var router = gin.Default()
	var api = router.Group("/v1/api")

	station.Initiate(api)

	router.Run(":8080")
}
