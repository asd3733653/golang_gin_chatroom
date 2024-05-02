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
	data := `
	<!DOCTYPE html>
	<html lang="zh-TW">
	<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>簡單聊天室</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			margin: 0;
			padding: 0;
		}
		#chat-container {
			width: 80%;
			margin: 0 auto;
			padding: 20px;
		}
		#messages {
			border: 1px solid #ccc;
			padding: 10px;
			height: 300px;
			overflow-y: scroll;
		}
		#input-container {
			margin-top: 20px;
		}
	</style>
	</head>
	<body>
	<div id="chat-container">
		<div id="messages"></div>
		<div id="input-container">
			<img src='static/&user&.png' alt="your picture" width="20" height="20">
			<input type="text" id="messageInput" placeholder="輸入訊息...">
			<button onclick="sendMessage()">發送</button>
		</div>
	</div>
	
	<script>
		const socket = new WebSocket("ws://" + window.location.hostname + ":&port&/ws");

		// 監聽來自後端的訊息
		socket.addEventListener("message", function (event) {
			const messageContainer = document.getElementById("messages");
			const message = JSON.parse(event.data);
			const { username, messageText } = message;
			const messageElement = document.createElement("div");

			const imageElement = document.createElement("img");
			imageElement.src = "static/" + username + ".png"; // 設定圖片路徑
			imageElement.style.width = "20px"; // 設定圖片寬度，可依需求調整
			imageElement.style.height = "20px"; // 設定圖片高度，可依需求調整
			messageElement.appendChild(imageElement);
			messageElement.appendChild(document.createTextNode(username + "：" + messageText));
			
			messageContainer.appendChild(messageElement);
		});
	
		// 發送訊息到後端
		function sendMessage() {
			const messageInput = document.getElementById("messageInput");
			const messageText = messageInput.value.trim();
			if (messageText !== "") {
				const message = {
					username: "&user&",
					messageText: messageText
				};
				socket.send(JSON.stringify(message));
				messageInput.value = "";
			}
		}
	</script>
	</body>
	</html>
	`
	data = strings.Replace(data, "&user&", user, -1)
	data = strings.Replace(data, "&port&", port, -1)
	c.Data(http.StatusOK, "chatroom.html", []byte(data))
}

func ChatHandle(c *gin.Context) {
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
