package gin

import (
	"context"
	"fmt"
	"time"
)

var _ Context = (*wrapper)(nil)

type Context interface {
	context.Context
	Bind(interface{}) error
	Reply(int, interface{}) error
	File(filepath string) error
}

type wrapper struct {
	router *Router
}

// Bind implements Context.
func (*wrapper) Bind(interface{}) error {
	panic("unimplemented")
}

// Reply implements Context.
func (c *wrapper) Reply(code int, v interface{}) error {
	tolog := fmt.Sprintf("reply code: %d", code)
	fmt.Println(tolog)
	return c.router.srv.enc(c, v)
}

// File implements Context.
func (*wrapper) File(filepath string) error {
	panic("unimplemented")
}

// Deadline implements Context.
func (*wrapper) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements Context.
func (*wrapper) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements Context.
func (*wrapper) Err() error {
	panic("unimplemented")
}

// Value implements Context.
func (*wrapper) Value(key any) any {
	panic("unimplemented")
}
