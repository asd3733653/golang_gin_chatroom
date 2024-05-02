package filesystem

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	file_model "github.com/jacob/modules/modules/filesystem/model"
	file_service "github.com/jacob/modules/modules/filesystem/service"
)

func UploadFileEndpoint(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無法獲取檔案" + err.Error()})
		return
	}
	defer file.Close()

	newFileName := file_service.RandomString(8)
	err = file_service.CreateAndCopyFile(newFileName, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	log.Printf("檔案已成功上傳: origin_file:%s, file:%s", header.Filename, newFileName)

	data := new(file_model.UploadFileData)
	data.User = newFileName

	c.JSON(http.StatusOK, data)
}
