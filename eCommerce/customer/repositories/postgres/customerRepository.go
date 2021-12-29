package postgres

import (
	"eCommerce/customer/entities"

	"github.com/google/uuid"
)

type CustomerRespository struct {
}

func (c *CustomerRespository) FindByCustomerId(id uuid.UUID) *entities.Customer {
	println("hit")
	return &entities.Customer{FirstName: "Test", LastName: "User", Id: uuid.New(), InternalId: 1}
}

func (c *CustomerRespository) CreateCustomer(id uuid.UUID, firstName, lastName string) *entities.Customer {
	return nil
}
