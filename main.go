package main

import (
	"log"

	"github.com/gin-gonic/gin"
	admin_end "github.com/jacob/modules/modules/admin/api"
	chat_end "github.com/jacob/modules/modules/chat/api"
	filesystem_end "github.com/jacob/modules/modules/filesystem/api"
	home_end "github.com/jacob/modules/modules/home/api"
)

// program entry point
func main() {
	// defaut gin engine
	server := gin.Default()

	// static file path
	server.Static("/static", "./static")

	// modules endpoint
	server.GET("/", home_end.HomeEndpoint)
	server.POST("/upload", filesystem_end.UploadFileEndpoint)
	server.GET("/chatroom", chat_end.ChatRoomEndpoint)
	server.GET("/admin", admin_end.AdminEndpoint)
	server.GET("/admin/:username", admin_end.AdminUserEndpoint)

	// websocket endpoint
	server.GET("/ws", chat_end.ChatWebSocket)

	// message handle
	go chat_end.HandleMessages()

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("無法啟動伺服器: ", err)
	}
}
