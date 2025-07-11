package metrics

import (
	"context"
	"sync"
	"time"

	"github.com/gscaffold/helpers/logger"
	"github.com/smira/go-statsd"
)

var (
	defaultClient *statsd.Client
	once          sync.Once
)

type Tag = statsd.Tag

func getClient() *statsd.Client {
	if defaultClient == nil {
		once.Do(func() {
			// todo 测试
			// defaultClient = statsd.NewClient()("127.0.0.1:8125")
		})
	}
	return defaultClient
}

func Close() error {
	if defaultClient == nil {
		return nil
	}
	return getClient().Close()
}

func Incr(name string, tags ...Tag) {
	getClient().Incr(name, 1, tags...)
}

func Count(name string, value int64, tags ...Tag) {
	getClient().Incr(name, value, tags...)
}

func FCount(name string, value float64, tags ...Tag) {
	getClient().FIncr(name, value, tags...)
}

func Guage(name string, value int64, tags ...Tag) {
	getClient().Gauge(name, value, tags...)
}

func FGuage(name string, value float64, tags ...Tag) {
	getClient().FGaugeDelta(name, value, tags...)
}

func Timing(name string, value time.Duration, tags ...Tag) {
	getClient().Timing(name, value.Milliseconds(), tags...)
}

// TimingSince 耗时计算语法糖, defer TimingSince 直接计算耗时.
func TimingSince(name string, now time.Time, tags ...Tag) {
	Timing(name, time.Since(now), tags...)
}

// 语法糖, 输入日志并且打点
func LoggerErrorAndMetrics(ctx context.Context, metric, format string, args ...interface{}) {
	logger.Errorf(ctx, format, args...)
	Incr(metric)
}
