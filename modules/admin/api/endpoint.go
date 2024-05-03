package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	admin_model "github.com/jacob/modules/modules/admin/model"
	admin_service "github.com/jacob/modules/modules/admin/service"

	chat_model "github.com/jacob/modules/modules/chat/model"
	chat_service "github.com/jacob/modules/modules/chat/service"
)

func AdminEndpoint(c *gin.Context) {
	// define AdminQuery mapper 查詢物件
	var query admin_model.AdminQuery

	// 解析用戶輸入的查詢條件
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的查詢條件"})
		return
	}

	// AdminQuery validator
	startTime, endTime, err := admin_service.AdminQueryValidator(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// 從對話記錄中找出符合查詢條件的訊息
	chatRecords := chat_service.MessagesRecord
	var result []chat_model.Message
	for _, msg := range chatRecords {
		if msg.Timestamp.After(startTime) && msg.Timestamp.Before(endTime) {
			result = append(result, msg)
		}
	}

	// 回傳查詢結果
	c.JSON(http.StatusOK, result)
}

func AdminUserEndpoint(c *gin.Context) {
	username := c.Param("username")
	chatRecords := chat_service.MessagesRecord

	var result []admin_model.Conversation
	var conv admin_model.Conversation

	// 遍歷對話記錄
	for i, msg := range chatRecords {
		if msg.Username == username {
			if i > 0 {
				conv.Previous = chatRecords[i-1]
			}
			conv.Current = msg
			if i < len(chatRecords)-1 {
				conv.Next = chatRecords[i+1]
			}
			result = append(result, conv)
		}
	}

	// 回傳查詢結果
	c.JSON(http.StatusOK, result)
}
