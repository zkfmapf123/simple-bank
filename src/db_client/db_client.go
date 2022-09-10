package db_client

import (
	"log"
	"zkfmapf123/src/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema =`
	create table accounts (
		id integer,
		owner text,
		balance integer,
		currency text,
		created_at integer,
		updated_at integer,
		timezone text
	)

	create table entries (
		id integer,
		account_id integer,
		amount integer
	)

	create table transfers (
		id integer,
		from_account_id integer,
		to_account_id integer,
		amount integer,
		created_at integer
	)
`

// init query
func Conn() *sqlx.DB{
	config, configErr := utils.LoadConfig(".")
	if configErr != nil {
		log.Fatal(configErr.Error())
	}
	
	db, err := sqlx.Open(config.DBDriver,config.DBSource)

	if err != nil{
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
