package crosscutting

import (
	"github.com/google/uuid"
)

type UUID struct{}

func (u *UUID) Generate() string {
	return uuid.New().String()
}

func (u *UUID) IsValid(str string) bool {
	if _, err := uuid.Parse(str); err != nil {
		return false
	}
	return true
}
