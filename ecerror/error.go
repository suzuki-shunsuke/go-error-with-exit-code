package ecerror

import (
	"errors"
)

type withExitCodeError struct {
	err  error
	code int
}

func Wrap(err error, code int) error {
	return withExitCodeError{
		err:  err,
		code: code,
	}
}

func (err withExitCodeError) ExitCode() int {
	return err.code
}

func (err withExitCodeError) Error() string {
	if err.err == nil {
		return ""
	}
	return err.err.Error()
}

func (err withExitCodeError) Unwrap() error {
	return err.err
}

func GetExitCode(err error) int {
	if err == nil {
		return 0
	}
	var ecerr withExitCodeError
	if errors.As(err, &ecerr) {
		return ecerr.ExitCode()
	}
	return 1
}
