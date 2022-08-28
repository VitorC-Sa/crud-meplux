package routers

import (
	"github.com/VitorC-sa/crud-meplux/controllers/clientController"
	"github.com/VitorC-sa/crud-meplux/controllers/productController"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	{
		r.GET("/client/:id", clientController.GET)
		r.GET("/client", clientController.GETALL)
		r.POST("/client", clientController.POST)
		r.PUT("/client/:id", clientController.PUT)
	}

	{
		r.GET("/product/:id", productController.GET)
		r.GET("/product", productController.GETALL)
		r.POST("/product", productController.POST)
		r.PUT("/product/:id", productController.PUT)
	}

	r.Run()
}
