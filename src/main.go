package main

import (
	"log"
	"zkfmapf123/src/api"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main(){
	err := api.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start server : ", err.Error())	
	}
}