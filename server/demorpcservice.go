package main

import (
	"errors"
	"log"
	"net/http"
)

var errPersonNotFound = errors.New("demorpcservice: person not found with the given ID")

// DemoRPCService provides some RPC methods useful to create and get a Person
type DemoRPCService struct {
	// This is for demo purposes, the RPC service should have a data provider
	// such a database or whatsoever. DON'T DO THIS IN PRODUCTION!
	personList map[string]Person
}

// NewDemoRPCService instantiates a new DemoRPCService
func NewDemoRPCService() *DemoRPCService {
	return &DemoRPCService{
		personList: make(map[string]Person),
	}
}

// CreatePerson instantiates a new Person with the given args, if they're correct, stores the Person
// in the local list and returns it
func (s *DemoRPCService) CreatePerson(r *http.Request, args *CreatePersonArgs, reply *Person) error {
	log.Printf("got CreatePerson request from %s\n", r.RemoteAddr)

	newPerson, err := NewPerson(args.Name, args.Surname, args.Age)
	if err != nil {
		return err
	}

	s.personList[newPerson.ID] = *newPerson

	*reply = *newPerson
	return nil
}

// GetPerson returns the Person matching the given args.ID, if present, or an error otherwise
func (s *DemoRPCService) GetPerson(r *http.Request, args *GetPersonArgs, reply *Person) error {
	log.Printf("got GetPerson request from %s\n", r.RemoteAddr)

	person, personExists := s.personList[args.ID]
	if !personExists {
		return errPersonNotFound
	}

	*reply = person
	return nil
}

// CreatePersonArgs contains all the parameters used in the CreatePerson method
type CreatePersonArgs struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     uint   `json:"age"`
}

// GetPersonArgs contains all the parameters used in the GetPerson method
type GetPersonArgs struct {
	ID string `json:"id"`
}
