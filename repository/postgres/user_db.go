package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/madxiii/tsarka_task/model"
)

const (
	usersTable = "users"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) Create(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name) VALUES ($1, $2) RETURNING id`, usersTable)

	row := u.db.QueryRow(query, user.FirstName, user.LastName)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) One(id int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf(`SELECT id, first_name, last_name FROM %s WHERE id=$1`, usersTable)

	err := u.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
	)

	return user, err
}

func (u *User) Update(id int, user model.User) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if user.FirstName != "" {
		args = append(args, user.FirstName)
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", len(args)))
	}

	if user.LastName != "" {
		args = append(args, user.LastName)
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", len(args)))
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d", usersTable, setQuery, id)

	_, err := u.db.Exec(query, args...)

	return err
}

func (u *User) Delete(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, usersTable)

	_, err := u.db.Exec(query, id)

	return err
}
