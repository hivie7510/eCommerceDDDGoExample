package interfaces

import (
	"eCommerce/order/entities"

	"github.com/google/uuid"
)

type OrderServiceInterface interface {
	FindByOrderId(id uuid.UUID) *entities.Order
	CreateOrder() *entities.Order
}
