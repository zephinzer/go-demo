package echo

import (
	"encoding/json"
	"net/http"
	"time"
)

var HandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]string)
	for name, values := range r.Header{
		for _, value := range values {
			headers[name] = value
		}
	}
	response := Response{
		ContentLength: r.ContentLength,
		Headers: headers,
		Host: r.Host,
		Method: r.Method,
		Protocol: r.Proto,
		RemoteAddr: r.RemoteAddr,
		RequestURI: r.RequestURI,
		Timestamp: time.Now().Format(time.RFC3339),
		TLSEnabled: r.TLS != nil,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		w.Write(NewErrorResponse(err.Error()))
	}
	w.WriteHeader(200)
	w.Write(responseBytes)
})