package helper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prathishbv/notes-api/data/response"
)

func GetIDParam(ctx *gin.Context, paramName string) int {
	param := ctx.Param(paramName)
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid ID parameter"))
		ctx.Abort()
	}
	return id
}
