package db

import (
	"github.com/Younkyum/nomadcoin/utils"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

func DB() {
	if db == nil {
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)
		db = dbPointer
	}
	return db
}
