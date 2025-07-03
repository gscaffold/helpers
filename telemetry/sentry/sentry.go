/*
sentry 对程序中的错误事件进行聚合分析.

用法:
1. 程序启动时, 调用 Init 初始化 sentry,
2. 程序退出时, 调用 Close 将数据刷盘. 建议在启动时将 close 函数添加到析构方法中.
*/
package sentry

/*
SampleRate 采样
BeforeSend 过滤或定制要发送的事件
sentry.Recover() recover + 上报事件
sentry.Flush()
设置 level, 并且和 logger 交互

资源发现(环境变量)

add logger hook
recover hook: panic 时上报 sentry, 避免程序崩溃, 作为 logger、metrics 的补充.
一般服务场景不存在panic需要重启程序的场景, 如有避免引入该组件.
	add gin hook
	add grpc hook

*/

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gscaffold/helpers/devops"
	"github.com/gscaffold/utils"
)

func Init(opts ...Option) {
	clientOpts := sentry.ClientOptions{
		Dsn:         devops.Sentry(),
		ServerName:  utils.GetApp(),
		Release:     utils.AppVersion, // 默认从构建环境中取
		Environment: utils.GetEnv(),
	}
	for _, opt := range opts {
		opt(&clientOpts)
	}
	err := sentry.Init(clientOpts)
	utils.HandleFatalError(err, "sentry", "sentry init error")
}

func Close() {
	sentry.Flush(5 * time.Second)
}

// --------------------------
type Option func(*sentry.ClientOptions)

func OptionDSN(dsn string) Option {
	return func(opts *sentry.ClientOptions) {
		opts.Dsn = dsn
	}
}

func OptionDebug() Option {
	return func(opts *sentry.ClientOptions) {
		opts.Debug = true
	}
}

func OptionSampleRate(rate float64) Option {
	return func(opts *sentry.ClientOptions) {
		// 0 在 sentry 中表示全采, 只能通过取消dsn关闭采样
		if rate == 0 {
			opts.Dsn = ""
		} else {
			opts.SampleRate = rate
		}
	}
}

func OptionIgnoreErrors(errs []string) Option {
	return func(opts *sentry.ClientOptions) {
		opts.IgnoreErrors = errs
	}
}

func OptionBeforeSend(fn func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event) Option {
	return func(opts *sentry.ClientOptions) {
		opts.BeforeSend = fn
	}
}
