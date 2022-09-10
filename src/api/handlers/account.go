package handlers

import (
	"fmt"
	"zkfmapf123/src/api/helpers"
	"zkfmapf123/src/api/services"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner string `form:"owner" json:"owner" binding:"required"`
	Currency string `form:"currency" json:"currency" binding:"required"`
	TimeZone string `form:"timezone" json:"timezone" binding:"required"`
}

func CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil{
		helpers.BadRequestResponse(ctx,err, "")
		return 
	}

	createErr := services.AccountCreate(ctx, req.Owner, req.Currency, req.TimeZone)
	if createErr != nil {
		return
	}

	user, findErr := services.AccountFindUser(ctx, req.Owner)
	if findErr != nil {
		return
	}

	helpers.SuccessResponse(ctx, user)
}

type getAccountRequest struct {
	ID int `uri:"id" binding:"required" min:"1"`
}

func GetAccount(ctx *gin.Context){
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		helpers.BadRequestResponse(ctx,err,"")
		return 
	}

	user, err := services.AccountFindUser(ctx, req.ID)
	if err != nil {
		return
	}

	helpers.SuccessResponse(ctx, user)
}

type listAccountRequest struct {
	Offset string `form:"offset" binding:"required"`
	Limit string `form:"limit" binding:"required"`
}

func ListAccount(ctx *gin.Context) {	
	var req listAccountRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		helpers.BadRequestResponse(ctx,err,"")
		return
	}

	userList, listErr := services.AccountListUser(ctx, req.Limit, req.Offset)
	if listErr != nil {
		return
	}

	fmt.Println(userList)

	helpers.SuccessResponse(ctx, userList)
}