package server

import (
	"net"
	"strings"

	"gitee.com/haiyux/kratos-layout/api/testing"
	"gitee.com/haiyux/kratos-layout/internal/controller"
	"gitee.com/haiyux/kratos-layout/internal/server/grpc"
	"gitee.com/haiyux/kratos-layout/internal/server/http"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/soheilhy/cmux"
)

func InitServer(addr string, log log.Logger) (*kratos.App, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	m := cmux.New(l)
	grpcSrv := grpc.NewGrpcServer(m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc")), log)
	httpSrv := http.NewHttpServer(m.Match(cmux.Any()), log)

	// register grpc and http server
	testing.RegisterTestingServer(grpcSrv, controller.NewTestingService())
	testing.RegisterTestingHTTPServer(httpSrv, controller.NewTestingService())

	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
		// kratos.Registrar(r),
	)

	go func() {
		if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
			panic(err)
		}
	}()

	return app, nil
}
