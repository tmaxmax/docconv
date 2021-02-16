package util

import "github.com/pkg/errors"

// ErrorPrefixer is a function that given an error prepends the passed message string to it.
type ErrorPrefixer func(error, string) error

// NewErrorPrefixer returns an ErrorPrefixer that prepends the given prefix to the error message.
func NewErrorPrefixer(prefix string) ErrorPrefixer {
	return func(err error, s string) error {
		return errors.WithMessage(err, prefix+": "+s)
	}
}

// DeferErrorHandler is used to handle deffered errors, such as those from
// closing an io.ReadCloser. It should be used in defer statements,
// with a pointer to the error variable the caller function will return.
//
// The function replaces returnError with defferedError, if the error returnError points to is nil.
func DeferErrorHandler(defferedError error, returnError *error) {
	if returnError == nil || *returnError != nil {
		return
	}
	*returnError = defferedError
}
