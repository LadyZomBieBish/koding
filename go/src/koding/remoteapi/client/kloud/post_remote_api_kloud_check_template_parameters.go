package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIKloudCheckTemplateParams creates a new PostRemoteAPIKloudCheckTemplateParams object
// with the default values initialized.
func NewPostRemoteAPIKloudCheckTemplateParams() *PostRemoteAPIKloudCheckTemplateParams {
	var ()
	return &PostRemoteAPIKloudCheckTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIKloudCheckTemplateParamsWithTimeout creates a new PostRemoteAPIKloudCheckTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIKloudCheckTemplateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIKloudCheckTemplateParams {
	var ()
	return &PostRemoteAPIKloudCheckTemplateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIKloudCheckTemplateParamsWithContext creates a new PostRemoteAPIKloudCheckTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIKloudCheckTemplateParamsWithContext(ctx context.Context) *PostRemoteAPIKloudCheckTemplateParams {
	var ()
	return &PostRemoteAPIKloudCheckTemplateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIKloudCheckTemplateParams contains all the parameters to send to the API endpoint
for the post remote API kloud check template operation typically these are written to a http.Request
*/
type PostRemoteAPIKloudCheckTemplateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIKloudCheckTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) WithContext(ctx context.Context) *PostRemoteAPIKloudCheckTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) WithBody(body models.DefaultSelector) *PostRemoteAPIKloudCheckTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API kloud check template params
func (o *PostRemoteAPIKloudCheckTemplateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIKloudCheckTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
