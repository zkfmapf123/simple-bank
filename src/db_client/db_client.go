package db_client

import (
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

type Accounts struct {
	id int 
	owner string
	balance int
	currency string
	createdAt string `db:"created_at"`
	updatedAt string `db:"updated_at"`
	timezone string `db:"timezone"`
}

type Entries struct{
	id int
	accountId string `db:"account_id"`
	amount int
}

type Transers struct {
	id int
	fromAccountId int `db:"from_account_id"`
	toAccountId int `db:"to_account_id"`
	amount int
	createdAt int `db:"created_at"`
}


// init query
func Conn() *sqlx.DB{
	db, err := sqlx.Open("mysql","root:1234@tcp(localhost:3306)/simple_bank?parseTime=true")

	if err != nil{
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	
	return db
}

