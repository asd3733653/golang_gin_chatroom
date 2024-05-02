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
