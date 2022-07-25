package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func main() {
  rdb := redis.NewClient()
  rdb.Set()
  rdb.AddHook()
}


type Client struct {

}

type cmdable func(ctx context.Context, str string) error


