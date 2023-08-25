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

	err1 := os.WriteFile("/tmp/data.txt", []byte(txt), 777)
	if err1 != nil {
		log.Fatal(err1)
	}

	err2 := os.WriteFile("/var/tmp/data.txt", []byte(txt), 777)
	if err2 != nil {
		log.Fatal(err2)
	}
}
