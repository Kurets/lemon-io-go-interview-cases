package case_20_custom_error

type MyError struct {
	msg string
	err error
}

func (e *MyError) Error() string {
	return e.msg + ": " + e.err.Error()
}

func (e *MyError) Unwrap() error {
	return e.err
}
