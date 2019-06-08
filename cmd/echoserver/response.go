package main

import (
	"encoding/json"
)

type Response struct {
	ContentLength int64 `json:"contentLength"`
	Headers map[string]string `json:"headers"`
	Host string `json:"host"`
	Port uint64 `json:"port"`
	Method string `json:"method"`
	Protocol string `json:"protocol"`
	RemoteAddr string `json:"remoteAddr"`
	RequestURI string `json:"requestUri"`
	Timestamp string `json:"timestamp"`
	TLSEnabled bool `json:"tlsEnabled"`
}

type ErrorResponse struct {
	StatusCode uint `json:"statusCode"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) []byte {
	errorResponse := ErrorResponse{
		StatusCode: 500,
		Message: message,
	}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		return []byte("something went seriously wrong")
	}
	return response
}
