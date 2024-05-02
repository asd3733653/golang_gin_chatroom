package chat_service

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	chat_model "github.com/jacob/modules/modules/chat/model"
)

var (
	Clients           = make(map[*websocket.Conn]bool)
	Broadcast         = make(chan chat_model.Message)
	MessagesRecord    []chat_model.Message
	MessagesRecordMux sync.Mutex
)

// 升級連接 http to ws
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleMessages() {
	for {
		msg := <-Broadcast

		// 將訊息添加到最近的聊天訊息中
		MessagesRecordMux.Lock()
		MessagesRecord = append(MessagesRecord, msg)
		MessagesRecordMux.Unlock()

		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				// 如果發送失敗，刪除該客戶端
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
