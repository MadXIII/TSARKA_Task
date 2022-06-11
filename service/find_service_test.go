package service

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCheckStr(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		waitRes string
	}{
		{
			name:    "SuccessCheckStr",
			input:   "a",
			waitRes: "a",
		},
		{
			name:    "SuccessFindSubstring",
			input:   "dasdf",
			waitRes: "asdf",
		},
	}

	for _, test := range tests {
		s := NewFinder()
		t.Run(test.name, func(t *testing.T) {
			checkRes := s.CheckStr(test.input)
			if checkRes == test.waitRes {
				assert.Equal(t, checkRes, test.waitRes)
			} else {
				findRes := FindSubstring(test.input)
				assert.Equal(t, findRes, test.waitRes)
			}
		})
	}
}
