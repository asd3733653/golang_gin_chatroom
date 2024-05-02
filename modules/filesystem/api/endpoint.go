package filesystem

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	file_model "github.com/jacob/modules/modules/filesystem/model"
	file_service "github.com/jacob/modules/modules/filesystem/service"
)

func UploadFileEndpoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "無法獲取檔案")
		return
	}
	defer file.Close()

	newFileName := file_service.RandomString(8)

	out, err := os.Create("static/" + newFileName + ".png")
	if err != nil {
		log.Fatal("無法建立檔案: ", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal("無法複製檔案內容: ", err)
	}

	log.Printf("檔案已成功上傳: origin_file:%s, file:%s", header.Filename, newFileName)

	data := new(file_model.UploadFileData)
	data.User = newFileName

	c.JSON(http.StatusOK, data)
}
