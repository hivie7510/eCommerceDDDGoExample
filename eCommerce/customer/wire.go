//go:build wireinject
// +build wireinject

package main

import (
	"eCommerce/customer/controllers"
	"eCommerce/customer/infrastructure"
	"eCommerce/customer/repositories"
	"eCommerce/customer/services"

	"github.com/google/wire"
)

func Init() func() {
	panic(wire.Build(repositories.RepositoryProviderSet, infrastructure.InfrastructureProviderSet, services.ServicesProviderSet, controllers.ControllersProviderSet))

}
