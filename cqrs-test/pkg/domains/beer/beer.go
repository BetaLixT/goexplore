package beer

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/pkg/domains/beer/cmds"
	"main.go/pkg/domains/beer/evts"
)

type BeerDomain struct {
}

func NewBeerDomain() *BeerDomain {
  return &BeerDomain{}
}

func (b *BeerDomain) RegisterCommandQueryHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.CommandHandler {
  return []cqrs.CommandHandler {
    cmds.NewOrderBeerHandler(eb),
  }
}

func (b *BeerDomain) RegisterEventHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.EventHandler {
  return []cqrs.EventHandler {
    evts.NewOrderBeerOnRoomBooked(cb),
  }
}
