package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
	CREATE TABLE IF NOT EXISTS person (
		first_name text,
		last_name text,
		email text
	);
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(schema)

	// insertamos usuarios con una transacción
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Roberto", "Suárez", "electrosonix12@gmail.com")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Carlos", "Litardo", "carlos@gmail.com")
	tx.Commit()

	people := []Person{}

	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")

	fmt.Printf("%#v \n", people)

}
