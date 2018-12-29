package infrastructure

import (
	"github.com/google/wire"
	"github.com/morio-pg/stubtool/src/infrastructure/firebase"
)

// WireSet is infrastructure ProviderSet
var WireSet = wire.NewSet(firebase.NewAccountRepository, firebase.NewStubRepository)
