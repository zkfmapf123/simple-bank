package main

import (
	"log"
	"zkfmapf123/src/api"
	"zkfmapf123/src/utils"
)

func main(){

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	serverErr := api.Start(config.ServerAddress)
	if serverErr != nil{
		log.Fatal("cannot start server : ", err.Error())	
	}
}