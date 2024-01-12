package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(Context) error

type Router struct {
	srv *Server
}

func newRouter(srv *Server) *Router {
	r := &Router{
		srv: srv,
	}

	return r
}

func (r *Router) Handle(method, relativePath string, h HandlerFunc) {
	r.srv.engine.Handle(method, relativePath, func(c *gin.Context) {
		ctx := &wrapper{router: r}
		if err := h(ctx); err != nil {
			r.srv.ene(ctx, err)
		} else {
			r.srv.enc(ctx, nil)
		}
		c.Next()
	})
}

func (r *Router) HEAD(path string, h HandlerFunc) {
	r.Handle(http.MethodHead, path, h)
}

func (r *Router) GET(path string, h HandlerFunc) {
	r.Handle(http.MethodGet, path, h)
}

func (r *Router) POST(path string, h HandlerFunc) {
	r.Handle(http.MethodPost, path, h)
}

func (r *Router) PUT(path string, h HandlerFunc) {
	r.Handle(http.MethodPut, path, h)
}

func (r *Router) PATCH(path string, h HandlerFunc) {
	r.Handle(http.MethodPatch, path, h)
}

func (r *Router) DELETE(path string, h HandlerFunc) {
	r.Handle(http.MethodDelete, path, h)
}

func (r *Router) OPTIONS(path string, h HandlerFunc) {
	r.Handle(http.MethodOptions, path, h)
}
