package repository

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"zkfmapf123/src/api/helpers"
	"zkfmapf123/src/base"
	"zkfmapf123/src/db_client"
	"zkfmapf123/src/utils"

	"github.com/fatih/structs"
)

type accounts struct {
	Owner string `json:"owner"`
	Balance int `json:"balance"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	Currency string `json:"currency"`
	Timezone string `json:"timezone"`
}

type FindParmas interface {
	int | string
}

func NewAccount(owner string, balance int, currency string, tz string) *accounts {
	now := utils.Now()
	return &accounts{Owner: owner, Balance: balance, Currency: currency, Timezone: tz, CreatedAt: now, UpdatedAt: now}
}

func (ac *accounts) CreateAccount () error{
	db := db_client.Conn()
	_, err := db.NamedExec("insert into accounts (owner, balance,currency, created_at, updated_at, timezone) values (:Owner, :Balance,:Currency, :CreatedAt, :UpdatedAt, :Timezone)",structs.Map(ac))

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()
	return nil
}

func (ac *accounts) CreateAccountUseQuery (query string) error {
	db := db_client.Conn()
	_ ,err := db.NamedExec(query, structs.Map(ac))
	defer db.Close()
	return err
}

/*
* @todo use _.first
*/
func GetAccount[T FindParmas](param T) (*base.AccountModels, error) {
	accounts := []base.AccountModels{}

	var rows []base.AccountModels
	var err error

	if reflect.TypeOf(param).Name() == "string" {
		rows, err = FindOne(accounts, "select * from accounts where owner = ?", param)
	}else{
		rows, err = FindOne(accounts, "select * from accounts where id = ?", param)
	}

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New(helpers.NOT_EXISTS_OWNER)
	}

	return &rows[0], nil
} 

func GetAccountList(limit string, offset string) (*[]base.AccountModels, error){
	accounts := []base.AccountModels{}

	_limit, _ := strconv.Atoi(limit)
	_offset, _ := strconv.Atoi(offset)

	query := "select * from accounts"
	query += fmt.Sprintf(" limit %d offset %d", _limit, _offset)
	rows, err := FindList(accounts,query)

	if err != nil{
		return nil, err
	}
	
	return &rows, nil
}
