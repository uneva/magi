package errors

import "net/http"

func BadRequest(reason, message string) *Error {
	return New(http.StatusBadRequest, reason, message)
}

func IsBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

func Unauthorized(reason, message string) *Error {
	return New(http.StatusUnauthorized, reason, message)
}

func IsUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

func Forbidden(reason, message string) *Error {
	return New(http.StatusForbidden, reason, message)
}

func IsForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

func NotFound(reason, message string) *Error {
	return New(http.StatusNotFound, reason, message)
}

func IsNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

func InternalServer(reason, message string) *Error {
	return New(http.StatusInternalServerError, reason, message)
}

func IsInternalServer(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

func ServiceUnavailable(reason, message string) *Error {
	return New(http.StatusServiceUnavailable, reason, message)
}

func IsServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

func GatewayTimeout(reason, message string) *Error {
	return New(http.StatusGatewayTimeout, reason, message)
}

func IsGatewayTimeout(err error) bool {
	return Code(err) == http.StatusGatewayTimeout
}
