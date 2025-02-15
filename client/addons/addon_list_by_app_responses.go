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

// AddonListByAppReader is a Reader for the AddonListByApp structure.
type AddonListByAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonListByAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddonListByAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddonListByAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAddonListByAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddonListByAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddonListByAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddonListByAppOK creates a AddonListByAppOK with default headers values
func NewAddonListByAppOK() *AddonListByAppOK {
	return &AddonListByAppOK{}
}

/*AddonListByAppOK handles this case with default header values.

OK
*/
type AddonListByAppOK struct {
	Payload *models.V0AppAddOnsListResponseModel
}

func (o *AddonListByAppOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/addons][%d] addonListByAppOK  %+v", 200, o.Payload)
}

func (o *AddonListByAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppAddOnsListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByAppBadRequest creates a AddonListByAppBadRequest with default headers values
func NewAddonListByAppBadRequest() *AddonListByAppBadRequest {
	return &AddonListByAppBadRequest{}
}

/*AddonListByAppBadRequest handles this case with default header values.

Bad Request
*/
type AddonListByAppBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByAppBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/addons][%d] addonListByAppBadRequest  %+v", 400, o.Payload)
}

func (o *AddonListByAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByAppUnauthorized creates a AddonListByAppUnauthorized with default headers values
func NewAddonListByAppUnauthorized() *AddonListByAppUnauthorized {
	return &AddonListByAppUnauthorized{}
}

/*AddonListByAppUnauthorized handles this case with default header values.

Unauthorized
*/
type AddonListByAppUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByAppUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/addons][%d] addonListByAppUnauthorized  %+v", 401, o.Payload)
}

func (o *AddonListByAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByAppNotFound creates a AddonListByAppNotFound with default headers values
func NewAddonListByAppNotFound() *AddonListByAppNotFound {
	return &AddonListByAppNotFound{}
}

/*AddonListByAppNotFound handles this case with default header values.

Not Found
*/
type AddonListByAppNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByAppNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/addons][%d] addonListByAppNotFound  %+v", 404, o.Payload)
}

func (o *AddonListByAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByAppInternalServerError creates a AddonListByAppInternalServerError with default headers values
func NewAddonListByAppInternalServerError() *AddonListByAppInternalServerError {
	return &AddonListByAppInternalServerError{}
}

/*AddonListByAppInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddonListByAppInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByAppInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/addons][%d] addonListByAppInternalServerError  %+v", 500, o.Payload)
}

func (o *AddonListByAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
