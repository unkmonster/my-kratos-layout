package server

import (
	"net/http"

	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"
	"github.com/gorilla/handlers"
	"github.com/unkmonster/go-kit/middleware/http/realip"
	"github.com/unkmonster/go-kit/middleware/http/requestid"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func newCorsHandler(c *conf.Server_HTTP_CORS) func(http.Handler) http.Handler {
	opts := []handlers.CORSOption{
		handlers.AllowedHeaders(c.AllowedHeaders),
		handlers.AllowedMethods(c.AllowedMethods),
		handlers.AllowedOrigins(c.AllowedOrigins),
	}

	if c.AllowCredentials {
		opts = append(opts, handlers.AllowCredentials())
	}

	return handlers.CORS(
		opts...,
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			requestid.Server(),
			realip.Server(
				logger,
				realip.WithIpHeaders(c.Http.RealIp.IpHeaders),
				realip.WithTrustedHeader(c.Http.RealIp.TrustedHeader),
				realip.WithTrustedProxies(c.Http.RealIp.TrustedProxies),
			),
			metadata.Server(),
			logging.Server(logger),
			validate.ProtoValidate(),
		),
		khttp.Filter(
			newCorsHandler(c.Http.Cors),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, khttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, khttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, khttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := khttp.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
