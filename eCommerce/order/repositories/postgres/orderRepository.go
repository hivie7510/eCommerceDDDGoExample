package postgres

import (
	"eCommerce/order/entities"

	"github.com/google/uuid"
)

type OrderRespository struct {
}

func (c *OrderRespository) FindByCustomerId(id uuid.UUID) *entities.Order {
	println("hit")
	return &entities.Order{Id: uuid.New(), InternalId: 1}
}

func (c *OrderRespository) CreateOrder(id uuid.UUID) *entities.Order {
	return nil
}
