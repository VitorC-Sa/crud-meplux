package clientController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "OK",
	})
}

func GETALL(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "OK",
	})
}

func POST(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "OK",
	})
}

func PUT(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "OK",
	})
}
