package domain

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/domain/service"
)

// WireSet is domain ProviderSet
var WireSet = wire.NewSet(service.NewStubService)
