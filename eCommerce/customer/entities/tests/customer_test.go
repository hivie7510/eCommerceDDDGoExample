package tests

import (
	"eCommerce/customer/entities"
	"testing"

	"github.com/google/uuid"
)

func TestWithoutGlobalId(t *testing.T) {
	uut := entities.Customer{FirstName: "Heath", LastName: "Ivie", InternalId: 1}
	result := uut.IsValid()
	if result == true {
		t.Errorf("Customer was passed validation without a global id")
	}
}

func TestWithoutInternalId(t *testing.T) {
	uut := entities.Customer{Id: uuid.New(), FirstName: "Heath", LastName: "Ivie"}
	result := uut.IsValid()
	if result == true {
		t.Errorf("Customer was passed validation without a internal id")
	}
}

func TestWithoutName(t *testing.T) {
	uut := entities.Customer{Id: uuid.New(), FirstName: "Heath", InternalId: 1}
	result := uut.IsValid()
	if result == true {
		t.Errorf("Customer was passed validation without a name")
	}

	uut = entities.Customer{Id: uuid.New(), LastName: "Ivie", InternalId: 1}
	result = uut.IsValid()
	if result == true {
		t.Errorf("Customer was passed validation without a name")
	}
}
