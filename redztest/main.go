package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	rop := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	client := redis.NewClient(
		rop,
	)
	err := client.Ping(context.Background()).Err()
	if err != nil {
	  panic(err)
	}

	scan(client)
}

func add(client *redis.Client) {
  d := map[string]interface{}{
	  "courtroomId": "xyz-abc-325",
	}
	dj, _ := json.Marshal(d)
	err := client.ZAdd(
	  context.Background(),
	  "ztesting",
	  &redis.Z{
	    Member: dj,
	  },
	).Err() 
	if err != nil {
	  panic(err)
	}
}

func scan(client *redis.Client) {
  res, cur, err := client.ZScan(
    context.Background(),
    "ztestingg",
    0,
    "*\"xyz-abc-322\"*",
    100,
  ).Result()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%v\n", res)
  fmt.Printf("%v\n", cur)
}
