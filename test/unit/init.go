package unit

import (
	"zkfmapf123/src/db_client"
	"zkfmapf123/src/utils"
)

func Setup(){
	db := db_client.Conn()
	db.NamedExec("insert into accounts (owner,balance,currency,created_at, updated_at,timezone) values (:owner,:balance,:currency,:createdAt,:updatedAt,:timezone)",
	map[string]interface{}{
		"owner": "owner1",
		"balance":1000,
		"currency" : "KR",
		"createdAt": utils.Now(),
		"updatedAt" : utils.Now(),
		"timezone" : "Asia/Seoul",
	})

	db.NamedExec("insert into accounts (owner,balance,currency,created_at, updated_at,timezone) values (:owner,:balance,:currency,:createdAt,:updatedAt,:timezone)",
	map[string]interface{}{
		"owner": "owner2",
		"balance":1000,
		"currency" : "KR",
		"createdAt": utils.Now(),
		"updatedAt" : utils.Now(),
		"timezone" : "Asia/Seoul",
	})

	defer db.Close()
}

func ShutDown(){
	db := db_client.Conn()

	db.QueryRow("delete from accounts where owner=?", "owner1")
	db.QueryRow("delete from accounts where owner=?", "owner2")
	defer db.Close()
}
