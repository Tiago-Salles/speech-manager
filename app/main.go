package main

import (
	"fmt"
	"speech_manager/server"
)

func main() {
	var server server.ServerHandler = &server.HttpServer{}
	server.InitServer()

	fmt.Print("Initializing server...")
}
