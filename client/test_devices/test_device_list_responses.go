// Code generated by go-swagger; DO NOT EDIT.

package test_devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// TestDeviceListReader is a Reader for the TestDeviceList structure.
type TestDeviceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestDeviceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestDeviceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTestDeviceListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewTestDeviceListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTestDeviceListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTestDeviceListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestDeviceListOK creates a TestDeviceListOK with default headers values
func NewTestDeviceListOK() *TestDeviceListOK {
	return &TestDeviceListOK{}
}

/*TestDeviceListOK handles this case with default header values.

OK
*/
type TestDeviceListOK struct {
	Payload *models.V0TestDeviceListResponseModel
}

func (o *TestDeviceListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/test-devices][%d] testDeviceListOK  %+v", 200, o.Payload)
}

func (o *TestDeviceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0TestDeviceListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestDeviceListBadRequest creates a TestDeviceListBadRequest with default headers values
func NewTestDeviceListBadRequest() *TestDeviceListBadRequest {
	return &TestDeviceListBadRequest{}
}

/*TestDeviceListBadRequest handles this case with default header values.

Bad Request
*/
type TestDeviceListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *TestDeviceListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/test-devices][%d] testDeviceListBadRequest  %+v", 400, o.Payload)
}

func (o *TestDeviceListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestDeviceListUnauthorized creates a TestDeviceListUnauthorized with default headers values
func NewTestDeviceListUnauthorized() *TestDeviceListUnauthorized {
	return &TestDeviceListUnauthorized{}
}

/*TestDeviceListUnauthorized handles this case with default header values.

Unauthorized
*/
type TestDeviceListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *TestDeviceListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/test-devices][%d] testDeviceListUnauthorized  %+v", 401, o.Payload)
}

func (o *TestDeviceListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestDeviceListNotFound creates a TestDeviceListNotFound with default headers values
func NewTestDeviceListNotFound() *TestDeviceListNotFound {
	return &TestDeviceListNotFound{}
}

/*TestDeviceListNotFound handles this case with default header values.

Not Found
*/
type TestDeviceListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *TestDeviceListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/test-devices][%d] testDeviceListNotFound  %+v", 404, o.Payload)
}

func (o *TestDeviceListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestDeviceListInternalServerError creates a TestDeviceListInternalServerError with default headers values
func NewTestDeviceListInternalServerError() *TestDeviceListInternalServerError {
	return &TestDeviceListInternalServerError{}
}

/*TestDeviceListInternalServerError handles this case with default header values.

Internal Server Error
*/
type TestDeviceListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *TestDeviceListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/test-devices][%d] testDeviceListInternalServerError  %+v", 500, o.Payload)
}

func (o *TestDeviceListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
