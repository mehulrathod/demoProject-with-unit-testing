package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetValue(key string, value interface{}) error {
	b,err :=  json.Marshal(value)
	//if err  {
	//
	//}

	err = rdb.Set(ctx, key, string(b), 20*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetValue(key string) interface{} {
	var response []interface{}
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("redis file", res)
	json.Unmarshal([]byte(res), &response)
	return response
}
