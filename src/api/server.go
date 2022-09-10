package api

import (
	"net/http"
	"zkfmapf123/src/api/handlers"

	"github.com/gin-gonic/gin"
)

const (
	BASE = "/"
	ACCOUNT = "/account"
	FIND_ACCOUT_USER = "/account/:id"
)

func Start(address string) error{
	r := gin.Default()

	// middlewares

	// routers
	router(r)

	// add routes to router
	return r.Run(address)
}


func router(r *gin.Engine) {
	r.GET(BASE, func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{"status" : "Hello world"})
		return
	})

	r.POST(ACCOUNT, handlers.CreateAccount)
	r.GET(FIND_ACCOUT_USER, handlers.GetAccount)
	r.GET(ACCOUNT, handlers.ListAccount)
}

func middleware(){
	
}
