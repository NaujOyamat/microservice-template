package models

import "time"

type Person struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Birthdate time.Time `json:"birthdate"`
	Merried   bool      `json:"merried"`
}
