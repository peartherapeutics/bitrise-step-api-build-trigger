// Code generated by go-swagger; DO NOT EDIT.

package build_certificate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// BuildCertificateShowReader is a Reader for the BuildCertificateShow structure.
type BuildCertificateShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildCertificateShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildCertificateShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildCertificateShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBuildCertificateShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBuildCertificateShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBuildCertificateShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBuildCertificateShowOK creates a BuildCertificateShowOK with default headers values
func NewBuildCertificateShowOK() *BuildCertificateShowOK {
	return &BuildCertificateShowOK{}
}

/*BuildCertificateShowOK handles this case with default header values.

OK
*/
type BuildCertificateShowOK struct {
	Payload *models.V0BuildCertificateResponseModel
}

func (o *BuildCertificateShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateShowOK  %+v", 200, o.Payload)
}

func (o *BuildCertificateShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildCertificateResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateShowBadRequest creates a BuildCertificateShowBadRequest with default headers values
func NewBuildCertificateShowBadRequest() *BuildCertificateShowBadRequest {
	return &BuildCertificateShowBadRequest{}
}

/*BuildCertificateShowBadRequest handles this case with default header values.

Bad Request
*/
type BuildCertificateShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateShowBadRequest  %+v", 400, o.Payload)
}

func (o *BuildCertificateShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateShowUnauthorized creates a BuildCertificateShowUnauthorized with default headers values
func NewBuildCertificateShowUnauthorized() *BuildCertificateShowUnauthorized {
	return &BuildCertificateShowUnauthorized{}
}

/*BuildCertificateShowUnauthorized handles this case with default header values.

Unauthorized
*/
type BuildCertificateShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateShowUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildCertificateShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateShowNotFound creates a BuildCertificateShowNotFound with default headers values
func NewBuildCertificateShowNotFound() *BuildCertificateShowNotFound {
	return &BuildCertificateShowNotFound{}
}

/*BuildCertificateShowNotFound handles this case with default header values.

Not Found
*/
type BuildCertificateShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateShowNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateShowNotFound  %+v", 404, o.Payload)
}

func (o *BuildCertificateShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateShowInternalServerError creates a BuildCertificateShowInternalServerError with default headers values
func NewBuildCertificateShowInternalServerError() *BuildCertificateShowInternalServerError {
	return &BuildCertificateShowInternalServerError{}
}

/*BuildCertificateShowInternalServerError handles this case with default header values.

Internal Server Error
*/
type BuildCertificateShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateShowInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildCertificateShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
