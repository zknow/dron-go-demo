package main

import (
	"dron-go-demo/util"
	"net/http"
)

// 2023/09/22
func main() {
	http.HandleFunc("/", util.SayHello)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
