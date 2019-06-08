package server

import (
	"net/http"
	"time"
	"go-demo/internal/logger"
	"go-demo/internal/config"
)

const MaxHeaderBytes = 1024
var (
	ReadHeaderTimeout = 5 * time.Second
	ReadTimeout = 10 * time.Second
	WriteTimeout = 10 * time.Second
	IdleTimeout = 10 * time.Second
	log = logger.New()
)

func New(handler http.Handler) http.Server {
	serverConfig := config.NewServer()
	return http.Server{
		Addr:              serverConfig.GetAddr(),
		Handler:           logger.ServerMiddleware(handler),
		MaxHeaderBytes:  	 1024,
		ReadHeaderTimeout: ReadHeaderTimeout,
		ReadTimeout:       ReadTimeout,
		WriteTimeout:      WriteTimeout,
		IdleTimeout:       IdleTimeout,
	}
}

func Start(applicationName string, server http.Server) error {
	tlsKeyPairConfig := config.NewTLSKeyPair()
	var serverStartError error
	if tlsKeyPairConfig.Exists() {
		log.Printf("starting %s at https://%s...", applicationName, server.Addr)
		serverStartError = server.ListenAndServeTLS(tlsKeyPairConfig.CertPath, tlsKeyPairConfig.KeyPath)
	} else {
		log.Printf("starting %s at http://%s...", applicationName, server.Addr)
		serverStartError = server.ListenAndServe()
	}
	if serverStartError != nil {
		log.Fatal(serverStartError)
	}
	return serverStartError
}
