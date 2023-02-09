package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   0,
		},
	)
	ctx := context.Background()
	rdb.Set(ctx, "wow:a", "12", time.Hour)
	rdb.Set(ctx, "wow:b", "12", time.Hour)
	rdb.Set(ctx, "wow:c", "12", 0)
	rdb.Set(ctx, "wow:d", "12", time.Hour)
	rdb.Set(ctx, "wow:e", "12", time.Hour)
	rdb.Set(ctx, "wow:f", "12", time.Hour)

	pipe := rdb.Pipeline()
	pipe.TTL(ctx, "wow:a")
	pipe.TTL(ctx, "wow:b")
	pipe.TTL(ctx, "wow:c")
	res, err := pipe.Exec(ctx)
	if err != nil {
		println("failure lmao")
	}
	for idx := range res {
		dur, ok :=  (res[idx]).(*redis.DurationCmd)
		if !ok {
			println("wtf why?")
		} else {
			print((dur.Args()[1]).(string))
			println(dur.Val())
		}
	}
	println("lmao")
}

type Client struct {
}

type cmdable func(ctx context.Context, str string) error
