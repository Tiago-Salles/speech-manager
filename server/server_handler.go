package server

type ServerHandler interface {
	InitServer()
	InitializeHandlers() []HttpHandler
}
