package main

import (
	"time"
	"github.com/spf13/viper"
)

func NewConfig() *Config {
	config := viper.New()
	config.SetDefault("body", "")
	config.SetDefault("method", "GET")
	config.SetDefault("request_concurrency", 1)
	config.SetDefault("request_rate", 1.0)
	config.SetDefault("request_interval", 1*time.Second)
	config.SetDefault("url", "https://localhost:11111")
	config.AutomaticEnv()

	return &Config{
		Body:               config.GetString("body"),
		Method:             config.GetString("method"),
		RequestConcurrency: config.GetUint("request_concurrency"),
		RequestInterval:    config.GetDuration("request_interval"),
		RequestRate:        config.GetFloat64("request_rate"),
		URL:                config.GetString("url"),
	}
}

type Config struct {
	// Body of the request to be sent
	Body string `json:"body"`

	// Method to use when sending the HTTP request
	Method string `json:"method"`

	// RequestConcurrency specifies how many requester threads should be running at once
	RequestConcurrency uint `json:"requestConcurrency"`

	// RequestInterval specifies a duration between requester threads triggering a request to the .URL
	RequestInterval time.Duration `json:"requestInterval"`

	// RequestRate specifies a number from 0.0 - 1.0 which indicates the percentage of times the request is actually sent
	RequestRate float64 `json:"requestRate"`

	// URL specifies where to send the request to
	URL string `json:"url"`
}
