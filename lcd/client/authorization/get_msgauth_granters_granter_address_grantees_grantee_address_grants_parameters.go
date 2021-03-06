// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams creates a new GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams object
// with the default values initialized.
func NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams() *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	var ()
	return &GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithTimeout creates a new GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithTimeout(timeout time.Duration) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	var ()
	return &GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams{

		timeout: timeout,
	}
}

// NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithContext creates a new GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithContext(ctx context.Context) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	var ()
	return &GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams{

		Context: ctx,
	}
}

// NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithHTTPClient creates a new GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParamsWithHTTPClient(client *http.Client) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	var ()
	return &GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams{
		HTTPClient: client,
	}
}

/*GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams contains all the parameters to send to the API endpoint
for the get msgauth granters granter address grantees grantee address grants operation typically these are written to a http.Request
*/
type GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams struct {

	/*GranteeAddress
	  grantee address you want to lookup

	*/
	GranteeAddress string
	/*GranterAddress
	  granter address you want to lookup

	*/
	GranterAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WithTimeout(timeout time.Duration) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WithContext(ctx context.Context) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WithHTTPClient(client *http.Client) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGranteeAddress adds the granteeAddress to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WithGranteeAddress(granteeAddress string) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	o.SetGranteeAddress(granteeAddress)
	return o
}

// SetGranteeAddress adds the granteeAddress to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) SetGranteeAddress(granteeAddress string) {
	o.GranteeAddress = granteeAddress
}

// WithGranterAddress adds the granterAddress to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WithGranterAddress(granterAddress string) *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams {
	o.SetGranterAddress(granterAddress)
	return o
}

// SetGranterAddress adds the granterAddress to the get msgauth granters granter address grantees grantee address grants params
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) SetGranterAddress(granterAddress string) {
	o.GranterAddress = granterAddress
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param granteeAddress
	if err := r.SetPathParam("granteeAddress", o.GranteeAddress); err != nil {
		return err
	}

	// path param granterAddress
	if err := r.SetPathParam("granterAddress", o.GranterAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
