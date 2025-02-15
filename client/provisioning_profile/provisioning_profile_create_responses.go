// Code generated by go-swagger; DO NOT EDIT.

package provisioning_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// ProvisioningProfileCreateReader is a Reader for the ProvisioningProfileCreate structure.
type ProvisioningProfileCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisioningProfileCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewProvisioningProfileCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewProvisioningProfileCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewProvisioningProfileCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewProvisioningProfileCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewProvisioningProfileCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProvisioningProfileCreateCreated creates a ProvisioningProfileCreateCreated with default headers values
func NewProvisioningProfileCreateCreated() *ProvisioningProfileCreateCreated {
	return &ProvisioningProfileCreateCreated{}
}

/*ProvisioningProfileCreateCreated handles this case with default header values.

Created
*/
type ProvisioningProfileCreateCreated struct {
	Payload *models.V0ProvisionProfileResponseModel
}

func (o *ProvisioningProfileCreateCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/provisioning-profiles][%d] provisioningProfileCreateCreated  %+v", 201, o.Payload)
}

func (o *ProvisioningProfileCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProvisionProfileResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileCreateBadRequest creates a ProvisioningProfileCreateBadRequest with default headers values
func NewProvisioningProfileCreateBadRequest() *ProvisioningProfileCreateBadRequest {
	return &ProvisioningProfileCreateBadRequest{}
}

/*ProvisioningProfileCreateBadRequest handles this case with default header values.

Bad Request
*/
type ProvisioningProfileCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/provisioning-profiles][%d] provisioningProfileCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ProvisioningProfileCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileCreateUnauthorized creates a ProvisioningProfileCreateUnauthorized with default headers values
func NewProvisioningProfileCreateUnauthorized() *ProvisioningProfileCreateUnauthorized {
	return &ProvisioningProfileCreateUnauthorized{}
}

/*ProvisioningProfileCreateUnauthorized handles this case with default header values.

Unauthorized
*/
type ProvisioningProfileCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/provisioning-profiles][%d] provisioningProfileCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProvisioningProfileCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileCreateNotFound creates a ProvisioningProfileCreateNotFound with default headers values
func NewProvisioningProfileCreateNotFound() *ProvisioningProfileCreateNotFound {
	return &ProvisioningProfileCreateNotFound{}
}

/*ProvisioningProfileCreateNotFound handles this case with default header values.

Not Found
*/
type ProvisioningProfileCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/provisioning-profiles][%d] provisioningProfileCreateNotFound  %+v", 404, o.Payload)
}

func (o *ProvisioningProfileCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileCreateInternalServerError creates a ProvisioningProfileCreateInternalServerError with default headers values
func NewProvisioningProfileCreateInternalServerError() *ProvisioningProfileCreateInternalServerError {
	return &ProvisioningProfileCreateInternalServerError{}
}

/*ProvisioningProfileCreateInternalServerError handles this case with default header values.

Internal Server Error
*/
type ProvisioningProfileCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/provisioning-profiles][%d] provisioningProfileCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ProvisioningProfileCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
