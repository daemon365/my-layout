package http

import (
	"net"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHttpServer(listener net.Listener, log log.Logger) *http.Server {
	return http.NewServer(
		http.Listener(listener),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(log),
		),
	)
}
