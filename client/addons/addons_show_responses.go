// Code generated by go-swagger; DO NOT EDIT.

package addons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// AddonsShowReader is a Reader for the AddonsShow structure.
type AddonsShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonsShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddonsShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddonsShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddonsShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddonsShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddonsShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddonsShowOK creates a AddonsShowOK with default headers values
func NewAddonsShowOK() *AddonsShowOK {
	return &AddonsShowOK{}
}

/*AddonsShowOK handles this case with default header values.

OK
*/
type AddonsShowOK struct {
	Payload *models.V0AddonsShowResponseModel
}

func (o *AddonsShowOK) Error() string {
	return fmt.Sprintf("[GET /addons/{addon-id}][%d] addonsShowOK  %+v", 200, o.Payload)
}

func (o *AddonsShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AddonsShowResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsShowBadRequest creates a AddonsShowBadRequest with default headers values
func NewAddonsShowBadRequest() *AddonsShowBadRequest {
	return &AddonsShowBadRequest{}
}

/*AddonsShowBadRequest handles this case with default header values.

Bad Request
*/
type AddonsShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonsShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /addons/{addon-id}][%d] addonsShowBadRequest  %+v", 400, o.Payload)
}

func (o *AddonsShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsShowUnauthorized creates a AddonsShowUnauthorized with default headers values
func NewAddonsShowUnauthorized() *AddonsShowUnauthorized {
	return &AddonsShowUnauthorized{}
}

/*AddonsShowUnauthorized handles this case with default header values.

Unauthorized
*/
type AddonsShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonsShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /addons/{addon-id}][%d] addonsShowUnauthorized  %+v", 401, o.Payload)
}

func (o *AddonsShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsShowNotFound creates a AddonsShowNotFound with default headers values
func NewAddonsShowNotFound() *AddonsShowNotFound {
	return &AddonsShowNotFound{}
}

/*AddonsShowNotFound handles this case with default header values.

Not Found
*/
type AddonsShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonsShowNotFound) Error() string {
	return fmt.Sprintf("[GET /addons/{addon-id}][%d] addonsShowNotFound  %+v", 404, o.Payload)
}

func (o *AddonsShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonsShowInternalServerError creates a AddonsShowInternalServerError with default headers values
func NewAddonsShowInternalServerError() *AddonsShowInternalServerError {
	return &AddonsShowInternalServerError{}
}

/*AddonsShowInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddonsShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonsShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /addons/{addon-id}][%d] addonsShowInternalServerError  %+v", 500, o.Payload)
}

func (o *AddonsShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
