package filesystem

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UploadFileData struct {
	User string `json:"user"`
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	str := make([]byte, length)
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}
	return string(str)
}

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "無法獲取檔案")
		return
	}
	defer file.Close()

	newFileName := randomString(8)

	out, err := os.Create("static/" + newFileName + ".png")
	if err != nil {
		log.Fatal("無法建立檔案: ", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal("無法複製檔案內容: ", err)
	}
	log.Printf("檔案已成功上傳: %s", header.Filename)
	data := new(UploadFileData)
	data.User = newFileName
	c.JSON(http.StatusOK, data)
}
