package chatroom

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ChatRoomEndpoint(c *gin.Context) {
	user := c.Query("user")
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
		const socket = new WebSocket("ws://" + window.location.hostname + ":32161/ws");

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
	c.Data(http.StatusOK, "chatroom.html", []byte(data))
}
