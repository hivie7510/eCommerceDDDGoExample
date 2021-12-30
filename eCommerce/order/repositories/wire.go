package repositories

import (
	"eCommerce/order/repositories/interfaces"
	"eCommerce/order/repositories/postgres"

	"github.com/google/wire"
)

func ProvideOrderRepository() *postgres.OrderRespository {
	return &postgres.OrderRespository{}
}

var RepositoryProviderSet = wire.NewSet(ProvideOrderRepository, wire.Bind(new(interfaces.OrderRespositoryInterface), new(*postgres.OrderRespository)))
