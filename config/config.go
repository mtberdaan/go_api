package config

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	host        = "localhost"
	port        = 8001
	user        = "postgres"
	password    = "passwd"
	dbname      = "postgres"
	IdentityKey = "id"
	Key         = "my_secret_key_8f6e2p"
)

var DB *gorm.DB

func Init() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
