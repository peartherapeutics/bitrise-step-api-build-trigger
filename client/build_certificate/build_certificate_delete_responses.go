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

// BuildCertificateDeleteReader is a Reader for the BuildCertificateDelete structure.
type BuildCertificateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildCertificateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildCertificateDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildCertificateDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBuildCertificateDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBuildCertificateDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBuildCertificateDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBuildCertificateDeleteOK creates a BuildCertificateDeleteOK with default headers values
func NewBuildCertificateDeleteOK() *BuildCertificateDeleteOK {
	return &BuildCertificateDeleteOK{}
}

/*BuildCertificateDeleteOK handles this case with default header values.

OK
*/
type BuildCertificateDeleteOK struct {
	Payload *models.V0BuildCertificateResponseModel
}

func (o *BuildCertificateDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateDeleteOK  %+v", 200, o.Payload)
}

func (o *BuildCertificateDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildCertificateResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateDeleteBadRequest creates a BuildCertificateDeleteBadRequest with default headers values
func NewBuildCertificateDeleteBadRequest() *BuildCertificateDeleteBadRequest {
	return &BuildCertificateDeleteBadRequest{}
}

/*BuildCertificateDeleteBadRequest handles this case with default header values.

Bad Request
*/
type BuildCertificateDeleteBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *BuildCertificateDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateDeleteUnauthorized creates a BuildCertificateDeleteUnauthorized with default headers values
func NewBuildCertificateDeleteUnauthorized() *BuildCertificateDeleteUnauthorized {
	return &BuildCertificateDeleteUnauthorized{}
}

/*BuildCertificateDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type BuildCertificateDeleteUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildCertificateDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateDeleteNotFound creates a BuildCertificateDeleteNotFound with default headers values
func NewBuildCertificateDeleteNotFound() *BuildCertificateDeleteNotFound {
	return &BuildCertificateDeleteNotFound{}
}

/*BuildCertificateDeleteNotFound handles this case with default header values.

Not Found
*/
type BuildCertificateDeleteNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateDeleteNotFound  %+v", 404, o.Payload)
}

func (o *BuildCertificateDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateDeleteInternalServerError creates a BuildCertificateDeleteInternalServerError with default headers values
func NewBuildCertificateDeleteInternalServerError() *BuildCertificateDeleteInternalServerError {
	return &BuildCertificateDeleteInternalServerError{}
}

/*BuildCertificateDeleteInternalServerError handles this case with default header values.

Internal Server Error
*/
type BuildCertificateDeleteInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildCertificateDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
