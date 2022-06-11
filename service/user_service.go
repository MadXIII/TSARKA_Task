package service

import (
	"errors"
	"strconv"

	"github.com/madxiii/tsarka_task/model"
	"github.com/madxiii/tsarka_task/repository"
)

var errWrongID = errors.New("invalid id") // TODO: too specific, simplify error value

type User struct {
	repo repository.User
}

func NewUser(repo repository.User) *User {
	return &User{repo: repo}
}

func (u *User) Create(user model.User) (int, error) {
	return u.repo.Create(user)
}

func (u *User) ByID(param string) (model.User, error) {
	var user model.User

	id, err := strconv.Atoi(param)
	if err != nil {
		return user, errWrongID
	}

	user, err = u.repo.One(id)
	return user, err
}

func (u *User) Update(param string, user model.User) error {
	id, err := strconv.Atoi(param)
	if err != nil {
		return errWrongID
	}

	err = u.repo.Update(id, user)
	return err
}

func (u *User) Delete(param string) error {
	id, err := strconv.Atoi(param)
	if err != nil {
		return errWrongID
	}

	err = u.repo.Delete(id)
	return err
}
