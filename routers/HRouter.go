package routers

import (
	"github.com/VitorC-sa/crud-meplux/controllers/clientController"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/client", clientController.GET)
	r.GET("/client", clientController.GETALL)
	r.POST("/client", clientController.POST)
	r.PUT("/client", clientController.PUT)

	r.Run()
}
