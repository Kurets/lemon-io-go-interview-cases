package case_20_custom_error

import (
	"errors"
	"testing"
)

func TestMyErrorWrapping(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := &MyError{
		msg: "something went wrong",
		err: originalErr,
	}

	errMsg := wrappedErr.Error()
	if errMsg == "" {
		t.Error("expected non-empty error message, got empty string")
	}

	if !errors.Is(wrappedErr, originalErr) {
		t.Errorf("expected errors.Is(wrappedErr, originalErr) to be true, got false")
	}

	unwrapped := errors.Unwrap(wrappedErr)
	if unwrapped != originalErr {
		t.Errorf("expected unwrapped error to be %v, got %v", originalErr, unwrapped)
	}
}

func TestNilCause(t *testing.T) {
	wrappedErr := &MyError{
		msg: "no cause here",
		err: nil,
	}

	unwrapped := errors.Unwrap(wrappedErr)
	if unwrapped != nil {
		t.Errorf("expected nil from Unwrap, got %v", unwrapped)
	}
}
