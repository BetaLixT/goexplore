package cmds

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/inputs"
)

// OrderBeerHandler is a command handler, which handles OrderBeer command and emits BeerOrdered.
// BeerOrdered is not handled by any event handler, but we may use persistent Pub/Sub to handle it in the future.
type OrderBeerHandler struct {
	eventBus *cqrs.EventBus
}

func NewOrderBeerHandler(
	eb *cqrs.EventBus,
) *OrderBeerHandler {
	return &OrderBeerHandler{
		eventBus: eb,
	}
}

func (o OrderBeerHandler) HandlerName() string {
	return "OrderBeerHandler"
}

func (o OrderBeerHandler) NewCommand() interface{} {
	return &inputs.OrderBeer{}
}

func (o OrderBeerHandler) Handle(ctx context.Context, c interface{}) error {
	cmd := c.(*inputs.OrderBeer)

	if rand.Int63n(10) == 0 {
		// sometimes there is no beer left, command will be retried
		return errors.Errorf("no beer left for room %s, please try later", cmd.RoomId)
	}

	if err := o.eventBus.Publish(ctx, &inputs.BeerOrdered{
		RoomId: cmd.RoomId,
		Count:  cmd.Count,
	}); err != nil {
		return err
	}

	log.Printf("%d beers ordered to room %s", cmd.Count, cmd.RoomId)
	return nil
}
