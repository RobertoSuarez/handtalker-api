package data

import (
	"fmt"

	"github.com/RobertoSuarez/vinculacion_api_graph/db"
)

type User struct {
	ID        uint    `db:"id" json:"id"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password,omitempty"`
	Firstname string  `db:"firstname" json:"firstname"`
	Lastname  string  `db:"lastname" json:"lastname"`
	URLPhoto  *string `db:"url_photo" json:"url_photo"`
	RoleID    uint    `db:"role_id" json:"role_id"`
	Status    bool    `db:"status" json:"status"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (r *CreateUserRequest) Validate() error {
	// validamos el correo electrónico
	if len(r.Email) < 3 {
		return fmt.Errorf("el correo no es correcto")
	}
	return nil
}

func CreateUser(reqUser *CreateUserRequest) (*User, error) {
	photo := "http://foto.com"
	user := &User{
		Email:     reqUser.Email,
		Password:  reqUser.Password,
		Firstname: reqUser.Firstname,
		Lastname:  reqUser.Lastname,
		URLPhoto:  &photo,
		RoleID:    1,
		Status:    true,
	}

	var lastInsertId int

	err := db.DB.QueryRowx(
		`INSERT INTO auth.users (email, password, firstname, lastname, url_photo, role_id, status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		user.Email, user.Password, user.Firstname, user.Lastname, user.URLPhoto, user.RoleID, user.Status).Scan(&lastInsertId)

	if err != nil {
		return user, err
	}

	user.Password = ""
	user.ID = uint(lastInsertId)

	return user, err

}

func GetUsers() ([]User, error) {
	users := []User{}
	err := db.DB.Select(&users, "SELECT id, email, firstname, lastname, url_photo FROM auth.users")
	return users, err
}
