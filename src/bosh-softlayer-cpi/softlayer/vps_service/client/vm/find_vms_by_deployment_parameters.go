package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindVmsByDeploymentParams creates a new FindVmsByDeploymentParams object
// with the default values initialized.
func NewFindVmsByDeploymentParams() *FindVmsByDeploymentParams {
	var ()
	return &FindVmsByDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindVmsByDeploymentParamsWithTimeout creates a new FindVmsByDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindVmsByDeploymentParamsWithTimeout(timeout time.Duration) *FindVmsByDeploymentParams {
	var ()
	return &FindVmsByDeploymentParams{

		timeout: timeout,
	}
}

// NewFindVmsByDeploymentParamsWithContext creates a new FindVmsByDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindVmsByDeploymentParamsWithContext(ctx context.Context) *FindVmsByDeploymentParams {
	var ()
	return &FindVmsByDeploymentParams{

		Context: ctx,
	}
}

/*FindVmsByDeploymentParams contains all the parameters to send to the API endpoint
for the find vms by deployment operation typically these are written to a http.Request
*/
type FindVmsByDeploymentParams struct {

	/*Deployment
	  deployment values that need to be considered for filter

	*/
	Deployment []string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the find vms by deployment params
func (o *FindVmsByDeploymentParams) WithTimeout(timeout time.Duration) *FindVmsByDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find vms by deployment params
func (o *FindVmsByDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find vms by deployment params
func (o *FindVmsByDeploymentParams) WithContext(ctx context.Context) *FindVmsByDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find vms by deployment params
func (o *FindVmsByDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDeployment adds the deployment to the find vms by deployment params
func (o *FindVmsByDeploymentParams) WithDeployment(deployment []string) *FindVmsByDeploymentParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the find vms by deployment params
func (o *FindVmsByDeploymentParams) SetDeployment(deployment []string) {
	o.Deployment = deployment
}

// WriteToRequest writes these params to a swagger request
func (o *FindVmsByDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDeployment := o.Deployment

	joinedDeployment := swag.JoinByFormat(valuesDeployment, "multi")
	// query array param deployment
	if err := r.SetQueryParam("deployment", joinedDeployment...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
