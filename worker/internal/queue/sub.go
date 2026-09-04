package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Sub(ctx context.Context, channel string) <-chan *redis.Message {

	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	sub := rds.Subscribe(ctx, channel)

	go func() {
		<-ctx.Done()
		sub.Close()
	}()

	return sub.Channel()
}
