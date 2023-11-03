package main

import (
	"log"

	"github.com/RobertoSuarez/vinculacion_api_graph/models"
	"github.com/gofiber/fiber/v2"
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
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=handtalker_db host=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// db.MustExec(schema)

	// insertamos usuarios con una transacción
	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Roberto", "Suárez", "electrosonix12@gmail.com")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Carlos", "Litardo", "carlos@gmail.com")
	// tx.Commit()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		users := []models.User{}
		err := db.Select(&users, "SELECT id, email, password, firstname, lastname, url_photo, role_id, status FROM auth.users ORDER BY firstname ASC")
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusOK).JSON(users)
	})

	log.Fatal(app.Listen(":3000"))
}
