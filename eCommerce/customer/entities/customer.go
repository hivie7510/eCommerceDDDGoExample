package entities

import (
	"github.com/google/uuid"
)

type Customer struct {
	Id                  uuid.UUID
	InternalId          int64
	FirstName, LastName string
}

func (c *Customer) IsValid() bool {

	if c.InternalId < 1 {
		println("customer is missing internal id")
		return false
	} else if c.Id == uuid.Nil {
		println(c.Id == uuid.Nil)
		println("customer is missing global id")
		return false
	} else if len(c.FirstName) < 1 || len(c.LastName) < 1 {
		println("customer is missing name")
		return false
	}

	return true
}
