package db_client

import (
	"testing"
	"zkfmapf123/src/db_client"
	"zkfmapf123/src/repository"
	"zkfmapf123/src/utils"

	"github.com/stretchr/testify/assert"
)

/*
	go test ./test
*/

func TestAccounts(t *testing.T){

	db := db_client.Conn()
	_, err := db.NamedExec("insert into accounts (owner,balance,currency,created_at, updated_at,timezone) values (:owner,:balance,:currency,:createdAt,:updatedAt,:timezone)",
	map[string]interface{}{
		"owner": "owner",
		"balance":1000,
		"currency" : "KR",
		"createdAt": utils.Now(),
		"updatedAt" : utils.Now(),
		"timezone" : "Asia/Seoul",
	})

	assert.Equal(t, err,nil)
}

func TestTransaction(t *testing.T){
	err1 := repository.Transaction("update accounts set balance = ? where id=?", []interface{}{900,17})
	assert.Equal(t, err1, nil)
}

func TestTrasnfer(t *testing.T){

	fromAccount, err1 := repository.GetAccount(17)
	toAccount, err2 := repository.GetAccount(18)

	err3 := repository.NewTransfer(fromAccount.Id, toAccount.Id, 500).Transfer(fromAccount, toAccount)

	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
	assert.Equal(t, err3, nil)
}

func TestTransferToFail(t *testing.T){

	fromAcount, _ := repository.GetAccount(17)
	toAccount, _:= repository.GetAccount(18)
	err := repository.NewTransfer(fromAcount.Id, toAccount.Id, 2000).Transfer(fromAcount, toAccount)

	assert.Equal(t, err.Error(), "more money than you currently have")
}

func TestTransferTx(t *testing.T){
	fromAccount, _ := repository.GetAccount(17)
	toAccount, _ := repository.GetAccount(18)

	// receiver
	errs := make(chan error)
	n := 5
	amount:= int64(300)

	for i := 0; i< n; i++ {
		go func() {
			err := repository.NewTransfer(fromAccount.Id, toAccount.Id, int(amount)).Transfer(fromAccount, toAccount)
			errs <- err // inject
		}()
	}

	// check result
	for i := 0; i<n; i++ {
		err := <-errs
		assert.Equal(t, err, nil)
	}
 }