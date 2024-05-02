package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestData struct {
	Hello string `json:"hello"`
}

func TestEndpoint(c *gin.Context) {
	data := new(TestData)
	data.Hello = "world!"
	c.JSON(http.StatusOK, data)
}
