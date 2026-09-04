package queue

import (
	"github.com/redis/go-redis/v9"
	"context"
)

func Publish(path string) (bool, error){
	ctx := context.Background()
	rds := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	err :=rds.Publish(
		ctx,
		"fileParse",
		path,
	).Err()

	if err != nil {
		return false, err
	}
	
	return true, nil
}