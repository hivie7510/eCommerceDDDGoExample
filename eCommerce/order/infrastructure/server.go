package infrastructure

import (
	"eCommerce/order/controllers"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	OrderController controllers.OrderController
}

func ProvideServer(handlers Handlers) func() {
	router := gin.Default()
	router.GET("/test", handlers.OrderController.GetById)
	return func() {
		router.Run()
	}

}
