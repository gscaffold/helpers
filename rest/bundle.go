package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gscaffold/helpers/logger"
	"github.com/gscaffold/utils"
)

type HTTPBundle struct {
	name       string
	router     http.Handler
	httpServer http.Server

	timeout      time.Duration // 针对 client 端的限制，到时间返回
	writeTimeout time.Duration // httpServer.writeTimeout
	readTimeout  time.Duration // httpServer.readTimeout

	port int
}

func (bundle *HTTPBundle) LoadDefault() {
	if bundle.name == "" {
		bundle.name = utils.GetApp()
	}
	if bundle.port == 0 {
		bundle.port = 8080
	}
}

func New(router http.Handler, opts ...Option) *HTTPBundle {
	if router == nil {
		err := errors.New("must have router")
		utils.HandleFatalError(err, "new_http_bundle", "")
	}

	api := &HTTPBundle{
		router: router,
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func (bundle *HTTPBundle) GetName() string {
	return bundle.name
}

func (bundle *HTTPBundle) Run(ctx context.Context) {
	bundle.LoadDefault()

	bundle.httpServer = http.Server{
		Addr:         ":" + strconv.Itoa(bundle.port),
		Handler:      http.TimeoutHandler(bundle.router, bundle.timeout, ""),
		ReadTimeout:  bundle.readTimeout,
		WriteTimeout: bundle.writeTimeout,
	}
	// 允许取消TimeoutHandler以支持ws等协议
	if bundle.timeout == 0 {
		bundle.httpServer.Handler = bundle.router
	}

	logger.Infof(ctx, "HTTP Server is listening on %s", bundle.httpServer.Addr)
	if err := bundle.httpServer.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logger.Errorf(ctx, "HTTP Server start error. err:%s", err)
			panic(err)
		}
	}
}

func (bundle *HTTPBundle) Stop(ctx context.Context) {
	logger.Info(ctx, "HTTP Server is shutdown")
	if err := bundle.httpServer.Shutdown(ctx); err != nil {
		logger.Errorf(ctx, "HTTP Server shutdown error. err:%s", err)
	}
}
