package interfaces

import (
	"eCommerce/customer/entities"
	"eCommerce/customer/values"

	"github.com/google/uuid"
)

type CustomerServiceInterface interface {
	FindByCustomerId(id uuid.UUID) *entities.Customer
	FindAddressesByCustomerId(id uuid.UUID) *[]values.CustomerAddress
	CreateCustomer(firstName, lastName string) *entities.Customer
}
