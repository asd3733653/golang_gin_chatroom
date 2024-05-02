package chat_model

import "time"

type Message struct {
	Username    string    `json:"username"`
	MessageText string    `json:"messageText"`
	Timestamp   time.Time `json:"timestamp"`
}
