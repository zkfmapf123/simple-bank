package main

import (
	"fmt"
	"zkfmapf123/src/db_client"
)

func main(){
	db := db_client.Conn()
	fmt.Println(db)

	accounts := []string{}
	err := db.Select(&accounts, "select owner from accounts order by created_at asc")
	if err != nil{
		panic(err)
	}

	fmt.Println(accounts)


	
	

}