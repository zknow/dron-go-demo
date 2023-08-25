package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	t := time.Now().UTC().Add(time.Hour * 8)
	txt := "Ping" + t.String()
	c.String(http.StatusOK, txt)
	wlog(txt)
}

func runGin() {
	server := gin.Default()
	server.GET("/", index)
	server.Run(":8888")
}

func wlog(txt string) {

	log.Println(txt)
	err1 := os.WriteFile("/tmp/data.txt", []byte(txt), 777)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func main() {
	runGin()
}
