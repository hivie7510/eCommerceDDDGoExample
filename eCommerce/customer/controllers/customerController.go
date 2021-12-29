package controllers

import (
	"eCommerce/customer/services/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerController struct {
	CustomerService interfaces.CustomerServiceInterface
}

func (c *CustomerController) GetById(g *gin.Context) {
	customer := c.CustomerService.FindByCustomerId(uuid.New())

	g.JSON(200, customer)
}
