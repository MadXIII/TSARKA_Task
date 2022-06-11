package repository

import (
	"github.com/madxiii/tsarka_task/model"
)

type Count interface {
	Value() (string, error)
	Add(num int) error
	Sub(num int) error
}

type User interface {
	Create(userData model.User) (int, error)
	One(id int) (model.User, error)
	Update(id int, userData model.User) error
	Delete(id int) error
}

// FIXME: return with errors
type Hash interface {
	StoreKey(key string)
	StoreValByKey(key string, val int)
	GetValueByKey(key string) (int, error)
}

type Repository struct {
	Count
	User
	Hash
}

func New(user User, count Count, hash Hash) *Repository {
	return &Repository{Count: count, User: user, Hash: hash}
}
