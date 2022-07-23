package domain

import (
	"waterforcast/pkg/domain/forcast"

	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	forcast.NewCommands,
	forcast.NewQueries,
)
