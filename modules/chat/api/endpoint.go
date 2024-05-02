package chat

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	chat_model "github.com/jacob/modules/modules/chat/model"
	chat_service "github.com/jacob/modules/modules/chat/service"
)

func ChatRoomEndpoint(c *gin.Context) {
	user := c.Query("user")
	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = "8080"
	}

	// read home html template
	templatePath := "chatRoom.html"
	data, err := os.ReadFile("modules/chat/template/" + templatePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	template := string(data)
	template = strings.Replace(template, "&user&", user, -1)
	template = strings.Replace(template, "&port&", port, -1)
	c.Data(http.StatusOK, templatePath, []byte(template))
}

func ChatWebSocket(c *gin.Context) {
	ws, err := chat_service.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer ws.Close()

	chat_service.Clients[ws] = true

	// 將最近的聊天訊息發送給新連接的客戶端
	chat_service.MessagesRecordMux.Lock()
	start := 0
	if len(chat_service.MessagesRecord) > 10 {
		start = len(chat_service.MessagesRecord) - 10
	}
	print(start)
	for _, msg := range chat_service.MessagesRecord[start:] {
		err := ws.WriteJSON(msg)
		if err != nil {
			log.Println("無法發送訊息給客戶端:", err)
			break
		}
	}
	chat_service.MessagesRecordMux.Unlock()

	for {
		var msg chat_model.Message
		msg.Timestamp = time.Now()
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(chat_service.Clients, ws)
			break
		}
		chat_service.Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-chat_service.Broadcast

		// 將訊息添加到最近的聊天訊息中
		chat_service.MessagesRecordMux.Lock()
		chat_service.MessagesRecord = append(chat_service.MessagesRecord, msg)
		chat_service.MessagesRecordMux.Unlock()

		for client := range chat_service.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				// 如果發送失敗，刪除該客戶端
				client.Close()
				delete(chat_service.Clients, client)
			}
		}
	}
}
