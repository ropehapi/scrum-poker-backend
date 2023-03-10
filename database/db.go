package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetConn() {
	connString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connString))

	if err != nil {
		panic(err.Error())
	}
}
