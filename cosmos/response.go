package cosmos

import (
	"math"
	"net/http"
)

type Response struct {
	Header http.Header
}

// Continuation returns continuation token for paged request.
// Pass this value to next request to get next page of documents.
func (r *Response) Continuation() string {
	return r.Header.Get(HeaderContinuation)
}

// RetryAfterMs returns the number of seconds until you can try to send the response again
// only applicable to status code 429 To Many Requests
func (r *Response) RetryAfterMs() string {
	return r.Header.Get(HeaderRetryAfterMs)
}

type statusCodeValidatorFunc func(statusCode int) bool

func expectStatusCode(expected int) statusCodeValidatorFunc {
	return func(statusCode int) bool {
		return expected == statusCode
	}
}

func expectStatusCodeXX(expected int) statusCodeValidatorFunc {
	beginning := int(math.Floor(float64(expected/100))) * 100
	end := beginning + 99
	return func(statusCode int) bool {
		return (statusCode >= beginning) && (statusCode <= end)
	}
}
