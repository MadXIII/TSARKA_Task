package service

import (
	"errors"
	"regexp"
)

var errEmail = errors.New("invalid email") // TODO: too specific, simplify error value, no 'error' text in error

type Search struct{}

func NewSearcher() *Search {
	return &Search{}
}

func (s *Search) CheckEmail(str string) (string, error) {
	reg := regexp.MustCompile(`\s+`)
	cleaned := reg.ReplaceAllString(str, "")

	reg = regexp.MustCompile(`^(Email:)[a-zA-Z0-9][a-zA-Z0-9-_.]+[^\^!#\$%&'\@()*+\/=\?\^\n_{\|}~-]@[a-z]{2,16}\.[a-zA-Z]{2,3}$`)
	if ok := reg.MatchString(cleaned); ok {
		return cleaned, nil
	}
	return "", errEmail
}
