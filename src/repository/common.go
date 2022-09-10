package repository

import (
	"log"
	"zkfmapf123/src/base"
	"zkfmapf123/src/db_client"
)

type storeable interface {
	base.AccountModels | base.EntrieModels | base.TranserModels
}

/*
* @desc Must Todo Update, Delete
*/
func Transaction(query string, args []interface{}) error {

	db := db_client.Conn()

	// transaction begin
	tx ,err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil{
		tx.Rollback()
		return err
	}

	defer db.Close()
	return nil
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

func FindList[T storeable](stores []T, query string) ([]T, error) {
	var store T
	db := db_client.Conn()
	rows, err := db.Queryx(query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next(){
		err := rows.StructScan(&store)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	defer db.Close()
	return stores, nil
}