package devops

import "os"

var (
	sentryDSN = os.Getenv("SENTRY_DSN")
)

func Sentry() string {
	return sentryDSN
}
