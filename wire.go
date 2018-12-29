//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/application"
	"github.com/morio-pg/stubtool/src/domain"
	"github.com/morio-pg/stubtool/src/infrastructure"
	"github.com/morio-pg/stubtool/src/interfaces"
	"github.com/morio-pg/stubtool/src/interfaces/api/server/router"
)

var wireSet = wire.NewSet(application.WireSet, domain.WireSet, infrastructure.WireSet, interfaces.WireSet)

func InitializeRouter() (router router.Router) {
	wire.Build(wireSet)
	return
}
