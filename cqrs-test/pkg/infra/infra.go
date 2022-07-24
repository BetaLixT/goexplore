package infra

import "github.com/google/wire"

var DependencySet = wire.NewSet(
	NewInfra,
)

type Infra struct {

}

func NewInfra() *Infra {
	return &Infra{}
}

func (*Infra) Start() {

}

func (*Infra) Stop() {

}
