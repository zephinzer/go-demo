// # fwdserver
//
// # running
// use `go run .` in this directory to run the server
//
// # testing
// test this using `curl -k https://localhost:11111/<registered server>`
//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go-demo/internal/config"
	"go-demo/internal/logger"

	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var nextHopConfig = viper.New()

func init() {
	pflag.Duration("auto-ping", 0, "when set together with NEXT_HOP_URL, sends a request to the next hop URL at the specified duration interval")
	pflag.Parse()
	nextHopConfig.BindPFlags(pflag.CommandLine)
	nextHopConfig.SetDefault("next_hop_url", "")
	nextHopConfig.AutomaticEnv()
	if nextHopConfig.GetDuration("auto-ping") > 0 && nextHopConfig.GetString("next_hop_url") == "" {
		log.Fatalln("flag:'--auto-ping' was set but envvar:'NEXT_HOP_URL' seems to be empty - fwdserver needs to know who to auto-ping!")
	}
}

func main() {
	serverConfig := config.NewServer()
	tlsKeyPair := config.NewTLSKeyPair()
	log.Printf("host: %s", serverConfig.Host)
	log.Printf("port: %v", serverConfig.Port)
	log.Printf("cert: %s", tlsKeyPair.CertPath)
	log.Printf("key:  %s", tlsKeyPair.KeyPath)
	log.Printf("next_hop_url: %s", nextHopConfig.GetString("next_hop_url"))
	log.Printf("auto-ping:    %v", nextHopConfig.GetDuration("auto-ping"))

	if nextHopConfig.GetDuration("auto-ping") > 0 && nextHopConfig.GetString("next_hop_url") != "" {
		nextHopUrl := nextHopConfig.GetString("next_hop_url")
		go func(every <-chan time.Time) {
			for {
				select {
				case <-every:
					log.Printf("pinging %s...", nextHopUrl)
					nextHop := &NextHop{
						Key: "next_hop_url",
						Body: "",
						URL: nextHopConfig.GetString("next_hop_url"),
						Method: "GET",
						Path: "/",
					}
					nextHopRes, err := nextHop.Request()
					if err != nil {
						log.Println(err)
					} else {
						log.Println(nextHopRes)
					}

				}
			}
		}(time.Tick(nextHopConfig.GetDuration("auto-ping")))
	}

	handler := mux.NewRouter()
	handler.Handle(`/{nextHop}`, getNextHopHandler())
	handler.Handle(`/{nextHop}/{path:.+}`, getNextHopHandler())
	addr := serverConfig.GetAddr()
	server := http.Server{
		Addr:              addr,
		Handler:           logger.ServerMiddleware(handler),
		MaxHeaderBytes:  	 1024,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	var serverStartError error
	if tlsKeyPair.Exists() {
		log.Printf("starting pingserver at https://%s...", addr)
		serverStartError = server.ListenAndServeTLS(tlsKeyPair.CertPath, tlsKeyPair.KeyPath)
	} else {
		log.Printf("starting pingserver at http://%s...", addr)
		serverStartError = server.ListenAndServe()
	}
	if serverStartError != nil {
		log.Panic(serverStartError)
	}
}

func getNextHopHandler() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var err error
		params := mux.Vars(r)
		contentLength := r.Header.Get("content-length")
		contentEncoding := r.Header.Get("content-encoding")
		nextHop := NextHop{
			Key: params["nextHop"],
			Method: r.Method,
			Path: params["path"],
		}
	
		// assign url iff nextHop.Key is valid
		var nextHopURL string
		if nextHop.Key != "" {
			if nextHopURL = nextHopConfig.GetString(nextHop.Key); nextHopURL == "" {
				if nextHopURL = nextHopConfig.GetString("next_hop_url"); nextHopURL == "" {
					err = fmt.Errorf("no url found for key '%s'", nextHop.Key)
					handleError(w, err, &nextHop)
					return
				}
			}
			nextHop.URL = nextHopURL
		}

		// assign body if "Content-(Length|Encoding)" http headers are available
		if contentLength != ""  || contentEncoding != "" {
			nextHopBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				handleError(w, err, &nextHop)
				return
			}
			nextHop.Body = string(nextHopBody)
		}

		response, err := nextHop.Request()
		if err != nil {
			handleError(w, err, &nextHop)
			return
		}

		var responseBody []byte
		responseContentLength := response.Header.Get("Content-Length")
		responseContentEncoding := response.Header.Get("Content-Encoding")
		if responseContentLength != "" || responseContentEncoding != "" {
			responseBody, err = ioutil.ReadAll(response.Body)
			if err != nil {
				handleError(w, err, &nextHop)
				return
			}
		}

		var responseData interface{}
		var unmarshalledResponse map[string]interface{}
		err = json.Unmarshal(responseBody, &unmarshalledResponse)
		if err != nil {
			responseData = string(responseBody)
		} else {
			responseData = unmarshalledResponse
		}
		w.Write(Response{
			Data: responseData,
			NextHop: &nextHop,
		}.ToBytes())
	})
}

func handleError(w http.ResponseWriter, err error, nextHop *NextHop) {
	log.Println(err)
	w.Write(Response{
		Error: err.Error(),
		NextHop: nextHop,
	}.ToBytes())
}
