package main

import (
	"log"

	"github.com/gin-gonic/gin"
	admin_end "github.com/jacob/modules/modules/admin/api"
	chat_end "github.com/jacob/modules/modules/chat/api"
	chat_service "github.com/jacob/modules/modules/chat/service"
	filesystem_end "github.com/jacob/modules/modules/filesystem/api"
	home_end "github.com/jacob/modules/modules/home/api"
)

// program entry pointＦ
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

	// websocket
	server.GET("/ws", chat_end.ChatHandle)

	go chat_service.HandleMessages()

	err := server.Run(":8080")
	if err != nil {
		log.Fatal("無法啟動伺服器: ", err)
	}
}
