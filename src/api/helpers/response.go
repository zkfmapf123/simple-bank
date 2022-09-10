package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(ctx *gin.Context, err error, msg string){
	ctx.JSON(http.StatusBadRequest, ErrorReason(err, msg))
}

func SuccessResponse(ctx *gin.Context, params any) {
	ctx.JSON(http.StatusOK, params)
}