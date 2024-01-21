package test

import (
	"speech_manager/server"
	"testing"
)

func TestHttpServer(t *testing.T) {
	var server server.HttpServer = server.HttpServer{}

	server.InitServer()
}
