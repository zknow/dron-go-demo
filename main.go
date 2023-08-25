package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	txt := "HELLO_WORLD_To_小蔡!!"

	fmt.Println(txt)

	err := os.WriteFile("data.txt", []byte(txt), 777)
	if err != nil {
		log.Fatal(err)
	}
}
