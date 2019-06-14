// # traffic-generator
//
// ## running
// use `go run .` in this directory to run the application
//
// ## testing
// test this by starting a server on port 11111 (assuming no
// extra configuration and defaults are in play)
//
package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"go-demo/internal/logger"

	"github.com/sirupsen/logrus"
)

var config *Config
var log *logrus.Logger

func init() {
	log = logger.New()
	config = NewConfig()
	log.Infof("body            : %s", config.Body)
	log.Infof("method          : %s", config.Method)
	log.Infof("req-concurrency : %v", config.RequestConcurrency)
	log.Infof("req-interval    : %v", config.RequestInterval)
	log.Infof("req-rate        : %v", config.RequestRate)
	log.Infof("url             : %s", config.URL)
	rand.Seed(time.Now().Unix())
}

func main() {
	var waiter sync.WaitGroup

	for i := uint(0); i < config.RequestConcurrency; i++ {
		waiter.Add(1)
		<-time.After(time.Duration(math.Floor(float64(config.RequestInterval) * rand.Float64())))
		go repeatedlyRequest(time.Tick(config.RequestInterval), config.RequestRate, config.Method, config.URL, config.Body, i)
	}

	waiter.Wait()
}

func repeatedlyRequest(every <-chan time.Time, rate float64, method, url, body string, threadID uint) {
	for {
		select {
		case <-every:
			randomFloat := rand.Float64()
			if randomFloat < rate {
				timer := time.Now()
				log.Infof("sending '%s' request to '%s' with data '%s'", method, url, body)
				req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
				if err != nil {
					log.Error(err)
					continue
				}
				client := http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
						},
					},
				}
				requestCreated := time.Since(timer)
				timer = time.Now()
				res, err := client.Do(req)
				if err != nil {
					log.Error(err)
					continue
				}
				requestFulfilled := time.Since(timer)
				timer = time.Now()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Error(err)
					continue
				}
				requestParsed := time.Since(timer)
				log.WithFields(logrus.Fields{
					"requestCreated": requestCreated,
					"requestFulfilled": requestFulfilled,
					"requestParsed": requestParsed,
				}).Infof("url '%s' responded with '%v'", url, res.StatusCode)
				log.Tracef("response: '%s'", string(body))
			} else {
				log.Infof("skipping request due to threshold not being met (generated: %v, threshold: %v)", randomFloat, rate)
			}
		}
	}
}

