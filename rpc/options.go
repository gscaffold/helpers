package rpc

import (
	hgrpc "github.com/gscaffold/helpers/rpc/grpc"
	"google.golang.org/grpc"
)

type GRPCOption func(bundle *GRPCBundle)

func GRPCOptionListen(listenAddr string) GRPCOption {
	return func(s *GRPCBundle) {
		s.listenAddr = listenAddr
	}
}

func GRPCOptionOrigin(opts ...grpc.ServerOption) GRPCOption {
	return func(s *GRPCBundle) {
		s.serverOptions = append(s.serverOptions, opts...)
	}
}

func GRPCOptionMetrics(prefix string) GRPCOption {
	return func(s *GRPCBundle) {
		s.serverOptions = append(s.serverOptions,
			grpc.ChainUnaryInterceptor(hgrpc.ServerMetricsInterceptor(prefix)))
	}
}
