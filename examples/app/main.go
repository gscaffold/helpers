package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gscaffold/helpers/app"
	"github.com/gscaffold/helpers/logger"
)

func main() {
	// 启动 app 服务
	appI := app.New(
		app.OptionName("example-app"),
		app.OptionWithProfiler(9999),
		app.OptionBeforeStart(func(ctx context.Context) error {
			fmt.Println("example app start...")
			return nil
		}),
		app.OptionAfterStop(func(ctx context.Context) error {
			fmt.Println("example app stop...")
			return nil
		}),
	)
	appI.AddBundle(app.NewDefaultBundle("example-bundle",
		func(ctx context.Context) {
			logger.Info(ctx, "example-bundle dosomething...")
			time.Sleep(time.Second * 10)
		},
		func(ctx context.Context) {
			logger.Info(ctx, "example-bundle stop...")
		}))
	appI.Run(context.TODO())
}
