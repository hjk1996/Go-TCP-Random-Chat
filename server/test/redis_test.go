package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSetAndGetValue(t *testing.T) {
	ctx := context.Background()

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		},
	)

	key := "testKey"
	value := "testValue"

	// Set value
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		t.Fatalf("Could not set value: %v", err)
	}

	// Get value
	got, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("Could not get value: %v", err)
	}

	if got != value {
		t.Errorf("Expected value %s, got %s", value, got)
	}
}
