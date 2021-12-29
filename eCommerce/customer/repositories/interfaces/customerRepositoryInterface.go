package interfaces

import (
	"eCommerce/customer/entities"

	"github.com/google/uuid"
)

type CustomerRespositoryInterface interface {
	FindByCustomerId(id uuid.UUID) *entities.Customer
	CreateCustomer(id uuid.UUID, firstName, lastName string) *entities.Customer
}
