package infra

import (
	"waterforcast/pkg/domain/forcast"
	"waterforcast/pkg/infra/repos"

	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	repos.NewForcastRepository,
	wire.Bind(new(forcast.IQueryRepository), new(*repos.ForcastRepository)),
	wire.Bind(new(forcast.ICommandRepository), new(*repos.ForcastRepository)),
)
