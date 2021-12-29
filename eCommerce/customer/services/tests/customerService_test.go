package tests

import (
	"eCommerce/customer/entities"
	"eCommerce/customer/repositories/interfaces"
	"eCommerce/customer/services"
	"testing"

	"github.com/google/uuid"
)

type MockGeneratedWithGlobalIdRepository struct {
	interfaces.CustomerRespositoryInterface
}

var (
	withGlobalId = uuid.New()
)

func (m MockGeneratedWithGlobalIdRepository) CreateCustomer(id uuid.UUID, firstName, lastName string) *entities.Customer {
	return &entities.Customer{Id: withGlobalId, FirstName: "Heath", LastName: "Ivie"}
}

func TestNewGlobalIdIsGenerated(t *testing.T) {
	d := MockGeneratedWithGlobalIdRepository{}

	uut := services.CustomerService{Repository: d}
	customer := uut.CreateCustomer("Heath", "Ivie")
	if customer == nil || customer.Id == uuid.Nil {
		t.Errorf("Customer was not created")
	}
}

var (
	withInternalId = int64(1)
)

type MockGeneratedWithInteralIdRepository struct {
	interfaces.CustomerRespositoryInterface
}

func (m MockGeneratedWithInteralIdRepository) CreateCustomer(id uuid.UUID, firstName, lastName string) *entities.Customer {
	return &entities.Customer{Id: withGlobalId, InternalId: withInternalId, FirstName: "Heath", LastName: "Ivie"}
}

func TestNewInternalIsGenerated(t *testing.T) {
	d := MockGeneratedWithInteralIdRepository{}

	uut := services.CustomerService{Repository: d}
	customer := uut.CreateCustomer("Heath", "Ivie")
	if customer == nil || customer.InternalId < 1 {
		t.Errorf("Customer was not created or was missing internal id")
	}

}
