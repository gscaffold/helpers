package rest

import (
	"time"
)

type Option func(bundle *HTTPBundle)

func OptionPort(port int) Option {
	return func(bundle *HTTPBundle) {
		bundle.port = port
	}
}

func OptionTimeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.timeout = t
	}
}

func OptionReadTimeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.readTimeout = t
	}
}

func OptionWriteTimeout(t time.Duration) Option {
	return func(bundle *HTTPBundle) {
		bundle.writeTimeout = t
	}
}
