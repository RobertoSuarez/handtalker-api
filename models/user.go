package models

type User struct {
	ID        uint    `db:"id"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	Firstname string  `db:"firstname"`
	Lastname  string  `db:"lastname"`
	URLPhoto  *string `db:"url_photo"`
	RoleID    uint    `db:"role_id"`
	Status    bool    `db:"status"`
}
