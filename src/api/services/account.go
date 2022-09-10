package services

import (
	"zkfmapf123/src/api/helpers"
	"zkfmapf123/src/repository"

	"github.com/gin-gonic/gin"
)

func AccountCreate(ctx *gin.Context, owner string, currency string, tz string) error{
	err := repository.NewAccount(owner, 0, currency, tz).CreateAccount()
	if err != nil {
		helpers.BadRequestResponse(ctx , err, "")
		return err
	}

	return nil
}

func AccountFindUser[T string | int](ctx *gin.Context, param T) (any, error){
	user, err := repository.GetAccount(param)
	if err != nil {
		helpers.BadRequestResponse(ctx, err, helpers.ALREADY_EXISTS_OWNER)
		return nil, err
	}

	return user, nil
}

func AccountListUser(ctx *gin.Context, limit string, offset string) (any, error) {
	users, err := repository.GetAccountList(limit, offset)

	if err != nil{
		helpers.BadRequestResponse(ctx, err, "")
		return nil, err
	}

	return users, nil
}