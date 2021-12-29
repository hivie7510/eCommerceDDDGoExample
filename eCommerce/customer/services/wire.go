package services

import (
	"eCommerce/customer/services/interfaces"

	"github.com/google/wire"
)

var ProvideServices = wire.NewSet(

	wire.Struct(new(CustomerService), "*"),
)

var ServicesProviderSet = wire.NewSet(ProvideServices, wire.Bind(new(interfaces.CustomerServiceInterface), new(*CustomerService)))
