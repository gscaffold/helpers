package rest

import (
	"time"
)

type Option func(bundle *HTTPBundle)

func Port(port int) Option {
	return func(bundle *HTTPBundle) {
		bundle.port = port
	}
}

func Timeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.timeout = t
	}
}

func ReadTimeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.readTimeout = t
	}
}

func WriteTimeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.writeTimeout = t
	}
}
