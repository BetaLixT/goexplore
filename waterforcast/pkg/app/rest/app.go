package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	stdHttp "net/http"
	"time"
	"waterforcast/pkg/common"
	"waterforcast/pkg/domain"
	"waterforcast/pkg/domain/forcast/cmds"
	"waterforcast/pkg/infra"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/google/wire"
)

var dependencySet = wire.NewSet(
	infra.DependencySet,
	domain.DependencySet,
	NewRestApp,
)

func Start() {
	a, err := InitializeRestApp()
	if err != nil {
		panic(err)
	}
	a.startService()
}

type RestApp struct {
	addForcastHandler *cmds.AddForcastHandler
}

func NewRestApp(
	addForcastHandler *cmds.AddForcastHandler,
) *RestApp {
	return &RestApp{
		addForcastHandler: addForcastHandler,
	}
}

var (
	// For this example, we're using just a simple logger implementation,
	// You probably want to ship your own implementation of `watermill.LoggerAdapter`.
	logger = watermill.NewStdLogger(false, false)
)

func (app *RestApp) startService() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)

	httpSubscriber, err := http.NewSubscriber(
		":8080",
		http.SubscriberConfig{
			UnmarshalMessageFunc: func(
				topic string,
				request *stdHttp.Request,
			) (*message.Message, error) {
				b, err := ioutil.ReadAll(request.Body)
				if err != nil {
					return nil, common.NewBodyReadingError(err)
				}

				return message.NewMessage(watermill.NewUUID(), b), nil
			},
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	cqrsMarshaler := cqrs.ProtobufMarshaler{}
	// just for debug, we are printing all messages received on `incoming_messages_topic`
	_, err = cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return commandName
		},
		CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
			return []cqrs.CommandHandler{
				app.addForcastHandler,
			}
		},
		CommandsPublisher: pubSub,
		CommandsSubscriberConstructor: func(
			handlerName string,
		) (message.Subscriber, error) {
			// we can reuse subscriber, because all commands have separated topics
			return httpSubscriber, nil
		},
		GenerateEventsTopic: func(eventName string) string {
			// because we are using PubSub RabbitMQ config, we can use one topic for all events
			return "notifications"

			// we can also use topic per event type
			// return eventName
		},
		EventHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
			return []cqrs.EventHandler{}
		},
		EventsPublisher: pubSub,
		EventsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {	

			return pubSub, nil
		},
		Router:                router,
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})
	if err != nil {
		panic(err)
	}

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}

type structHandler struct {
	// we can add some dependencies here
}

func (s structHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	log.Println("structHandler received message", msg.UUID)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return message.Messages{msg}, nil
}
