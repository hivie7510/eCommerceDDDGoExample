package services

import (
	"eCommerce/order/services/interfaces"

	"github.com/google/wire"
)

var ProvideServices = wire.NewSet(

	wire.Struct(new(OrderService), "*"),
)

var ServicesProviderSet = wire.NewSet(ProvideServices, wire.Bind(new(interfaces.OrderServiceInterface), new(*OrderService)))
