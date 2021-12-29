package infrastructure

import (
	"eCommerce/customer/controllers"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	CustomerController controllers.CustomerController
}

func ProvideServer(handlers Handlers) func() {
	router := gin.Default()
	router.GET("/test", handlers.CustomerController.GetById)
	return func() {
		router.Run()
	}

}
