package app

import (
	"github.com/google/wire"
	"main.go/pkg/domains"
	"main.go/pkg/infra"
)

func Start() {

}

var dependencySet = wire.NewSet(
	infra.DependencySet,
	domains.DependencySet,
	NewApp,
)

type app struct {

}

func NewApp() *app {
  return &app{}
}

func (app *app) startService() {

}
