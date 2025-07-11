package hgrpc

import (
	"context"
	"strings"
	"time"

	"github.com/gscaffold/helpers/telemetry/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	overwriteOpts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     5 * time.Minute,
			MaxConnectionAge:      5 * time.Minute,
			MaxConnectionAgeGrace: 10 * time.Second,
			Time:                  time.Second,
			Timeout:               time.Millisecond * 100,
		}),
	}
	opts = append(opts, overwriteOpts...)

	s := grpc.NewServer(opts...)
	// 允许客户端查询服务端支持的服务和方法
	reflection.Register(s)
	// k8s 检活服务
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	return s
}

func ServerMetricsInterceptor(prefix string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		path := strings.ReplaceAll(info.FullMethod, "/", ".")
		metric := prefix + path
		defer metrics.TimingSince(metric, time.Now())

		resp, err = handler(ctx, req)
		if err == nil {
			metric += ".success"
		} else {
			metric += ".error"
		}

		metrics.Incr(metric)
		return resp, err
	}
}
