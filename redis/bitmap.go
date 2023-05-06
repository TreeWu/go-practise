package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.64.3:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.SetBit(ctx, "bitmap", 2, 1).Err()
	if err != nil {
		fmt.Println(err)
	}

	result, err := rdb.GetBit(ctx, "bitmap", 2).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("get bitmap 2 ", result)

	result, err = rdb.GetBit(ctx, "bitmap", 100).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("get bitmap 100 ", result)
}
