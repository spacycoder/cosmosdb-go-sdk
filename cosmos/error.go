package cosmos

import "fmt"

type CosmosErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error reports cosmos related errors.
type Error struct {
	message    *CosmosErrorMessage
	statusCode int
}

// NewCosmosError creates a new cosmos error struct
func NewCosmosError(message *CosmosErrorMessage, statusCode int) *Error {
	return &Error{
		message:    message,
		statusCode: statusCode,
	}
}

// Error implements the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%d, %s, %s", e.statusCode, e.message.Code, e.message.Message)
}

// StatusCode returns status code.
func (e *Error) StatusCode() int {
	return e.statusCode
}

// StatusOK checks is cosmos response was ok.
func (e *Error) StatusOK() bool {
	return e.statusCode == 200
}

// Created will be true if a resource was created.
func (e *Error) Created() bool {
	return e.statusCode == 201
}

// BadRequest see `https://docs.microsoft.com/en-us/rest/api/cosmos-db/http-status-codes-for-cosmosdb` for possible cause.
func (e *Error) BadRequest() bool {
	return e.statusCode == 400
}

// Unauthorized is true of operation was not authorized.
func (e *Error) Unauthorized() bool {
	return e.statusCode == 401
}

// Forbidden see `https://docs.microsoft.com/en-us/rest/api/cosmos-db/http-status-codes-for-cosmosdb` for possible cause.
func (e *Error) Forbidden() bool {
	return e.statusCode == 403
}

// NotFound could not find resource.
func (e *Error) NotFound() bool {
	return e.statusCode == 404
}

// RequestTimeout the operation did not complete within the allotted amount of time.
func (e *Error) RequestTimeout() bool {
	return e.statusCode == 408
}

// Conflict returns true if the ID for a resource has already been taken by an existing resource.
// Only applicable to PUT and POST request. e.g. Create and Replace.
func (e *Error) Conflict() bool {
	return e.statusCode == 409
}

// EntityTooLarge returns true if the request exceeded the allowable document size for a request. the max allowable document size is 2 MB.
func (e *Error) EntityTooLarge() bool {
	return e.statusCode == 413
}

// The collection has exceeded the provisioned throughput limit. Retry the request after the server specified retry after duration.
func (e *Error) TooManyRequests() bool {
	return e.statusCode == 429
}

// RetryWith returns true if the operation encountered a transient error. This code only occurs on write operations. It is safe to retry the operation.
func (e *Error) RetryWith() bool {
	return e.statusCode == 449
}

// ServiceUnavailable operation could not be completed because the service was unavailable
func (e *Error) ServiceUnavailable() bool {
	return e.statusCode == 503
}
