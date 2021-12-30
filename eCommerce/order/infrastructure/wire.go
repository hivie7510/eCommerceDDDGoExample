package infrastructure

import "github.com/google/wire"

var Providehandlers = wire.NewSet(

	wire.Struct(new(Handlers), "*"),
)

var InfrastructureProviderSet = wire.NewSet(
	ProvideServer,
	Providehandlers,
)
