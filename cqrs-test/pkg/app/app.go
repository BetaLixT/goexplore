package app

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/wire"
	"main.go/inputs"
	"main.go/pkg/domains"
	"main.go/pkg/infra"
)

func Start() {
  app, err := InitializeApp()
  if err != nil {
    panic(err)
  }
  app.startService()
}

var dependencySet = wire.NewSet(
	infra.DependencySet,
	domains.DependencySet,
	newApp,
)

type app struct {
  dom *domains.Domain
}

func newApp(
  dom *domains.Domain,
) *app {
  return &app{
    dom: dom,
  }
}

var amqpAddress = "amqp://guest:guest@127.0.0.1:5672/"

func (app *app) startService() {
  logger := watermill.NewStdLogger(false, false)
	cqrsMarshaler := cqrs.ProtobufMarshaler{}

	// You can use any Pub/Sub implementation from here: https://watermill.io/docs/pub-sub-implementations/
	// Detailed RabbitMQ implementation: https://watermill.io/docs/pub-sub-implementations/#rabbitmq-amqp
	// Commands will be send to queue, because they need to be consumed once.
	commandsAMQPConfig := amqp.NewDurableQueueConfig(amqpAddress)
	commandsPublisher, err := amqp.NewPublisher(commandsAMQPConfig, logger)
	if err != nil {
		panic(err)
	}
	commandsSubscriber, err := amqp.NewSubscriber(commandsAMQPConfig, logger)
	if err != nil {
		panic(err)
	}

	// Events will be published to PubSub configured Rabbit, because they may be consumed by multiple consumers.
	// (in that case BookingsFinancialReport and OrderBeerOnRoomBooked).
	eventsPublisher, err := amqp.NewPublisher(amqp.NewDurablePubSubConfig(amqpAddress, nil), logger)
	if err != nil {
		panic(err)
	}

	// CQRS is built on messages router. Detailed documentation: https://watermill.io/docs/messages-router/
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	// Simple middleware which will recover panics from event or command handlers.
	// More about router middlewares you can find in the documentation:
	// https://watermill.io/docs/messages-router/#middleware
	//
	// List of available middlewares you can find in message/router/middleware.
	router.AddMiddleware(middleware.Recoverer)


	// cqrs.Facade is facade for Command and Event buses and processors.
	// You can use facade, or create buses and processors manually (you can inspire with cqrs.NewFacade)
	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return commandName
		},
		CommandHandlers: func(
		  cb *cqrs.CommandBus,
		  eb *cqrs.EventBus,
		) []cqrs.CommandHandler {
			return app.dom.RegisterCommandQueryHandlers(cb, eb)
		},
		CommandsPublisher: commandsPublisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			// we can reuse subscriber, because all commands have separated topics
			return commandsSubscriber, nil
		},
		GenerateEventsTopic: func(eventName string) string {
			// because we are using PubSub RabbitMQ config, we can use one topic for all events
			return "events"

			// we can also use topic per event type
			// return eventName
		},
		EventHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
			return app.dom.RegisterEventHandlers(cb, eb)
		},
		EventsPublisher: eventsPublisher,
		EventsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			config := amqp.NewDurablePubSubConfig(
				amqpAddress,
				amqp.GenerateQueueNameTopicNameWithSuffix(handlerName),
			)

			return amqp.NewSubscriber(config, logger)
		},
		Router:                router,
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})
	if err != nil {
		panic(err)
	}


	// publish BookRoom commands every second to simulate incoming traffic
	go publishCommands(cqrsFacade.CommandBus())

	// processors are based on router, so they will work when router will start
	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func publishCommands(commandBus *cqrs.CommandBus) func() {
	i := 0
	for {
		i++

		startDate, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			panic(err)
		}

		endDate, err := ptypes.TimestampProto(time.Now().Add(time.Hour * 24 * 3))
		if err != nil {
			panic(err)
		}

		bookRoomCmd := &inputs.BookRoom{
			RoomId:    fmt.Sprintf("%d", i),
			GuestName: "John",
			StartDate: startDate,
			EndDate:   endDate,
		}
		if err := commandBus.Send(context.Background(), bookRoomCmd); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
