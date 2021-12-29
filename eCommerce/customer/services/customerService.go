package services

import (
	"eCommerce/customer/entities"
	"eCommerce/customer/repositories/interfaces"
	"eCommerce/customer/values"

	"github.com/google/uuid"
)

type CustomerService struct {
	Repository interfaces.CustomerRespositoryInterface
}

func (c *CustomerService) FindByCustomerId(id uuid.UUID) *entities.Customer {
	return c.Repository.FindByCustomerId(id)
}

func (c *CustomerService) FindAddressesByCustomerId(id uuid.UUID) *[]values.CustomerAddress {
	return nil
}

func (c *CustomerService) CreateCustomer(firstName, lastName string) *entities.Customer {
	newId := uuid.New()
	newc := c.Repository.CreateCustomer(newId, firstName, lastName)
	return newc
}
