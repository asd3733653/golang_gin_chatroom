package home

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HomeEndpoint(c *gin.Context) {
	// read home html template
	templatePath := "home.html"
	data, err := os.ReadFile("modules/home/template/" + templatePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.Data(http.StatusOK, templatePath, []byte(data))
}
