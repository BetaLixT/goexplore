package cmds

import (
	"context"
	"waterforcast/pkg/domain/forcast"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type AddForcastHandler struct {
	repo forcast.ICommandRepository
	eventBus *cqrs.EventBus
}

func NewAddForcastHandler(
	repo forcast.ICommandRepository,
) *AddForcastHandler  {
	return &AddForcastHandler{
		repo: repo,
	}
}

func (cmd AddForcastHandler) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "AddForcast"
}

func (AddForcastHandler) NewCommand() interface{} {
	return &AddForcast{}
}

func (cmd *AddForcastHandler) Handle(ctx context.Context, e interface{}) error {

	event := e.(*AddForcast)

	// When we are using Pub/Sub which doesn't provide exactly-once delivery semantics, we need to deduplicate messages.
	// GoChannel Pub/Sub provides exactly-once delivery,
	// but let's make this example ready for other Pub/Sub implementations.
	cmd.repo.Create(
		ctx,
		event.Id,
		event.Title,
	)
	return nil
}

type AddForcast struct {
	Id    int
	Title string
}
