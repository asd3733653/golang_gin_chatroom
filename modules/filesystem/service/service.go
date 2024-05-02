package file_service

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"time"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	str := make([]byte, length)
	for i := range str {
		str[i] = charset[rng.Intn(len(charset))]
	}
	return string(str)
}

func CreateAndCopyFile(newFileName string, file multipart.File) error {
	// 建立新檔案
	out, err := os.Create("static/" + newFileName + ".png")
	if err != nil {
		return fmt.Errorf("無法建立檔案: %w", err)
	}
	defer out.Close()

	// 複製檔案內容
	_, err = io.Copy(out, file)
	if err != nil {
		return fmt.Errorf("無法複製檔案內容: %w", err)
	}
	return nil
}
