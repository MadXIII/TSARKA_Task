package service

import (
	"strconv"

	"github.com/madxiii/tsarka_task/repository"
)

type Counter struct {
	repo repository.Count
}

func NewCounter(repo repository.Count) *Counter {
	return &Counter{repo: repo}
}

func (s *Counter) CounterGet() (string, error) {
	res, err := s.repo.Value()
	return res, err
}

func (s *Counter) CounterAdd(param string) error {
	num, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	err = s.repo.Add(num)

	return err
}

func (s *Counter) CounterSub(param string) error {
	num, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	err = s.repo.Sub(num)

	return err
}
