package repositories

import (
	"eCommerce/customer/repositories/interfaces"
	"eCommerce/customer/repositories/postgres"

	"github.com/google/wire"
)

func ProvideCustomerRepository() *postgres.CustomerRespository {
	return &postgres.CustomerRespository{}
}

var RepositoryProviderSet = wire.NewSet(ProvideCustomerRepository, wire.Bind(new(interfaces.CustomerRespositoryInterface), new(*postgres.CustomerRespository)))
