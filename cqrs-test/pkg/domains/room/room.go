package room

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/pkg/domains/room/cmds"
)

type RoomDomain struct {

}

func NewFinancialDomain() *RoomDomain {
  return &RoomDomain{}
}

func (b *RoomDomain) RegisterCommandQueryHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.CommandHandler {
  return []cqrs.CommandHandler {
    cmds.NewBookRoomHandler(eb),
  }
}

func (b *RoomDomain) RegisterEventHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.EventHandler {
  return []cqrs.EventHandler {
  }
}
