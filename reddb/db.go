package reddb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    fmt.Print("Redis client created")
}

func SetOrder(orderID, orderStatus string) error {
    return rdb.Set(context.Background(), orderID, orderStatus, 0).Err()
}

func GetOrder(orderID string) (string, error) {
    orderStatus, err := rdb.Get(context.Background(), orderID).Result()
    if err == redis.Nil {
        return "", err
    } else if err != nil {
        return "", err
    }
    return orderStatus, nil
}
