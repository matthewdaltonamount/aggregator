package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

type ControllerChain struct {
	Err     error
	Writer  http.ResponseWriter
	Request *http.Request
}

func (cc *ControllerChain) E(code int, f func() error) *ControllerChain {
	if cc.Err == nil {
		err := f()
		cc.Err = err
		if err != nil {
			_ = render.Render(cc.Writer, cc.Request, NewErrorResponse(err, code))
		}
	}
	return cc
}
