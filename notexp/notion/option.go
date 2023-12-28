package notion

import (
	"time"

	"github.com/Soreing/gent"
	"github.com/Soreing/retrier"
)

type ClientOption func(*Configuration) *Configuration

type Configuration struct {
	retrier       gent.Retrier
	client        *gent.Client
	gentOptions   []gent.Option
	apiKey        *string
	apiVersion    *string
	notionVersion *string
}

func Build(config *Configuration) (*Client, error) {
	// TODO: research more auth stuff
	var apiKey, apiVer, notionVer string
	var client *gent.Client
	if config.apiKey == nil {
		return nil, ErrNoAPIKey
	}
	apiKey = *config.apiKey

	if config.apiVersion == nil {
		apiVer = apiVersion
	} else {
		apiVer = *config.apiVersion
	}

	if config.notionVersion == nil {
		notionVer = notionVersion
	} else {
		notionVer = *config.notionVersion
	}

	if config.client == nil {
		var retr gent.Retrier
		if config.retrier == nil {
			retr = gent.NewStatusCodeRetrier(
				5,
				retrier.LinearDelay(100*time.Millisecond),
				[]int{500, 503, 502, 504, 408},
			)
		} else {
			retr = config.retrier
		}

		gopts := make([]gent.Option, 0, len(config.gentOptions)+1)
		gopts = append(gopts, gent.UseRetrier(retr))
		gopts = append(gopts, config.gentOptions...)
		client = gent.NewClient(gopts...)
	} else {
		client = config.client
	}

	return &Client{
		client,
		apiKey,
		apiVer,
		notionVer,
	}, nil
}

func WithHTTPClient(client *gent.Client) ClientOption {
	return func(c *Configuration) *Configuration {
		c.client = client
		return c
	}
}

func WithRetrier(retr gent.Retrier) ClientOption {
	return func(c *Configuration) *Configuration {
		c.retrier = retr
		return c
	}
}

func WithGentOptions(options ...gent.Option) ClientOption {
	return func(c *Configuration) *Configuration {
		c.gentOptions = append(c.gentOptions, options...)
		return c
	}
}

func WithApiKey(apiKey string) ClientOption {
	return func(c *Configuration) *Configuration {
		c.apiKey = &apiKey
		return c
	}
}

func WithApiVersion(version string) ClientOption {
	return func(c *Configuration) *Configuration {
		c.apiVersion = &version
		return c
	}
}

func WithNotionVersion(version string) ClientOption {
	return func(c *Configuration) *Configuration {
		c.notionVersion = &version
		return c
	}
}

func New(
	apiKey string,
	options ...ClientOption,
) (*Client, error) {
	config := &Configuration{}

	for idx := range options {
		config = options[idx](config)
	}

	return Build(config)
}
