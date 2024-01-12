package gin

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerOption func(*Server)

func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func Address(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func Endpoint(endpoint *url.URL) ServerOption {
	return func(s *Server) {
		s.endpoint = endpoint
	}
}

func TLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

func Engine(r *gin.Engine) ServerOption {
	return func(o *Server) {
		o.engine = r
	}
}

type Server struct {
	*http.Server
	err      error
	network  string
	address  string
	timeout  time.Duration
	endpoint *url.URL
	tlsConf  *tls.Config
	engine   *gin.Engine
	enc      EncodeResponse
	ene      EncodeError
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		engine:  gin.New(),
		enc:     DefaultEncodeResponse,
		ene:     DefaultEncodeError,
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) Router() *Router {
	return newRouter(s)
}
