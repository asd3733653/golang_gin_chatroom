package admin_model

import chat_model "github.com/jacob/modules/modules/chat/model"

// AdminQuery 用於接收後台管理員的查詢條件
type AdminQuery struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// Conversation 對話結構，包含了上一句和下一句訊息
type Conversation struct {
	Previous chat_model.Message
	Current  chat_model.Message
	Next     chat_model.Message
}
