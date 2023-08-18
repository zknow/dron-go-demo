package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.String(http.StatusOK, "hello1")
}
func main() {
	// test1
	server := gin.Default()
	server.GET("/", test)
	server.Run(":8888")
}
