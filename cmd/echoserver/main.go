// # echoserver
//
// ## running
// use `go run .` in this directory to run the server
//
// ## testing
// test this using `curl http://localhost:11111`
//
package main

import (
	"net/http"
	"encoding/json"
	"time"

	"go-demo/internal/config"
	"go-demo/internal/logger"
	"go-demo/internal/server"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var serverConfig config.Server
var tlsKeyPairConfig config.TLSKeyPair
var log *logrus.Logger

func init() {
	log = logger.New()
	serverConfig = config.NewServer()
	tlsKeyPairConfig = config.NewTLSKeyPair()
	log.Infof("host         : %s", serverConfig.Host)
	log.Infof("port         : %v", serverConfig.Port)
	log.Infof("cert         : %s", tlsKeyPairConfig.CertPath)
	log.Infof("key          : %s", tlsKeyPairConfig.KeyPath)
}

func main() {
	handler := mux.NewRouter()
	handler.NotFoundHandler = getEchoHandler()
	if err := server.Start("echoserver", server.New(handler)); err != nil {
		panic(err)
	}
}

func getEchoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
}
