package beer

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/pkg/domains/beer/cmds"
	"main.go/pkg/domains/beer/evts"
)

var commandQueries = []cqrs.CommandHandler {
  &cmds.OrderBeerHandler{},
}

var eventHandlers = []cqrs.EventHandler {
  &evts.OrderBeerOnRoomBooked{},
}
