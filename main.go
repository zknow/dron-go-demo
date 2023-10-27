package main

import (
	"dron-go-demo/util"
	"net/http"
)

// 2023/10/27-1
func main() {
	http.HandleFunc("/", util.SayHello)
	if err := http.ListenAndServe(":12003", nil); err != nil {
		panic(err)
	}
}
