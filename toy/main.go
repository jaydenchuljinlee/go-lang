package main

import (
	"toy/config"
	"toy/redis"
	"toy/repository"
	"toy/router"
)

func main() {
	config.LoadConfig()
	// Redis 초기화
	redis.Init()

	// DB
	repository.InitDB()
	defer repository.DB.Close()

	router.InitRouter()
}
