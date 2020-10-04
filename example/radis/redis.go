package main

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"go-transportation-bot/pkg/modules/cache"
)

func main() {
	//klog.InitFlags(nil)
	flag.Set("v", "3")
	flag.Parse()
	ctx := context.Background()

	manger := cache.GetManager()
	client := manger.GetRedisClient("localhost:6379")
	err := client.Set(ctx, "Key", "Value", 0).Err()
	if err != nil {
		//klog.V(3).ErrorS(err, "an error occurred")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		//klog.V(3).ErrorS(err, "an error occurred")
	}
}
