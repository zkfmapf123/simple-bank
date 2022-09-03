package db_client

import (
	"testing"
	"zkfmapf123/src/db_client"
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
