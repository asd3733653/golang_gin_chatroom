package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jacob/modules/modules/admin"
	"github.com/jacob/modules/modules/chat"
	"github.com/jacob/modules/modules/chatroom"
	"github.com/jacob/modules/modules/filesystem"
	"github.com/jacob/modules/modules/home"
	"github.com/jacob/modules/modules/test"
)

func main() {
	// defaut gin engine
	server := gin.Default()

	// static file
	server.Static("/static", "./static")

	// test endpoint health endpoint
	server.GET("/test", test.TestEndpoint)

	// real module
	server.GET("/", home.HomeEndpoint)
	server.GET("/chatroom", chatroom.ChatRoomEndpoint)
	server.GET("/admin", admin.AdminEndpoint)
	server.GET("/admin/:username", admin.AdminUserEndpoint)

	// websocket
	server.GET("/ws", chat.HandleConnections)

	// function endpoint
	server.POST("/upload", filesystem.UploadHandler)

	go chat.HandleMessages()

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("無法啟動伺服器: ", err)
	}
}
