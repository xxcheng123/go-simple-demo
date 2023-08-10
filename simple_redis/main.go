package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.6:6379",
		Password: "TF4vjV0LfDaJ", // no password set
		DB:       12,             // use default DB
	})
	ctx := context.Background()
	//if str, err := client.HGet(ctx, "user:xxcheng", "name").Result(); err == nil {
	//	fmt.Println(str, err)
	//}
	fmt.Println(int(time.Hour * 4))
	if val, err := client.Set(ctx, "Alist-Path", "https://www.xxcheng.cn", time.Second*10).Result(); err == nil {
		fmt.Println(val, err)
	}
	if val, err := client.Get(ctx, "Alist-Path").Result(); err == redis.Nil {
		fmt.Printf("%s,%s", val, err)
	}

}
