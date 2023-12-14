package reddb

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
    fmt.Println("Attempting to create Redis client")
    rdb = redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    })

    _, err := rdb.Ping(context.Background()).Result()
    if err != nil {
        fmt.Println("Failed to create Redis client:", err)
    } else {
        fmt.Println("Redis client created")
    }
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
