// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateSensorUpdatePoliciesV2Reader is a Reader for the UpdateSensorUpdatePoliciesV2 structure.
type UpdateSensorUpdatePoliciesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSensorUpdatePoliciesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSensorUpdatePoliciesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSensorUpdatePoliciesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSensorUpdatePoliciesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSensorUpdatePoliciesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateSensorUpdatePoliciesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSensorUpdatePoliciesV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSensorUpdatePoliciesV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSensorUpdatePoliciesV2OK creates a UpdateSensorUpdatePoliciesV2OK with default headers values
func NewUpdateSensorUpdatePoliciesV2OK() *UpdateSensorUpdatePoliciesV2OK {
	return &UpdateSensorUpdatePoliciesV2OK{}
}

/*
UpdateSensorUpdatePoliciesV2OK describes a response with status code 200, with default header values.

OK
*/
type UpdateSensorUpdatePoliciesV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this update sensor update policies v2 o k response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update sensor update policies v2 o k response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 o k response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update sensor update policies v2 o k response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update sensor update policies v2 o k response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2OK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateSensorUpdatePoliciesV2OK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2OK  %+v", 200, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2OK) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2OK  %+v", 200, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2OK) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2BadRequest creates a UpdateSensorUpdatePoliciesV2BadRequest with default headers values
func NewUpdateSensorUpdatePoliciesV2BadRequest() *UpdateSensorUpdatePoliciesV2BadRequest {
	return &UpdateSensorUpdatePoliciesV2BadRequest{}
}

/*
UpdateSensorUpdatePoliciesV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSensorUpdatePoliciesV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this update sensor update policies v2 bad request response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update sensor update policies v2 bad request response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 bad request response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update sensor update policies v2 bad request response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update sensor update policies v2 bad request response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateSensorUpdatePoliciesV2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2BadRequest) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2BadRequest) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2Forbidden creates a UpdateSensorUpdatePoliciesV2Forbidden with default headers values
func NewUpdateSensorUpdatePoliciesV2Forbidden() *UpdateSensorUpdatePoliciesV2Forbidden {
	return &UpdateSensorUpdatePoliciesV2Forbidden{}
}

/*
UpdateSensorUpdatePoliciesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateSensorUpdatePoliciesV2Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this update sensor update policies v2 forbidden response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update sensor update policies v2 forbidden response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 forbidden response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update sensor update policies v2 forbidden response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update sensor update policies v2 forbidden response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateSensorUpdatePoliciesV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2Forbidden) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2NotFound creates a UpdateSensorUpdatePoliciesV2NotFound with default headers values
func NewUpdateSensorUpdatePoliciesV2NotFound() *UpdateSensorUpdatePoliciesV2NotFound {
	return &UpdateSensorUpdatePoliciesV2NotFound{}
}

/*
UpdateSensorUpdatePoliciesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateSensorUpdatePoliciesV2NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this update sensor update policies v2 not found response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update sensor update policies v2 not found response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 not found response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update sensor update policies v2 not found response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update sensor update policies v2 not found response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateSensorUpdatePoliciesV2NotFound) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2NotFound  %+v", 404, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2NotFound) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2NotFound  %+v", 404, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2NotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2TooManyRequests creates a UpdateSensorUpdatePoliciesV2TooManyRequests with default headers values
func NewUpdateSensorUpdatePoliciesV2TooManyRequests() *UpdateSensorUpdatePoliciesV2TooManyRequests {
	return &UpdateSensorUpdatePoliciesV2TooManyRequests{}
}

/*
UpdateSensorUpdatePoliciesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateSensorUpdatePoliciesV2TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update sensor update policies v2 too many requests response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update sensor update policies v2 too many requests response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 too many requests response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update sensor update policies v2 too many requests response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update sensor update policies v2 too many requests response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2InternalServerError creates a UpdateSensorUpdatePoliciesV2InternalServerError with default headers values
func NewUpdateSensorUpdatePoliciesV2InternalServerError() *UpdateSensorUpdatePoliciesV2InternalServerError {
	return &UpdateSensorUpdatePoliciesV2InternalServerError{}
}

/*
UpdateSensorUpdatePoliciesV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateSensorUpdatePoliciesV2InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// IsSuccess returns true when this update sensor update policies v2 internal server error response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update sensor update policies v2 internal server error response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update sensor update policies v2 internal server error response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update sensor update policies v2 internal server error response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update sensor update policies v2 internal server error response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateSensorUpdatePoliciesV2InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2InternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSensorUpdatePoliciesV2Default creates a UpdateSensorUpdatePoliciesV2Default with default headers values
func NewUpdateSensorUpdatePoliciesV2Default(code int) *UpdateSensorUpdatePoliciesV2Default {
	return &UpdateSensorUpdatePoliciesV2Default{
		_statusCode: code,
	}
}

/*
UpdateSensorUpdatePoliciesV2Default describes a response with status code -1, with default header values.

OK
*/
type UpdateSensorUpdatePoliciesV2Default struct {
	_statusCode int

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// Code gets the status code for the update sensor update policies v2 default response
func (o *UpdateSensorUpdatePoliciesV2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update sensor update policies v2 default response has a 2xx status code
func (o *UpdateSensorUpdatePoliciesV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update sensor update policies v2 default response has a 3xx status code
func (o *UpdateSensorUpdatePoliciesV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update sensor update policies v2 default response has a 4xx status code
func (o *UpdateSensorUpdatePoliciesV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update sensor update policies v2 default response has a 5xx status code
func (o *UpdateSensorUpdatePoliciesV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update sensor update policies v2 default response a status code equal to that given
func (o *UpdateSensorUpdatePoliciesV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateSensorUpdatePoliciesV2Default) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2Default) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/sensor-update/v2][%d] updateSensorUpdatePoliciesV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSensorUpdatePoliciesV2Default) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *UpdateSensorUpdatePoliciesV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
