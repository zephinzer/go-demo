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
	"log"
	"net/http"
	"time"

	"go-demo/internal/echo"
	"go-demo/internal/config"
	"go-demo/internal/logger"

	"github.com/gorilla/mux"
)

func main() {
	serverConfig := config.NewServer()
	handler := mux.NewRouter()
	handler.NotFoundHandler = echo.HandlerFunc
	addr := serverConfig.GetAddr()
	server := http.Server{
		Addr:              addr,
		Handler:           logger.ServerMiddleware(handler),
		MaxHeaderBytes: 	 1024,
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Printf("starting echoserver at http://%s...", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
