package main

import (
	"socket/redis"
	"socket/repository"
	"socket/router"
)

func main() {
	// Redis 초기화
	redis.Init()

	// DB
	repository.InitDB()
	defer repository.DB.Close()

	router.InitRouter()
}
