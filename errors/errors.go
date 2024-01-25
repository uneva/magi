package errors

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/uneva/magi/errors/v1"
)

const (
	UnknownCode   = http.StatusInternalServerError
	UnknownReason = ""
)

type Error struct {
	errors.Status
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v cause = %v", e.Code, e.Reason, e.Message, e.Metadata, e.cause)
}

func (e *Error) Unwrap() error { return e.cause }

func (e *Error) Is(err error) bool {
	if se := new(Error); stderrors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

func (e *Error) WithMetadata(md map[string]string) *Error {
	err := Clone(e)
	err.Metadata = md
	return err
}

func New(code int, reason, message string) *Error {
	return &Error{
		Status: errors.Status{
			Code:    int32(code),
			Message: message,
			Reason:  reason,
		},
	}
}

func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

func Errorf(code int, reason, format string, a ...interface{}) error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}
	return int(FromError(err).Code)
}

func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}

func Clone(err *Error) *Error {
	if err == nil {
		return nil
	}
	metadata := make(map[string]string, len(err.Metadata))
	for k, v := range err.Metadata {
		metadata[k] = v
	}
	return &Error{
		cause: err.cause,
		Status: errors.Status{
			Code:     err.Code,
			Reason:   err.Reason,
			Message:  err.Message,
			Metadata: metadata,
		},
	}
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); stderrors.As(err, &se) {
		return se
	}
	return New(UnknownCode, UnknownReason, err.Error())
}
