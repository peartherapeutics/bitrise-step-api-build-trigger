// Code generated by go-swagger; DO NOT EDIT.

package webhook_delivery_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/peartherapeutics/bitrise-step-api-build-trigger/models"
)

// WebhookDeliveryItemShowReader is a Reader for the WebhookDeliveryItemShow structure.
type WebhookDeliveryItemShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebhookDeliveryItemShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWebhookDeliveryItemShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWebhookDeliveryItemShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewWebhookDeliveryItemShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWebhookDeliveryItemShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWebhookDeliveryItemShowOK creates a WebhookDeliveryItemShowOK with default headers values
func NewWebhookDeliveryItemShowOK() *WebhookDeliveryItemShowOK {
	return &WebhookDeliveryItemShowOK{}
}

/*WebhookDeliveryItemShowOK handles this case with default header values.

OK
*/
type WebhookDeliveryItemShowOK struct {
	Payload *models.V0WebhookDeliveryItemResponseModel
}

func (o *WebhookDeliveryItemShowOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}][%d] webhookDeliveryItemShowOK  %+v", 200, o.Payload)
}

func (o *WebhookDeliveryItemShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0WebhookDeliveryItemResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemShowBadRequest creates a WebhookDeliveryItemShowBadRequest with default headers values
func NewWebhookDeliveryItemShowBadRequest() *WebhookDeliveryItemShowBadRequest {
	return &WebhookDeliveryItemShowBadRequest{}
}

/*WebhookDeliveryItemShowBadRequest handles this case with default header values.

Bad Request
*/
type WebhookDeliveryItemShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *WebhookDeliveryItemShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}][%d] webhookDeliveryItemShowBadRequest  %+v", 400, o.Payload)
}

func (o *WebhookDeliveryItemShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemShowUnauthorized creates a WebhookDeliveryItemShowUnauthorized with default headers values
func NewWebhookDeliveryItemShowUnauthorized() *WebhookDeliveryItemShowUnauthorized {
	return &WebhookDeliveryItemShowUnauthorized{}
}

/*WebhookDeliveryItemShowUnauthorized handles this case with default header values.

Unauthorized
*/
type WebhookDeliveryItemShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *WebhookDeliveryItemShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}][%d] webhookDeliveryItemShowUnauthorized  %+v", 401, o.Payload)
}

func (o *WebhookDeliveryItemShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWebhookDeliveryItemShowInternalServerError creates a WebhookDeliveryItemShowInternalServerError with default headers values
func NewWebhookDeliveryItemShowInternalServerError() *WebhookDeliveryItemShowInternalServerError {
	return &WebhookDeliveryItemShowInternalServerError{}
}

/*WebhookDeliveryItemShowInternalServerError handles this case with default header values.

Internal Server Error
*/
type WebhookDeliveryItemShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *WebhookDeliveryItemShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}/delivery-items/{webhook-delivery-item-slug}][%d] webhookDeliveryItemShowInternalServerError  %+v", 500, o.Payload)
}

func (o *WebhookDeliveryItemShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
