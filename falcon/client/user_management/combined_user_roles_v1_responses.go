// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// CombinedUserRolesV1Reader is a Reader for the CombinedUserRolesV1 structure.
type CombinedUserRolesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CombinedUserRolesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCombinedUserRolesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCombinedUserRolesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCombinedUserRolesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCombinedUserRolesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCombinedUserRolesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCombinedUserRolesV1OK creates a CombinedUserRolesV1OK with default headers values
func NewCombinedUserRolesV1OK() *CombinedUserRolesV1OK {
	return &CombinedUserRolesV1OK{}
}

/*
CombinedUserRolesV1OK describes a response with status code 200, with default header values.

OK
*/
type CombinedUserRolesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaUserGrantsResponse
}

// IsSuccess returns true when this combined user roles v1 o k response has a 2xx status code
func (o *CombinedUserRolesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this combined user roles v1 o k response has a 3xx status code
func (o *CombinedUserRolesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined user roles v1 o k response has a 4xx status code
func (o *CombinedUserRolesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined user roles v1 o k response has a 5xx status code
func (o *CombinedUserRolesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this combined user roles v1 o k response a status code equal to that given
func (o *CombinedUserRolesV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *CombinedUserRolesV1OK) Error() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1OK  %+v", 200, o.Payload)
}

func (o *CombinedUserRolesV1OK) String() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1OK  %+v", 200, o.Payload)
}

func (o *CombinedUserRolesV1OK) GetPayload() *models.DomainMsaUserGrantsResponse {
	return o.Payload
}

func (o *CombinedUserRolesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaUserGrantsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedUserRolesV1BadRequest creates a CombinedUserRolesV1BadRequest with default headers values
func NewCombinedUserRolesV1BadRequest() *CombinedUserRolesV1BadRequest {
	return &CombinedUserRolesV1BadRequest{}
}

/*
CombinedUserRolesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CombinedUserRolesV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaUserGrantsResponse
}

// IsSuccess returns true when this combined user roles v1 bad request response has a 2xx status code
func (o *CombinedUserRolesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined user roles v1 bad request response has a 3xx status code
func (o *CombinedUserRolesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined user roles v1 bad request response has a 4xx status code
func (o *CombinedUserRolesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined user roles v1 bad request response has a 5xx status code
func (o *CombinedUserRolesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this combined user roles v1 bad request response a status code equal to that given
func (o *CombinedUserRolesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CombinedUserRolesV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CombinedUserRolesV1BadRequest) String() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CombinedUserRolesV1BadRequest) GetPayload() *models.DomainMsaUserGrantsResponse {
	return o.Payload
}

func (o *CombinedUserRolesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaUserGrantsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedUserRolesV1Forbidden creates a CombinedUserRolesV1Forbidden with default headers values
func NewCombinedUserRolesV1Forbidden() *CombinedUserRolesV1Forbidden {
	return &CombinedUserRolesV1Forbidden{}
}

/*
CombinedUserRolesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CombinedUserRolesV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaUserGrantsResponse
}

// IsSuccess returns true when this combined user roles v1 forbidden response has a 2xx status code
func (o *CombinedUserRolesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined user roles v1 forbidden response has a 3xx status code
func (o *CombinedUserRolesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined user roles v1 forbidden response has a 4xx status code
func (o *CombinedUserRolesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined user roles v1 forbidden response has a 5xx status code
func (o *CombinedUserRolesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this combined user roles v1 forbidden response a status code equal to that given
func (o *CombinedUserRolesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CombinedUserRolesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CombinedUserRolesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CombinedUserRolesV1Forbidden) GetPayload() *models.DomainMsaUserGrantsResponse {
	return o.Payload
}

func (o *CombinedUserRolesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaUserGrantsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedUserRolesV1TooManyRequests creates a CombinedUserRolesV1TooManyRequests with default headers values
func NewCombinedUserRolesV1TooManyRequests() *CombinedUserRolesV1TooManyRequests {
	return &CombinedUserRolesV1TooManyRequests{}
}

/*
CombinedUserRolesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CombinedUserRolesV1TooManyRequests struct {

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

// IsSuccess returns true when this combined user roles v1 too many requests response has a 2xx status code
func (o *CombinedUserRolesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined user roles v1 too many requests response has a 3xx status code
func (o *CombinedUserRolesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined user roles v1 too many requests response has a 4xx status code
func (o *CombinedUserRolesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined user roles v1 too many requests response has a 5xx status code
func (o *CombinedUserRolesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this combined user roles v1 too many requests response a status code equal to that given
func (o *CombinedUserRolesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CombinedUserRolesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedUserRolesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedUserRolesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedUserRolesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedUserRolesV1InternalServerError creates a CombinedUserRolesV1InternalServerError with default headers values
func NewCombinedUserRolesV1InternalServerError() *CombinedUserRolesV1InternalServerError {
	return &CombinedUserRolesV1InternalServerError{}
}

/*
CombinedUserRolesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CombinedUserRolesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaUserGrantsResponse
}

// IsSuccess returns true when this combined user roles v1 internal server error response has a 2xx status code
func (o *CombinedUserRolesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined user roles v1 internal server error response has a 3xx status code
func (o *CombinedUserRolesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined user roles v1 internal server error response has a 4xx status code
func (o *CombinedUserRolesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined user roles v1 internal server error response has a 5xx status code
func (o *CombinedUserRolesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this combined user roles v1 internal server error response a status code equal to that given
func (o *CombinedUserRolesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CombinedUserRolesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedUserRolesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /user-management/combined/user-roles/v1][%d] combinedUserRolesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedUserRolesV1InternalServerError) GetPayload() *models.DomainMsaUserGrantsResponse {
	return o.Payload
}

func (o *CombinedUserRolesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaUserGrantsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
