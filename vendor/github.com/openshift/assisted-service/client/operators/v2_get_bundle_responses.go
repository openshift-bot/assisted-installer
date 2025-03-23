// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2GetBundleReader is a Reader for the V2GetBundle structure.
type V2GetBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewV2GetBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2GetBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2GetBundleOK creates a V2GetBundleOK with default headers values
func NewV2GetBundleOK() *V2GetBundleOK {
	return &V2GetBundleOK{}
}

/*
V2GetBundleOK describes a response with status code 200, with default header values.

Success
*/
type V2GetBundleOK struct {
	Payload *models.Bundle
}

// IsSuccess returns true when this v2 get bundle o k response has a 2xx status code
func (o *V2GetBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 get bundle o k response has a 3xx status code
func (o *V2GetBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get bundle o k response has a 4xx status code
func (o *V2GetBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get bundle o k response has a 5xx status code
func (o *V2GetBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get bundle o k response a status code equal to that given
func (o *V2GetBundleOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2GetBundleOK) Error() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleOK  %+v", 200, o.Payload)
}

func (o *V2GetBundleOK) String() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleOK  %+v", 200, o.Payload)
}

func (o *V2GetBundleOK) GetPayload() *models.Bundle {
	return o.Payload
}

func (o *V2GetBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Bundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetBundleNotFound creates a V2GetBundleNotFound with default headers values
func NewV2GetBundleNotFound() *V2GetBundleNotFound {
	return &V2GetBundleNotFound{}
}

/*
V2GetBundleNotFound describes a response with status code 404, with default header values.

Bundle not found
*/
type V2GetBundleNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get bundle not found response has a 2xx status code
func (o *V2GetBundleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get bundle not found response has a 3xx status code
func (o *V2GetBundleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get bundle not found response has a 4xx status code
func (o *V2GetBundleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get bundle not found response has a 5xx status code
func (o *V2GetBundleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get bundle not found response a status code equal to that given
func (o *V2GetBundleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2GetBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleNotFound  %+v", 404, o.Payload)
}

func (o *V2GetBundleNotFound) String() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleNotFound  %+v", 404, o.Payload)
}

func (o *V2GetBundleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetBundleInternalServerError creates a V2GetBundleInternalServerError with default headers values
func NewV2GetBundleInternalServerError() *V2GetBundleInternalServerError {
	return &V2GetBundleInternalServerError{}
}

/*
V2GetBundleInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type V2GetBundleInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get bundle internal server error response has a 2xx status code
func (o *V2GetBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get bundle internal server error response has a 3xx status code
func (o *V2GetBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get bundle internal server error response has a 4xx status code
func (o *V2GetBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get bundle internal server error response has a 5xx status code
func (o *V2GetBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 get bundle internal server error response a status code equal to that given
func (o *V2GetBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2GetBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetBundleInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/operators/bundles/{id}][%d] v2GetBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetBundleInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
