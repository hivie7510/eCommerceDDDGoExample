package services

import (
	"eCommerce/order/entities"
	"eCommerce/order/repositories/interfaces"

	"github.com/google/uuid"
)

type OrderService struct {
	Repository interfaces.OrderRespositoryInterface
}

func (c *OrderService) FindByOrderId(id uuid.UUID) *entities.Order {
	return c.Repository.FindByCustomerId(id)
}

func (c *OrderService) CreateOrder() *entities.Order {
	newId := uuid.New()
	newc := c.Repository.CreateOrder(newId)
	return newc
}
