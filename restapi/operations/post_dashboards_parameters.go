package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/influxdata/mrfusion/models"
)

// NewPostDashboardsParams creates a new PostDashboardsParams object
// with the default values initialized.
func NewPostDashboardsParams() PostDashboardsParams {
	var ()
	return PostDashboardsParams{}
}

// PostDashboardsParams contains all the bound params for the post dashboards operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostDashboards
type PostDashboardsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Defines the dashboard and queries of the cells within the dashboard.
	  In: body
	*/
	Dashboard *models.Dashboard
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostDashboardsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Dashboard
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("dashboard", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Dashboard = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
