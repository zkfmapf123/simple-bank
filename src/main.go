package main

import (
	"fmt"
	"zkfmapf123/src/repository"
)

func main(){

	fromAccount, _ := repository.GetAccount(17)
	toAccount, _ := repository.GetAccount(18)

	err := repository.NewTransfer(fromAccount.Id, toAccount.Id, 500).Transfer(fromAccount, toAccount)

	fmt.Println(err)
}