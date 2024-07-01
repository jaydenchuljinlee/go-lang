package router

import (
	"socket/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/userinfo", controller.GetUserInfo)

	// 웹소켓 엔드포인트
	// router.GET("/ws", socket.WsHandler)
	// router.GET("/redis", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	router.Run(":9998")

	return router
}
