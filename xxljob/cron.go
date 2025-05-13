package xxljob

import (
	"context"
	"errors"
	"fmt"

	"github.com/gscaffold/helpers/logger"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type Cron struct {
	xxl.Executor
}

func New(_opts ...Option) (*Cron, error) {
	opts := &options{}
	for _, opt := range _opts {
		opt(opts)
	}
	opts.LoadDefault()

	// validate
	{
		if opts.serverAddr == "" {
			return &Cron{}, errors.New("must have server_addr")
		}
		if opts.registryKey == "" {
			return &Cron{}, errors.New("must have registry_key")
		}
	}

	client := xxl.NewExecutor(
		xxl.ServerAddr(opts.serverAddr),
		xxl.AccessToken(opts.accessToken),
		xxl.ExecutorIp(opts.executorIP),
		xxl.ExecutorPort(fmt.Sprint(opts.port)),
		xxl.RegistryKey(opts.registryKey),
		xxl.SetLogger(opts.logger),
	)
	client.Init()
	client.Use(LogMiddleware, RecoverMiddleware)
	logger.Infof(context.TODO(), "create new xxljob executor. name:%s, opts:%+v", opts.registryKey, opts)
	return &Cron{
		Executor: client,
	}, nil
}

func LogMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(ctx context.Context, param *xxl.RunReq) string {
		logger.Infof(ctx, "task start %s, prams:%+v", param.ExecutorHandler, param)
		res := tf(ctx, param)
		logger.Infof(ctx, "task finished %s, result:%s, prams:%+v", param.ExecutorHandler, res, param)
		return res
	}
}

func RecoverMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(ctx context.Context, param *xxl.RunReq) string {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf(ctx, "task %s panic. param:%+v", param.ExecutorHandler, param)
			}
		}()
		res := tf(ctx, param)
		return res
	}
}
