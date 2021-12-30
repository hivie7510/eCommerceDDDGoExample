package interfaces

import (
	"eCommerce/order/entities"

	"github.com/google/uuid"
)

type OrderRespositoryInterface interface {
	FindByCustomerId(id uuid.UUID) *entities.Order
	CreateOrder(id uuid.UUID) *entities.Order
}
