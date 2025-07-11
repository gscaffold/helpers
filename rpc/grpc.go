package rpc

import (
	"context"
	"net"

	application "github.com/gscaffold/helpers/app"
	hgrpc "github.com/gscaffold/helpers/rpc/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GRPCBundle struct {
	name       string
	listenAddr string

	Server        *grpc.Server
	serverOptions []grpc.ServerOption
}

var _ application.IBundle = new(GRPCBundle)

func NewGRPCBundle(name string, opts ...GRPCOption) *GRPCBundle {
	s := &GRPCBundle{
		name:       name,
		listenAddr: "0.0.0.0:8000",
	}
	for _, opt := range opts {
		opt(s)
	}

	s.Server = hgrpc.NewServer(s.serverOptions...)

	return s
}

func (s *GRPCBundle) GetName() string {
	return s.name
}

func (s *GRPCBundle) Run(ctx context.Context) {
	addr, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		err := errors.Wrap(err, "listen failed")
		panic(err)
	}
	err = s.Server.Serve(addr)
	if err != nil {
		panic(err)
	}
}

func (s *GRPCBundle) Stop(ctx context.Context) {
	s.Server.Stop()
}
