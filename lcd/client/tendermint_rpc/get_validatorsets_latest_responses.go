// Code generated by go-swagger; DO NOT EDIT.

package tendermint_rpc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// GetValidatorsetsLatestReader is a Reader for the GetValidatorsetsLatest structure.
type GetValidatorsetsLatestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidatorsetsLatestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetValidatorsetsLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetValidatorsetsLatestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetValidatorsetsLatestOK creates a GetValidatorsetsLatestOK with default headers values
func NewGetValidatorsetsLatestOK() *GetValidatorsetsLatestOK {
	return &GetValidatorsetsLatestOK{}
}

/*GetValidatorsetsLatestOK handles this case with default header values.

The validator set at the latest block height
*/
type GetValidatorsetsLatestOK struct {
	Payload *GetValidatorsetsLatestOKBody
}

func (o *GetValidatorsetsLatestOK) Error() string {
	return fmt.Sprintf("[GET /validatorsets/latest][%d] getValidatorsetsLatestOK  %+v", 200, o.Payload)
}

func (o *GetValidatorsetsLatestOK) GetPayload() *GetValidatorsetsLatestOKBody {
	return o.Payload
}

func (o *GetValidatorsetsLatestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetValidatorsetsLatestOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidatorsetsLatestInternalServerError creates a GetValidatorsetsLatestInternalServerError with default headers values
func NewGetValidatorsetsLatestInternalServerError() *GetValidatorsetsLatestInternalServerError {
	return &GetValidatorsetsLatestInternalServerError{}
}

/*GetValidatorsetsLatestInternalServerError handles this case with default header values.

Server internal error
*/
type GetValidatorsetsLatestInternalServerError struct {
}

func (o *GetValidatorsetsLatestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /validatorsets/latest][%d] getValidatorsetsLatestInternalServerError ", 500)
}

func (o *GetValidatorsetsLatestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetValidatorsetsLatestOKBody get validatorsets latest o k body
swagger:model GetValidatorsetsLatestOKBody
*/
type GetValidatorsetsLatestOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetValidatorsetsLatestOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get validatorsets latest o k body
func (o *GetValidatorsetsLatestOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetValidatorsetsLatestOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getValidatorsetsLatestOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetValidatorsetsLatestOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetValidatorsetsLatestOKBody) UnmarshalBinary(b []byte) error {
	var res GetValidatorsetsLatestOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetValidatorsetsLatestOKBodyResult get validatorsets latest o k body result
swagger:model GetValidatorsetsLatestOKBodyResult
*/
type GetValidatorsetsLatestOKBodyResult struct {

	// block height
	BlockHeight string `json:"block_height,omitempty"`

	// validators
	Validators []*models.TendermintValidator `json:"validators"`
}

// Validate validates this get validatorsets latest o k body result
func (o *GetValidatorsetsLatestOKBodyResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetValidatorsetsLatestOKBodyResult) validateValidators(formats strfmt.Registry) error {

	if swag.IsZero(o.Validators) { // not required
		return nil
	}

	for i := 0; i < len(o.Validators); i++ {
		if swag.IsZero(o.Validators[i]) { // not required
			continue
		}

		if o.Validators[i] != nil {
			if err := o.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getValidatorsetsLatestOK" + "." + "result" + "." + "validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetValidatorsetsLatestOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetValidatorsetsLatestOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetValidatorsetsLatestOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
