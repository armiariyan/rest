package rest

import (
	"net/http"
	"time"
)

type ConfigTransport struct {
	TLSHandshakeTimeout int `json:"tls_handshake_timeout"` // in second, default 5 seconds
}

type ConfigClient struct {
	Timeout int `json:"timeout"` // in second, default 30 seconds
}

type Config struct {
	Transport ConfigTransport `json:"transport"`
	Client    ConfigClient    `json:"client"`
}

// NewHttpClient return golang native http client
func NewHttpClient(config Config) *http.Client {
	if config.Transport.TLSHandshakeTimeout <= 0 {
		config.Transport.TLSHandshakeTimeout = 5
	}

	if config.Client.Timeout <= 0 {
		config.Client.Timeout = 30
	}

	var netTransport = &http.Transport{
		TLSHandshakeTimeout: time.Duration(config.Transport.TLSHandshakeTimeout) * time.Second,
	}

	return &http.Client{
		Transport: netTransport,
		Timeout:   time.Duration(config.Client.Timeout) * time.Second,
	}
}
