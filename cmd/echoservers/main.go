// # echoservers
//
// ## pre-requisites
// generate the certificates by running `make ssl` from the project root
//
// ## running
// use `go run .` in this directory to run the server
//
// ## testing
// test this using `curl -k https://localhost:11111`
//
package main

import (
	"log"
	"net/http"
	"time"

	"go-demo/internal/echo"
	"go-demo/internal/config"

	"github.com/gorilla/mux"
)

func main() {
	serverConfig := config.NewServer()
	tlsKeyPair := config.NewTLSKeyPair()
	handler := mux.NewRouter()
	handler.NotFoundHandler = echo.HandlerFunc
	addr := serverConfig.GetAddr()

	server := http.Server{
		Addr:              addr,
		Handler:           handler,
		MaxHeaderBytes:  	1024,
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Printf("starting echoservers at https://%s...", addr)
	err := server.ListenAndServeTLS(tlsKeyPair.CertPath, tlsKeyPair.KeyPath)
	if err != nil {
		log.Panic(err)
	}
}
