// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetGovProposalsProposalIDDepositsDepositorReader is a Reader for the GetGovProposalsProposalIDDepositsDepositor structure.
type GetGovProposalsProposalIDDepositsDepositorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDDepositsDepositorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDDepositsDepositorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDDepositsDepositorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGovProposalsProposalIDDepositsDepositorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDDepositsDepositorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDDepositsDepositorOK creates a GetGovProposalsProposalIDDepositsDepositorOK with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorOK() *GetGovProposalsProposalIDDepositsDepositorOK {
	return &GetGovProposalsProposalIDDepositsDepositorOK{}
}

/*GetGovProposalsProposalIDDepositsDepositorOK handles this case with default header values.

OK
*/
type GetGovProposalsProposalIDDepositsDepositorOK struct {
	Payload *models.Deposit
}

func (o *GetGovProposalsProposalIDDepositsDepositorOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorOK  %+v", 200, o.Payload)
}

func (o *GetGovProposalsProposalIDDepositsDepositorOK) GetPayload() *models.Deposit {
	return o.Payload
}

func (o *GetGovProposalsProposalIDDepositsDepositorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deposit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorBadRequest creates a GetGovProposalsProposalIDDepositsDepositorBadRequest with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorBadRequest() *GetGovProposalsProposalIDDepositsDepositorBadRequest {
	return &GetGovProposalsProposalIDDepositsDepositorBadRequest{}
}

/*GetGovProposalsProposalIDDepositsDepositorBadRequest handles this case with default header values.

Invalid proposal id or depositor address
*/
type GetGovProposalsProposalIDDepositsDepositorBadRequest struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDDepositsDepositorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorNotFound creates a GetGovProposalsProposalIDDepositsDepositorNotFound with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorNotFound() *GetGovProposalsProposalIDDepositsDepositorNotFound {
	return &GetGovProposalsProposalIDDepositsDepositorNotFound{}
}

/*GetGovProposalsProposalIDDepositsDepositorNotFound handles this case with default header values.

Found no deposit
*/
type GetGovProposalsProposalIDDepositsDepositorNotFound struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorNotFound) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorNotFound ", 404)
}

func (o *GetGovProposalsProposalIDDepositsDepositorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorInternalServerError creates a GetGovProposalsProposalIDDepositsDepositorInternalServerError with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorInternalServerError() *GetGovProposalsProposalIDDepositsDepositorInternalServerError {
	return &GetGovProposalsProposalIDDepositsDepositorInternalServerError{}
}

/*GetGovProposalsProposalIDDepositsDepositorInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDDepositsDepositorInternalServerError struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDDepositsDepositorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}