//go:build wireinject
// +build wireinject

package main

import (
	"eCommerce/order/controllers"
	"eCommerce/order/infrastructure"
	"eCommerce/order/repositories"
	"eCommerce/order/services"

	"github.com/google/wire"
)

func Init() func() {
	panic(wire.Build(repositories.RepositoryProviderSet, infrastructure.InfrastructureProviderSet, services.ServicesProviderSet, controllers.ControllersProviderSet))

}
