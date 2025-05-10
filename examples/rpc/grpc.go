package main

import (
	"context"
	"os"
	"syscall"

	"github.com/gscaffold/helpers/app"
	example_pb "github.com/gscaffold/helpers/examples/rpc/proto"
	"github.com/gscaffold/helpers/logger"
	"github.com/gscaffold/helpers/rpc"
	"github.com/gscaffold/helpers/rpc/grpc"
	"github.com/gscaffold/utils"
)

func main() {
	appIns := app.New()

	// run server
	{
		rpcBundle := rpc.NewGRPCBundle("example-server")
		example_pb.RegisterHelperServiceServer(rpcBundle.Server, &DemoServer{})
		appIns.AddBundle(rpcBundle)
	}

	// run client
	{
		clientFn := func(ctx context.Context) {
			client, err := grpc.NewClient(ctx, "0.0.0.0:8000")
			utils.HandleFatalError(err, "client_init", "")
			svc := example_pb.NewHelperServiceClient(client.GetConn())

			for _, name := range []string{"王勃", "张三"} {
				resp, err := svc.SayHello(ctx, &example_pb.SayHelloReq{Name: name})
				utils.HandleFatalError(err, "client_call", "王勃")
				logger.Infof(context.Background(),
					"client received: name:%s, word:%s", name, resp.Word)
			}

			syscall.Kill(os.Getpid(), syscall.SIGQUIT)
		}
		clientBundle := app.NewDefaultBundle("example-client",
			clientFn,
			func(context.Context) {},
		)
		appIns.AddBundle(clientBundle)
	}

	appIns.Run(context.Background())
}

// ------------ server impl ----------

type DemoServer struct {
	example_pb.UnimplementedHelperServiceServer
}

func (svc *DemoServer) SayHello(ctx context.Context, req *example_pb.SayHelloReq) (*example_pb.SayHelloResp, error) {
	switch req.Name {
	case "王勃":
		return &example_pb.SayHelloResp{
			Word: "落霞与孤鹜齐飞, 秋水共长天一色",
		}, nil
	default:
		return &example_pb.SayHelloResp{
			Word: "Hello " + req.Name,
		}, nil
	}
}
