package controllers

import (
	"eCommerce/order/services/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderController struct {
	OrderService interfaces.OrderServiceInterface
}

func (c *OrderController) GetById(g *gin.Context) {
	order := c.OrderService.FindByOrderId(uuid.New())

	g.JSON(200, order)
}
