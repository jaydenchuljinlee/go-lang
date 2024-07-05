package main

import (
	"log"
	"sso/config"
	"sso/router"
)

func main() {
	config.LoadConfig()

	r := router.InitRouter()

	// 서버를 시작합니다. 포트 번호는 필요에 따라 변경할 수 있습니다.
	err := r.Run(":9998")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
