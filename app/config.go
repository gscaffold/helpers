package app

import "time"

const DefaultStopTimeout = time.Minute * 5

type appConfig struct {
	stopTimeout time.Duration // app stop timeoit
	profilePort int           // go profile port
}
