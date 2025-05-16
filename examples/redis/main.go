package main

import (
	"context"
	"time"

	"github.com/gscaffold/helpers/databases/redis"
	"github.com/gscaffold/helpers/logger"
)

func main() {
	rds := redis.MustDiscovery("test")
	ctx := context.TODO()

	err := rds.Set(ctx, "test_set", "test_value", time.Hour).Err()
	if err != nil {
		panic(err)
	}

	value, err := rds.Get(ctx, "test_set").Result()
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "get value:%s", value)
}
