// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CheckHealthHandlerFunc turns a function with the right signature into a check health handler
type CheckHealthHandlerFunc func(CheckHealthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CheckHealthHandlerFunc) Handle(params CheckHealthParams) middleware.Responder {
	return fn(params)
}

// CheckHealthHandler interface for that can handle valid check health params
type CheckHealthHandler interface {
	Handle(CheckHealthParams) middleware.Responder
}

// NewCheckHealth creates a new http.Handler for the check health operation
func NewCheckHealth(ctx *middleware.Context, handler CheckHealthHandler) *CheckHealth {
	return &CheckHealth{Context: ctx, Handler: handler}
}

/*CheckHealth swagger:route GET /healthz checkHealth

CheckHealth check health API

*/
type CheckHealth struct {
	Context *middleware.Context
	Handler CheckHealthHandler
}

func (o *CheckHealth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCheckHealthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
