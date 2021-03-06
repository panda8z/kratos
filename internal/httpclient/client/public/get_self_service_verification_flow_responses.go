// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// GetSelfServiceVerificationFlowReader is a Reader for the GetSelfServiceVerificationFlow structure.
type GetSelfServiceVerificationFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceVerificationFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceVerificationFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceVerificationFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceVerificationFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceVerificationFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSelfServiceVerificationFlowOK creates a GetSelfServiceVerificationFlowOK with default headers values
func NewGetSelfServiceVerificationFlowOK() *GetSelfServiceVerificationFlowOK {
	return &GetSelfServiceVerificationFlowOK{}
}

/* GetSelfServiceVerificationFlowOK describes a response with status code 200, with default header values.

verificationFlow
*/
type GetSelfServiceVerificationFlowOK struct {
	Payload *models.VerificationFlow
}

func (o *GetSelfServiceVerificationFlowOK) Error() string {
	return fmt.Sprintf("[GET /self-service/verification/flows][%d] getSelfServiceVerificationFlowOK  %+v", 200, o.Payload)
}
func (o *GetSelfServiceVerificationFlowOK) GetPayload() *models.VerificationFlow {
	return o.Payload
}

func (o *GetSelfServiceVerificationFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VerificationFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceVerificationFlowForbidden creates a GetSelfServiceVerificationFlowForbidden with default headers values
func NewGetSelfServiceVerificationFlowForbidden() *GetSelfServiceVerificationFlowForbidden {
	return &GetSelfServiceVerificationFlowForbidden{}
}

/* GetSelfServiceVerificationFlowForbidden describes a response with status code 403, with default header values.

genericError
*/
type GetSelfServiceVerificationFlowForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceVerificationFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/verification/flows][%d] getSelfServiceVerificationFlowForbidden  %+v", 403, o.Payload)
}
func (o *GetSelfServiceVerificationFlowForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceVerificationFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceVerificationFlowNotFound creates a GetSelfServiceVerificationFlowNotFound with default headers values
func NewGetSelfServiceVerificationFlowNotFound() *GetSelfServiceVerificationFlowNotFound {
	return &GetSelfServiceVerificationFlowNotFound{}
}

/* GetSelfServiceVerificationFlowNotFound describes a response with status code 404, with default header values.

genericError
*/
type GetSelfServiceVerificationFlowNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceVerificationFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/verification/flows][%d] getSelfServiceVerificationFlowNotFound  %+v", 404, o.Payload)
}
func (o *GetSelfServiceVerificationFlowNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceVerificationFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceVerificationFlowInternalServerError creates a GetSelfServiceVerificationFlowInternalServerError with default headers values
func NewGetSelfServiceVerificationFlowInternalServerError() *GetSelfServiceVerificationFlowInternalServerError {
	return &GetSelfServiceVerificationFlowInternalServerError{}
}

/* GetSelfServiceVerificationFlowInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type GetSelfServiceVerificationFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceVerificationFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/verification/flows][%d] getSelfServiceVerificationFlowInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSelfServiceVerificationFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceVerificationFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
