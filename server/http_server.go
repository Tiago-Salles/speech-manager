package server

import (
	"net/http"
)

type httpMethodHandler func(http.ResponseWriter, *http.Request)

type HttpResponse struct {
	statusCode int
	message    string
	header     map[string]string
}

type HttpHandler struct {
	Path   string
	Method func(params ...interface{}) HttpResponse
}

type HttpServer struct {
	ServerHandler
}

func makeAsHttpHandler(f func(params ...interface{}) HttpResponse) httpMethodHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var response HttpResponse = f(r.Body)
		for k, v := range response.header {
			w.Header().Add(k, v)
		}
		w.WriteHeader(response.statusCode)
		w.Write([]byte(response.message))
	}
}

func (s *HttpServer) InitServer() {
	for _, e := range s.InitializeHandlers() {
		http.HandleFunc(e.Path, makeAsHttpHandler(e.Method))
	}
}

func (s *HttpServer) InitializeHandlers() []HttpHandler {
	var handlers []HttpHandler = make([]HttpHandler, 0)
	handler := HttpHandler{}
	handlers = append(handlers, handler)
	return handlers
}
