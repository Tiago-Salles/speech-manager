package speech

import "speech_manager/server"

type SpeechHandler struct {
}

func registerSpeech(params ...interface{}) server.HttpResponse {
	var response server.HttpResponse = server.HttpResponse{}
	return response
}

func (sH *SpeechHandler) initHttpHandlerMethods() []server.HttpHandler {
	var methods []server.HttpHandler = make([]server.HttpHandler, 0)

	registerSpeech := server.HttpHandler{Path: "/register", Method: registerSpeech}
	methods = append(methods, registerSpeech)

	return methods
}
