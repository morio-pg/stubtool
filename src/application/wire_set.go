package application

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/application/usecase"
)

// WireSet is domain ProviderSet
var WireSet = wire.NewSet(
	usecase.NewAuthenticationUsecase,
	usecase.NewAdminUsecase,
	usecase.NewAccountUsecase,
	usecase.NewStubUsecase,
)
