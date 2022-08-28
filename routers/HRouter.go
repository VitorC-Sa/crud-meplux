package routers

import (
	"net/http"

	"github.com/VitorC-sa/crud-meplux/controllers/clientController"
	"github.com/VitorC-sa/crud-meplux/controllers/productController"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	//client
	{
		r.GET("/client/:id", clientController.GET)
		r.GET("/client", clientController.GETALL)
		r.POST("/client", clientController.POST)
		r.PUT("/client/:id", clientController.PUT)
	}

	//product
	{
		r.GET("/product/:id", productController.GET)
		r.GET("/product", productController.GETALL)
		r.POST("/product", productController.POST)
		r.PUT("/product/:id", productController.PUT)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Página não encontrada": "404",
		})
	})
	r.Run()
}
