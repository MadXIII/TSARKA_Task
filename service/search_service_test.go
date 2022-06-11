package service

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCheckEmail(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		waitResult string
		waitError  error
	}{
		{
			name:       "Success",
			input:      "Email:     test@test.ru",
			waitResult: "Email:test@test.ru",
			waitError:  nil,
		},
		{
			name:       "Wait Error Invalid email",
			input:      "Email:     test@test",
			waitResult: "",
			waitError:  errEmail,
		},
	}

	s := NewSearcher()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := s.CheckEmail(test.input)
			if err != nil {
				assert.Equal(t, res, test.waitResult)
				assert.Equal(t, err, test.waitError)
			} else {
				assert.Equal(t, res, test.waitResult)
			}
		})
	}
}
