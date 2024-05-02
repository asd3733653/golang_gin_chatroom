package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacob/modules/modules/chat"
)

// AdminQuery 用於接收後台管理員的查詢條件
type AdminQuery struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func AdminEndpoint(c *gin.Context) {

	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")
	// 將字符串轉換為時間
	startTime, err := time.Parse("2006-01-02T15:04:05", startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的起始時間"})
		return
	}
	endTime, err := time.Parse("2006-01-02T15:04:05", endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的結束時間"})
		return
	}

	// 從對話記錄中找出符合查詢條件的訊息
	chatRecords := chat.MessagesRecord
	var result []chat.Message
	for _, msg := range chatRecords {
		if msg.Timestamp.After(startTime) && msg.Timestamp.Before(endTime) {
			result = append(result, msg)
		}
	}

	// 回傳查詢結果
	c.JSON(http.StatusOK, result)
}

// Conversation 定義對話的結構，包含了上一句和下一句訊息
type Conversation struct {
	Previous chat.Message
	Current  chat.Message
	Next     chat.Message
}

func AdminUserEndpoint(c *gin.Context) {
	username := c.Param("username")
	chatRecords := chat.MessagesRecord

	var result []Conversation
	var conv Conversation

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
