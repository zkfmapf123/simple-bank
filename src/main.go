package main

import (
	"fmt"
	"zkfmapf123/src/base"
	"zkfmapf123/src/repository"
)

func main(){
	
	accounts := []base.AccountModels{}
	rows, err := repository.FindOne(accounts, "select * from accounts where owner=?","owner")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(rows)
}