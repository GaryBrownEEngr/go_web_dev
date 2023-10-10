package stacktrs

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type StackError struct {
	Err   error
	Stack string
}

func (e *StackError) Error() string {
	return fmt.Sprintf("Error: %s\nStack: %s", e.Err, e.Stack)
}

func (e *StackError) Unwrap() error {
	return e.Err
}

func Wrap(err error) error {
	return wrap(debug.Stack, err)
}

func Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	return wrap(debug.Stack, err)
}

func wrap(getStack func() []byte, err error) error {
	if err == nil {
		return nil
	}

	// Don't compound stack if we have already created a stack.
	var stackErr *StackError
	if ok := errors.As(err, &stackErr); ok {
		return err
	}

	return &StackError{Err: err, Stack: string(getStack())}
}
