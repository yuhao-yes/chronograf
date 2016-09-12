package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	context "golang.org/x/net/context"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutDashboardsIDHandlerFunc turns a function with the right signature into a put dashboards ID handler
type PutDashboardsIDHandlerFunc func(context.Context, PutDashboardsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutDashboardsIDHandlerFunc) Handle(ctx context.Context, params PutDashboardsIDParams) middleware.Responder {
	return fn(ctx, params)
}

// PutDashboardsIDHandler interface for that can handle valid put dashboards ID params
type PutDashboardsIDHandler interface {
	Handle(context.Context, PutDashboardsIDParams) middleware.Responder
}

// NewPutDashboardsID creates a new http.Handler for the put dashboards ID operation
func NewPutDashboardsID(ctx *middleware.Context, handler PutDashboardsIDHandler) *PutDashboardsID {
	return &PutDashboardsID{Context: ctx, Handler: handler}
}

/*PutDashboardsID swagger:route PUT /dashboards/{id} putDashboardsId

Replace dashboard configuration.

*/
type PutDashboardsID struct {
	Context *middleware.Context
	Handler PutDashboardsIDHandler
}

func (o *PutDashboardsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutDashboardsIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(context.Background(), Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
