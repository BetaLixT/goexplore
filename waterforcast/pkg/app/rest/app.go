package rest

import (
	"waterforcast/pkg/domain"
	"waterforcast/pkg/infra"

	"github.com/google/wire"
)


var dependencySet = wire.NewSet(
  infra.DependencySet,
  domain.DependencySet,
)

func Start() {
  
}
