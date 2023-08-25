package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func runGin() {
	server := gin.Default()
	server.GET("/", test)
	server.Run(":8888")
}

func main() {
	fmt.Println("HELLO_WORLD_To_小蔡!!")
}
