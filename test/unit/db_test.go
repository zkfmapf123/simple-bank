package unit

import (
	"testing"
	"zkfmapf123/src/repository"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T){
	Setup()
	err1 := repository.Transaction("update accounts set balance = ? where id=?", []interface{}{900,17})
	assert.Equal(t, err1, nil)
	defer ShutDown()
}

func TestTrasnfer(t *testing.T){
	Setup()
	fromAccount, err1 := repository.GetAccount("owner1")
	toAccount, err2 := repository.GetAccount("owner2")

	err3 := repository.NewTransfer(fromAccount.Id, toAccount.Id, 500).Transfer(fromAccount, toAccount)

	assert.Equal(t, err1, nil)
	assert.Equal(t, err2, nil)
	assert.Equal(t, err3, nil)
	defer ShutDown()
}

func TestTransferToFail(t *testing.T){
	Setup()
	fromAcount, _ := repository.GetAccount("owner1")
	toAccount, _:= repository.GetAccount("owner2")
	err := repository.NewTransfer(fromAcount.Id, toAccount.Id, 2000).Transfer(fromAcount, toAccount)

	assert.Equal(t, err.Error(), "more money than you currently have")
	defer ShutDown()
}

func TestTransferTx(t *testing.T){
	Setup()
	fromAccount, _ := repository.GetAccount("owner1")
	toAccount, _ := repository.GetAccount("owner2")

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
	defer ShutDown()
 }