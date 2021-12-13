// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// QueryIntelIndicatorIdsReader is a Reader for the QueryIntelIndicatorIds structure.
type QueryIntelIndicatorIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryIntelIndicatorIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryIntelIndicatorIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryIntelIndicatorIdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryIntelIndicatorIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryIntelIndicatorIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryIntelIndicatorIdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryIntelIndicatorIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryIntelIndicatorIdsOK creates a QueryIntelIndicatorIdsOK with default headers values
func NewQueryIntelIndicatorIdsOK() *QueryIntelIndicatorIdsOK {
	return &QueryIntelIndicatorIdsOK{}
}

/* QueryIntelIndicatorIdsOK describes a response with status code 200, with default header values.

OK
*/
type QueryIntelIndicatorIdsOK struct {

	/* Provides next page pagination URL. Available only if sorting was done using using _marker field, which is the default one.
	 */
	NextPage string

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryIntelIndicatorIdsOK) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] queryIntelIndicatorIdsOK  %+v", 200, o.Payload)
}
func (o *QueryIntelIndicatorIdsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Next-Page
	hdrNextPage := response.GetHeader("Next-Page")

	if hdrNextPage != "" {
		o.NextPage = hdrNextPage
	}

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorIdsBadRequest creates a QueryIntelIndicatorIdsBadRequest with default headers values
func NewQueryIntelIndicatorIdsBadRequest() *QueryIntelIndicatorIdsBadRequest {
	return &QueryIntelIndicatorIdsBadRequest{}
}

/* QueryIntelIndicatorIdsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryIntelIndicatorIdsBadRequest struct {

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

func (o *QueryIntelIndicatorIdsBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] queryIntelIndicatorIdsBadRequest  %+v", 400, o.Payload)
}
func (o *QueryIntelIndicatorIdsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelIndicatorIdsForbidden creates a QueryIntelIndicatorIdsForbidden with default headers values
func NewQueryIntelIndicatorIdsForbidden() *QueryIntelIndicatorIdsForbidden {
	return &QueryIntelIndicatorIdsForbidden{}
}

/* QueryIntelIndicatorIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryIntelIndicatorIdsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryIntelIndicatorIdsForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] queryIntelIndicatorIdsForbidden  %+v", 403, o.Payload)
}
func (o *QueryIntelIndicatorIdsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIntelIndicatorIdsTooManyRequests creates a QueryIntelIndicatorIdsTooManyRequests with default headers values
func NewQueryIntelIndicatorIdsTooManyRequests() *QueryIntelIndicatorIdsTooManyRequests {
	return &QueryIntelIndicatorIdsTooManyRequests{}
}

/* QueryIntelIndicatorIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryIntelIndicatorIdsTooManyRequests struct {

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

func (o *QueryIntelIndicatorIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] queryIntelIndicatorIdsTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryIntelIndicatorIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelIndicatorIdsInternalServerError creates a QueryIntelIndicatorIdsInternalServerError with default headers values
func NewQueryIntelIndicatorIdsInternalServerError() *QueryIntelIndicatorIdsInternalServerError {
	return &QueryIntelIndicatorIdsInternalServerError{}
}

/* QueryIntelIndicatorIdsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryIntelIndicatorIdsInternalServerError struct {

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

func (o *QueryIntelIndicatorIdsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] queryIntelIndicatorIdsInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryIntelIndicatorIdsInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelIndicatorIdsDefault creates a QueryIntelIndicatorIdsDefault with default headers values
func NewQueryIntelIndicatorIdsDefault(code int) *QueryIntelIndicatorIdsDefault {
	return &QueryIntelIndicatorIdsDefault{
		_statusCode: code,
	}
}

/* QueryIntelIndicatorIdsDefault describes a response with status code -1, with default header values.

OK
*/
type QueryIntelIndicatorIdsDefault struct {
	_statusCode int

	/* Provides next page pagination URL. Available only if sorting was done using using _marker field, which is the default one.
	 */
	NextPage string

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query intel indicator ids default response
func (o *QueryIntelIndicatorIdsDefault) Code() int {
	return o._statusCode
}

func (o *QueryIntelIndicatorIdsDefault) Error() string {
	return fmt.Sprintf("[GET /intel/queries/indicators/v1][%d] QueryIntelIndicatorIds default  %+v", o._statusCode, o.Payload)
}
func (o *QueryIntelIndicatorIdsDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryIntelIndicatorIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Next-Page
	hdrNextPage := response.GetHeader("Next-Page")

	if hdrNextPage != "" {
		o.NextPage = hdrNextPage
	}

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
