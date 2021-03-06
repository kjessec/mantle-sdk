// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeReader is a Reader for the PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgType structure.
type PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK creates a PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK with default headers values
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK() *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK {
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK{}
}

/*PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK handles this case with default header values.

OK
*/
type PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK struct {
	Payload *models.StdTx
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK) Error() string {
	return fmt.Sprintf("[POST /msgauth/granters/{granterAddress}/grantees/{granteeAddress}/grants/{msgType}][%d] postMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK  %+v", 200, o.Payload)
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest creates a PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest with default headers values
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest() *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest {
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest{}
}

/*PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest handles this case with default header values.

Bad request
*/
type PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest struct {
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest) Error() string {
	return fmt.Sprintf("[POST /msgauth/granters/{granterAddress}/grantees/{granteeAddress}/grants/{msgType}][%d] postMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest ", 400)
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError creates a PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError with default headers values
func NewPostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError() *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError {
	return &PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError{}
}

/*PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError struct {
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /msgauth/granters/{granterAddress}/grantees/{granteeAddress}/grants/{msgType}][%d] postMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError ", 500)
}

func (o *PostMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsMsgTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
