package main

import (
	"dron-go-demo/util"
	"net/http"
)

// 2023/11/17-1
func main() {
	http.HandleFunc("/", util.SayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
