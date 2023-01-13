// Code generated by go-swagger; DO NOT EDIT.

package internal_transfers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// GetTransferReader is a Reader for the GetTransfer structure.
type GetTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransferOK creates a GetTransferOK with default headers values
func NewGetTransferOK() *GetTransferOK {
	return &GetTransferOK{}
}

/* GetTransferOK describes a response with status code 200, with default header values.

Ok
*/
type GetTransferOK struct {
	Payload *models.Transferentity
}

func (o *GetTransferOK) Error() string {
	return fmt.Sprintf("[GET /transfers/{transferId}][%d] getTransferOK  %+v", 200, o.Payload)
}
func (o *GetTransferOK) GetPayload() *models.Transferentity {
	return o.Payload
}

func (o *GetTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Transferentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
