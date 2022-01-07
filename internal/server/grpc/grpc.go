package grpc

import (
	"net"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGrpcServer(listener net.Listener, log log.Logger) *grpc.Server {
	return grpc.NewServer(
		grpc.Listener(listener),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log),
		),
	)
}
