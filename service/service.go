package service

import (
	"github.com/madxiii/tsarka_task/repository"
)

type Service struct {
	Finder *Finder
	Search *Search
	Count  *Counter
	User   *User
	Hash   *Hasher
}

func New(repo repository.Repository) *Service {
	return &Service{
		Finder: NewFinder(),
		Search: NewSearcher(),
		User:   NewUser(repo.User),
		Count:  NewCounter(repo.Count),
		Hash:   NewHasher(repo.Hash),
	}
}
