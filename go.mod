module github.com/gscaffold/helpers

go 1.23

toolchain go1.23.9

replace github.com/gscaffold/utils => ../utils

require (
	github.com/gscaffold/utils v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/stretchr/testify v1.9.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)
