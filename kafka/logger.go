package kafka

import (
	"context"

	"github.com/gscaffold/helpers/logger"
)

func Infof(msg string, args ...interface{}) {
	logger.Infof(context.TODO(), msg, args...)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(context.TODO(), msg, args...)
}
