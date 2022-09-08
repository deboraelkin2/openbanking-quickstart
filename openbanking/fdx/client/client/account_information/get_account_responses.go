// Code generated by go-swagger; DO NOT EDIT.

package account_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/openbanking/fdx/client/models"
)

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetAccountNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAccountServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/* GetAccountOK describes a response with status code 200, with default header values.

This can be one of LoanAccount, DepositAccount, LocAccount, InvestmentAccount, InsuranceAccount or AnnuityAccount
*/
type GetAccountOK struct {
	Payload interface{}
}

func (o *GetAccountOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}][%d] getAccountOK  %+v", 200, o.Payload)
}
func (o *GetAccountOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNotFound creates a GetAccountNotFound with default headers values
func NewGetAccountNotFound() *GetAccountNotFound {
	return &GetAccountNotFound{}
}

/* GetAccountNotFound describes a response with status code 404, with default header values.

Account with id not found
*/
type GetAccountNotFound struct {
	Payload *models.Error
}

func (o *GetAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}][%d] getAccountNotFound  %+v", 404, o.Payload)
}
func (o *GetAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountInternalServerError creates a GetAccountInternalServerError with default headers values
func NewGetAccountInternalServerError() *GetAccountInternalServerError {
	return &GetAccountInternalServerError{}
}

/* GetAccountInternalServerError describes a response with status code 500, with default header values.

Catch all exception where request was not processed due to an internal outage/issue. Consider other more specific errors before using this error
*/
type GetAccountInternalServerError struct {
	Payload *models.Error
}

func (o *GetAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}][%d] getAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNotImplemented creates a GetAccountNotImplemented with default headers values
func NewGetAccountNotImplemented() *GetAccountNotImplemented {
	return &GetAccountNotImplemented{}
}

/* GetAccountNotImplemented describes a response with status code 501, with default header values.

Error when FdxVersion in Header is not one of those implemented at backend
*/
type GetAccountNotImplemented struct {
	Payload *models.Error
}

func (o *GetAccountNotImplemented) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}][%d] getAccountNotImplemented  %+v", 501, o.Payload)
}
func (o *GetAccountNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountServiceUnavailable creates a GetAccountServiceUnavailable with default headers values
func NewGetAccountServiceUnavailable() *GetAccountServiceUnavailable {
	return &GetAccountServiceUnavailable{}
}

/* GetAccountServiceUnavailable describes a response with status code 503, with default header values.

System is down for maintenance
*/
type GetAccountServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetAccountServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}][%d] getAccountServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetAccountServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
