package tests

import (
	"eCommerce/order/entities"
	"eCommerce/order/repositories/interfaces"
	"testing"

	"github.com/google/uuid"
)

type MockGeneratedWithGlobalIdRepository struct {
	interfaces.OrderRespositoryInterface
}

var (
	withGlobalId = uuid.New()
)

func (m MockGeneratedWithGlobalIdRepository) CreateOrder() *entities.Order {
	return &entities.Order{Id: withGlobalId}
}

func TestNewGlobalIdIsGenerated(t *testing.T) {
	// d := MockGeneratedWithGlobalIdRepository{}

}
