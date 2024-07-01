package socket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 업그레이더: HTTP 연결을 웹소켓 연결로 업그레이드
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 도메인 검사를 통해 연결 제한
		return true
	},
}

// 웹소켓 핸들러 함수
func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	for {
		// 클라이언트로부터 메시지 읽기
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// 클라이언트로 메시지 쓰기
		err = conn.WriteMessage(messageType, message)

		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
