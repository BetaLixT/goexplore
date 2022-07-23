package domain

import (
	"waterforcast/pkg/domain/forcast"
	"waterforcast/pkg/domain/forcast/cmds"

	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	cmds.NewAddForcastHandler,
	forcast.NewQueries,
)
