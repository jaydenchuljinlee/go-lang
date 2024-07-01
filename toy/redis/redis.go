package redis

import (
	"context"
	"log"
	"socket/domain/dto"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
)

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetGoogleInfo(userID int64) (*dto.GoogleInfo, error) {
	key := "user:token_info:GOOGLE:" + strconv.FormatInt(userID, 10)
	var googleInfo dto.GoogleInfo

	err := Rdb.HGetAll(ctx, key).Scan(&googleInfo)
	if err != nil {
		return nil, err
	}
	return &googleInfo, nil
}
