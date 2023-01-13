// Code generated by go-swagger; DO NOT EDIT.

package reward_program_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// GetRewardProgramReader is a Reader for the GetRewardProgram structure.
type GetRewardProgramReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRewardProgramReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRewardProgramOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRewardProgramNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRewardProgramInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetRewardProgramNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRewardProgramServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRewardProgramOK creates a GetRewardProgramOK with default headers values
func NewGetRewardProgramOK() *GetRewardProgramOK {
	return &GetRewardProgramOK{}
}

/* GetRewardProgramOK describes a response with status code 200, with default header values.

Data describing reward programs associated with accounts
*/
type GetRewardProgramOK struct {
	Payload *models.RewardProgramentity
}

func (o *GetRewardProgramOK) Error() string {
	return fmt.Sprintf("[GET /reward-programs/{rewardProgramId}][%d] getRewardProgramOK  %+v", 200, o.Payload)
}
func (o *GetRewardProgramOK) GetPayload() *models.RewardProgramentity {
	return o.Payload
}

func (o *GetRewardProgramOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RewardProgramentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRewardProgramNotFound creates a GetRewardProgramNotFound with default headers values
func NewGetRewardProgramNotFound() *GetRewardProgramNotFound {
	return &GetRewardProgramNotFound{}
}

/* GetRewardProgramNotFound describes a response with status code 404, with default header values.

Reward program Id not found
*/
type GetRewardProgramNotFound struct {
	Payload *models.Error
}

func (o *GetRewardProgramNotFound) Error() string {
	return fmt.Sprintf("[GET /reward-programs/{rewardProgramId}][%d] getRewardProgramNotFound  %+v", 404, o.Payload)
}
func (o *GetRewardProgramNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRewardProgramNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRewardProgramInternalServerError creates a GetRewardProgramInternalServerError with default headers values
func NewGetRewardProgramInternalServerError() *GetRewardProgramInternalServerError {
	return &GetRewardProgramInternalServerError{}
}

/* GetRewardProgramInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type GetRewardProgramInternalServerError struct {
	Payload *models.Error
}

func (o *GetRewardProgramInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reward-programs/{rewardProgramId}][%d] getRewardProgramInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRewardProgramInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRewardProgramInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRewardProgramNotImplemented creates a GetRewardProgramNotImplemented with default headers values
func NewGetRewardProgramNotImplemented() *GetRewardProgramNotImplemented {
	return &GetRewardProgramNotImplemented{}
}

/* GetRewardProgramNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type GetRewardProgramNotImplemented struct {
	Payload *models.Error
}

func (o *GetRewardProgramNotImplemented) Error() string {
	return fmt.Sprintf("[GET /reward-programs/{rewardProgramId}][%d] getRewardProgramNotImplemented  %+v", 501, o.Payload)
}
func (o *GetRewardProgramNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRewardProgramNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRewardProgramServiceUnavailable creates a GetRewardProgramServiceUnavailable with default headers values
func NewGetRewardProgramServiceUnavailable() *GetRewardProgramServiceUnavailable {
	return &GetRewardProgramServiceUnavailable{}
}

/* GetRewardProgramServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type GetRewardProgramServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetRewardProgramServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reward-programs/{rewardProgramId}][%d] getRewardProgramServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetRewardProgramServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRewardProgramServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
