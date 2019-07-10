// Code generated by go-swagger; DO NOT EDIT.

package avatar_candidate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// AvatarCandidateListReader is a Reader for the AvatarCandidateList structure.
type AvatarCandidateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvatarCandidateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAvatarCandidateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAvatarCandidateListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAvatarCandidateListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAvatarCandidateListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAvatarCandidateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvatarCandidateListOK creates a AvatarCandidateListOK with default headers values
func NewAvatarCandidateListOK() *AvatarCandidateListOK {
	return &AvatarCandidateListOK{}
}

/*AvatarCandidateListOK handles this case with default header values.

OK
*/
type AvatarCandidateListOK struct {
	Payload *models.V0FindAvatarCandidateResponse
}

func (o *AvatarCandidateListOK) Error() string {
	return fmt.Sprintf("[GET /v0.1/apps/{app-slug}/avatar-candidates][%d] avatarCandidateListOK  %+v", 200, o.Payload)
}

func (o *AvatarCandidateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0FindAvatarCandidateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateListBadRequest creates a AvatarCandidateListBadRequest with default headers values
func NewAvatarCandidateListBadRequest() *AvatarCandidateListBadRequest {
	return &AvatarCandidateListBadRequest{}
}

/*AvatarCandidateListBadRequest handles this case with default header values.

Bad Request
*/
type AvatarCandidateListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateListBadRequest) Error() string {
	return fmt.Sprintf("[GET /v0.1/apps/{app-slug}/avatar-candidates][%d] avatarCandidateListBadRequest  %+v", 400, o.Payload)
}

func (o *AvatarCandidateListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateListUnauthorized creates a AvatarCandidateListUnauthorized with default headers values
func NewAvatarCandidateListUnauthorized() *AvatarCandidateListUnauthorized {
	return &AvatarCandidateListUnauthorized{}
}

/*AvatarCandidateListUnauthorized handles this case with default header values.

Unauthorized
*/
type AvatarCandidateListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v0.1/apps/{app-slug}/avatar-candidates][%d] avatarCandidateListUnauthorized  %+v", 401, o.Payload)
}

func (o *AvatarCandidateListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateListNotFound creates a AvatarCandidateListNotFound with default headers values
func NewAvatarCandidateListNotFound() *AvatarCandidateListNotFound {
	return &AvatarCandidateListNotFound{}
}

/*AvatarCandidateListNotFound handles this case with default header values.

Not Found
*/
type AvatarCandidateListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateListNotFound) Error() string {
	return fmt.Sprintf("[GET /v0.1/apps/{app-slug}/avatar-candidates][%d] avatarCandidateListNotFound  %+v", 404, o.Payload)
}

func (o *AvatarCandidateListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateListInternalServerError creates a AvatarCandidateListInternalServerError with default headers values
func NewAvatarCandidateListInternalServerError() *AvatarCandidateListInternalServerError {
	return &AvatarCandidateListInternalServerError{}
}

/*AvatarCandidateListInternalServerError handles this case with default header values.

Internal Server Error
*/
type AvatarCandidateListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v0.1/apps/{app-slug}/avatar-candidates][%d] avatarCandidateListInternalServerError  %+v", 500, o.Payload)
}

func (o *AvatarCandidateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
