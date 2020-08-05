// Code generated by go-swagger; DO NOT EDIT.

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle/lcd/models"
)

// GetDistributionValidatorsValidatorAddrRewardsReader is a Reader for the GetDistributionValidatorsValidatorAddrRewards structure.
type GetDistributionValidatorsValidatorAddrRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistributionValidatorsValidatorAddrRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistributionValidatorsValidatorAddrRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDistributionValidatorsValidatorAddrRewardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDistributionValidatorsValidatorAddrRewardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistributionValidatorsValidatorAddrRewardsOK creates a GetDistributionValidatorsValidatorAddrRewardsOK with default headers values
func NewGetDistributionValidatorsValidatorAddrRewardsOK() *GetDistributionValidatorsValidatorAddrRewardsOK {
	return &GetDistributionValidatorsValidatorAddrRewardsOK{}
}

/*GetDistributionValidatorsValidatorAddrRewardsOK handles this case with default header values.

OK
*/
type GetDistributionValidatorsValidatorAddrRewardsOK struct {
	Payload []*models.Coin
}

func (o *GetDistributionValidatorsValidatorAddrRewardsOK) Error() string {
	return fmt.Sprintf("[GET /distribution/validators/{validatorAddr}/rewards][%d] getDistributionValidatorsValidatorAddrRewardsOK  %+v", 200, o.Payload)
}

func (o *GetDistributionValidatorsValidatorAddrRewardsOK) GetPayload() []*models.Coin {
	return o.Payload
}

func (o *GetDistributionValidatorsValidatorAddrRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistributionValidatorsValidatorAddrRewardsBadRequest creates a GetDistributionValidatorsValidatorAddrRewardsBadRequest with default headers values
func NewGetDistributionValidatorsValidatorAddrRewardsBadRequest() *GetDistributionValidatorsValidatorAddrRewardsBadRequest {
	return &GetDistributionValidatorsValidatorAddrRewardsBadRequest{}
}

/*GetDistributionValidatorsValidatorAddrRewardsBadRequest handles this case with default header values.

Invalid validator address
*/
type GetDistributionValidatorsValidatorAddrRewardsBadRequest struct {
}

func (o *GetDistributionValidatorsValidatorAddrRewardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /distribution/validators/{validatorAddr}/rewards][%d] getDistributionValidatorsValidatorAddrRewardsBadRequest ", 400)
}

func (o *GetDistributionValidatorsValidatorAddrRewardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDistributionValidatorsValidatorAddrRewardsInternalServerError creates a GetDistributionValidatorsValidatorAddrRewardsInternalServerError with default headers values
func NewGetDistributionValidatorsValidatorAddrRewardsInternalServerError() *GetDistributionValidatorsValidatorAddrRewardsInternalServerError {
	return &GetDistributionValidatorsValidatorAddrRewardsInternalServerError{}
}

/*GetDistributionValidatorsValidatorAddrRewardsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDistributionValidatorsValidatorAddrRewardsInternalServerError struct {
}

func (o *GetDistributionValidatorsValidatorAddrRewardsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/validators/{validatorAddr}/rewards][%d] getDistributionValidatorsValidatorAddrRewardsInternalServerError ", 500)
}

func (o *GetDistributionValidatorsValidatorAddrRewardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}