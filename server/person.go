package main

import (
	"errors"

	"github.com/rs/xid"
)

var errInvalidNameOrSurname = errors.New("person: invalid name or surname given")

// Person represents a human, having a unique ID, Name, Surname and Age.
type Person struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

// NewPerson instantiates a new Person with the given name, surname and age parameters
func NewPerson(name, surname string, age uint) (*Person, error) {
	if name == "" || surname == "" {
		return nil, errInvalidNameOrSurname
	}

	return &Person{
		ID:      xid.New().String(),
		Name:    name,
		Surname: surname,
		Age:     age,
	}, nil
}
