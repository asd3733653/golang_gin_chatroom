package chat

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Username    string    `json:"username"`
	MessageText string    `json:"messageText"`
	Timestamp   time.Time `json:"timestamp"`
}

var (
	clients           = make(map[*websocket.Conn]bool)
	broadcast         = make(chan Message)
	MessagesRecord    []Message
	MessagesRecordMux sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer ws.Close()

	clients[ws] = true

	// 將最近的聊天訊息發送給新連接的客戶端
	MessagesRecordMux.Lock()
	start := 0
	if len(MessagesRecord) > 10 {
		start = len(MessagesRecord) - 10
	}
	print(start)
	for _, msg := range MessagesRecord[start:] {
		err := ws.WriteJSON(msg)
		if err != nil {
			log.Println("無法發送訊息給客戶端:", err)
			break
		}
	}
	MessagesRecordMux.Unlock()

	for {
		var msg Message
		msg.Timestamp = time.Now()
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		// 將訊息添加到最近的聊天訊息中
		MessagesRecordMux.Lock()
		MessagesRecord = append(MessagesRecord, msg)
		MessagesRecordMux.Unlock()

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				// 如果發送失敗，刪除該客戶端
				client.Close()
				delete(clients, client)
			}
		}
	}
}
