package redis

import (
	"context"
	"log"
	"strconv"
	"toy/config"
	"toy/domain/dto"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	Rdb *redis.Client
)

func Init() {
	cfg := config.AppConfig.Redis

	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
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
