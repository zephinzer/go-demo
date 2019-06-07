package main

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"strings"
	"time"
)

type NextHop struct {
	Key string `json:"key"`
	Body string `json:"body"`
	Method string `json:"method"`
	URL string `json:"url"`
	Path string `json:"path"`
}

func (nextHop *NextHop) Request() (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	request, err := http.NewRequest(
		strings.ToUpper(nextHop.Method),
		nextHop.URL,
		bytes.NewBuffer([]byte(nextHop.Body)),
	)
	if err != nil {
		return nil, err
	}
	return httpClient.Do(request)

}