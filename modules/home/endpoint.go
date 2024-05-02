package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeEndpoint(c *gin.Context) {
	data := `
	<!DOCTYPE html>
	<html lang="zh-TW">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>檔案上傳</title>
	</head>
	<body>
		<form id="uploadForm" enctype="multipart/form-data">
			<input type="file" name="file" accept="image/png" id="file">
			<button type="submit">上傳檔案</button>
		</form>
	
		<script>
		document.getElementById("uploadForm").addEventListener("submit", function(event) {
			event.preventDefault(); // 防止表單直接提交
		
			var formData = new FormData(this);
		
			// 發送 POST 請求
			var xhr = new XMLHttpRequest();
			xhr.open("POST", "/upload", true);
			xhr.onreadystatechange = function() {
				if (xhr.readyState === XMLHttpRequest.DONE) {
					if (xhr.status === 200) {
						console.log("上傳成功!");
						json = JSON.parse(xhr.response)
						window.location.href = "/chatroom?user=" + json.user;
					} else {
						console.error("發生錯誤:", xhr.statusText);
					}
				}
			};
			xhr.send(formData);
		});
		</script>
	</body>
	</html>
	`
	c.Data(http.StatusOK, "home.html", []byte(data))
}
