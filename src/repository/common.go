package repository

import (
	"zkfmapf123/src/base"
	"zkfmapf123/src/db_client"
)

type storeable interface {
	base.AccountModels | base.EntrieModels | base.TranserModels
}

func FindOne[T storeable](stores []T,query string, args interface{}) ([]T, error){
	var store T
	db := db_client.Conn()
	rows , err :=db.Queryx(query,args)

	if err != nil {
		return nil,err
	}

	for	rows.Next(){
		err := rows.StructScan(&store)
		if err != nil{
			return nil, err
		}

		stores = append(stores, store)
	}

	defer db.Close()
	return stores, nil
}