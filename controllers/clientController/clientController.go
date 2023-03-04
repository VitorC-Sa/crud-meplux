package clientController

import (
	"net/http"
	"strconv"

	c "github.com/VitorC-sa/crud-meplux/models/clientModel"
	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"codigo":   "Requisicao Invalida",
			"mensagem": "ID fora do formato esperado",
		})
		return
	}

	c := c.GetOne(id)
	if c == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"codigo":   "Entidade invalida",
			"mensagem": "Cliente nao encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, c)
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
