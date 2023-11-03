package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func CreateDatabase() (*sqlx.DB, error) {
	uri := "user=postgres password=postgres dbname=handtalker_db host=db sslmode=disable"
	return sqlx.Connect("postgres", uri)
}

func Init() error {
	db, err := CreateDatabase()
	DB = db
	return err
}
