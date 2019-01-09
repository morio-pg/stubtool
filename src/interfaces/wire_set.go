package interfaces

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/handler"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/middleware"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/router"
)

// WireSet is interfaces ProviderSet
var WireSet = wire.NewSet(
	router.NewRouter,
	middleware.NewAuthenticationMiddleware,
	handler.NewAccountHandler,
	handler.NewAdminHandler,
	handler.NewStubsHandler,
)
