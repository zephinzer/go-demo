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
	"github.com/spf13/viper"
)

var serverConfig config.Server
var tlsKeyPairConfig config.TLSKeyPair
var healthCheckConfig *viper.Viper
var log *logrus.Logger

func init() {
	log = logger.New()
	serverConfig = config.NewServer()
	tlsKeyPairConfig = config.NewTLSKeyPair()
	healthCheckConfig = viper.New()
	healthCheckConfig.SetDefault("alive", true)
	healthCheckConfig.SetDefault("ready", true)
	log.Infof("name  : %s", serverConfig.Name)
	log.Infof("host  : %s", serverConfig.Host)
	log.Infof("port  : %v", serverConfig.Port)
	log.Infof("cert  : %s", tlsKeyPairConfig.CertPath)
	log.Infof("key   : %s", tlsKeyPairConfig.KeyPath)
	log.Infof("alive : %v", healthCheckConfig.GetBool("alive"))
	log.Infof("ready : %v", healthCheckConfig.GetBool("ready"))
}

func main() {
	handler := mux.NewRouter()
	handler.Handle("/liveness", getLivenessCheckHandler())
	handler.Handle("/liveness/{status}", setLivenessCheckHandler())
	handler.Handle("/readiness", getReadinessCheckHandler())
	handler.Handle("/readiness/{status}", setReadinessCheckHandler())
	go reportStatusPeriodically(time.Tick(3 * time.Second))
	if err := server.Start(serverConfig.Name, server.New(handler)); err != nil {
		log.Panic(err)
	}
}

func reportStatusPeriodically(every <-chan time.Time) {
	for {
		select {
		case <-every:
			healthCheckConfig.AutomaticEnv()
			log.Debugf("alive/ready: %v/%v", healthCheckConfig.GetBool("alive"), healthCheckConfig.GetBool("ready"))
		}
	}
}

func setLivenessCheckHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				respondNotOK(w)
			}
		}()
		params := mux.Vars(r)
		healthCheckConfig.Set("alive", params["status"])
		respondOK(w)
	})
}

func getLivenessCheckHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		healthCheckConfig.AutomaticEnv()
		if healthCheckConfig.GetBool("alive") {
			respondOK(w)
		} else {
			respondNotOK(w)
		}
	})
}

func setReadinessCheckHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				respondNotOK(w)
			}
		}()
		params := mux.Vars(r)
		healthCheckConfig.Set("ready", params["status"])
		respondOK(w)
	})
}

func getReadinessCheckHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		healthCheckConfig.AutomaticEnv()
		if healthCheckConfig.GetBool("ready") {
			respondOK(w)
		} else {
			respondNotOK(w)
		}
	})
}

func respondOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	response := Response{Status: "ok"}
	w.Write(response.ToBytes())
}

func respondNotOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	response := Response{Status: "not ok"}
	w.Write(response.ToBytes())
}

type Response struct {
	Status string `json:"status"`
	Alive bool `json:"alive"`
	Ready bool `json:"ready"`
}

func (res *Response) ToBytes() []byte {
	res.Alive = healthCheckConfig.GetBool("alive")
	res.Ready = healthCheckConfig.GetBool("ready")
	response, err := json.Marshal(res)
	if err != nil {
		response, err = json.Marshal(Response{
			Status: err.Error(),
			Alive: res.Alive,
			Ready: res.Ready,
		})
	}
	return response
}