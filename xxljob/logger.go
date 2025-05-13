package xxljob

import (
	"context"

	"github.com/gscaffold/helpers/logger"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type Logger struct{}

var _ xxl.Logger = new(Logger)

func (l Logger) Info(format string, args ...interface{}) {
	logger.Infof(context.TODO(), format, args...)
}

func (l Logger) Error(format string, args ...interface{}) {
	logger.Errorf(context.TODO(), format, args...)
}
