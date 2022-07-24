package evts

import (
	"context"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/inputs"
)

// OrderBeerOnRoomBooked is a event handler, which handles RoomBooked event and emits OrderBeer command.
type OrderBeerOnRoomBooked struct {
	commandBus *cqrs.CommandBus
}

func NewOrderBeerOnRoomBooked(
	cb *cqrs.CommandBus,
) *OrderBeerOnRoomBooked {
	return &OrderBeerOnRoomBooked{
		commandBus: cb,
	}
}

func (o OrderBeerOnRoomBooked) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "OrderBeerOnRoomBooked"
}

func (OrderBeerOnRoomBooked) NewEvent() interface{} {
	return &inputs.RoomBooked{}
}

func (o OrderBeerOnRoomBooked) Handle(ctx context.Context, e interface{}) error {
	event := e.(*inputs.RoomBooked)

	orderBeerCmd := &inputs.OrderBeer{
		RoomId: event.RoomId,
		Count:  rand.Int63n(10) + 1,
	}

	return o.commandBus.Send(ctx, orderBeerCmd)
}
