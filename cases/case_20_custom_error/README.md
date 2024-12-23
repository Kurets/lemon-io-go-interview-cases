# Custom Error Wrapping

## Task Description

You need to create a custom error type that **wraps** another error. Go’s `errors` package recognizes wrapped errors if your type implements the **`Unwrap()`** method to return the wrapped cause. The type should also implement the **`Error()`** method.

```go
package main

type MyError struct {
	// Your code here, e.g.:
	// message string
	// cause   error
}

func (e *MyError) Error() string {
	// Should return a meaningful error message, possibly combining e.message + e.cause
	return ""
}

func (e *MyError) Unwrap() error {
	// Should return the wrapped cause if present, or nil otherwise
	return nil
}
```

### Requirements

1. **Fields**
    - You may store an underlying cause (type `error`) and a custom message (type `string`) within your struct.

2. **`Error()`**
    - Returns a string describing the error.
    - Typically includes the custom message and possibly details about the cause.

3. **`Unwrap()`**
    - Returns the wrapped cause (`error`), if any. If no cause, return `nil`.
    - This method enables `errors.Is` and `errors.Unwrap` to work correctly on your custom error.

4. **Behavior**
    - The built-in `errors.Unwrap()` must retrieve the underlying error from your custom type.
    - `errors.Is(yourError, someOtherError)` should work as expected if `someOtherError` is the same as (or part of the chain from) your underlying cause.

5. **Edge Cases**
    - **Nil cause**: If your custom error has no underlying cause, `Unwrap()` should return `nil`.
    - **Empty message**: It’s valid to have an empty message, though typically an error message is expected.

### Example Usage

``` go
package main

import (
	"errors"
	"fmt"
)

func main() {
	orig := errors.New("original error")
	myErr := &MyError{
		message: "Something went wrong",
		cause:   orig,
	}

	fmt.Println("MyError says:", myErr) // calls .Error()

	// Unwrap the error to see if we get the original cause
	fmt.Println("Unwrapped cause:", errors.Unwrap(myErr))
}
```

### Hints

- Make sure your `MyError` struct fields are exported or at least accessible within the package, if needed by the tests (though the tests might rely purely on `Error()` and `Unwrap()`).
- Often, the string returned by `Error()` might look like `fmt.Sprintf("%s: %v", e.message, e.cause)`, but you can customize.
- If you set `cause` to `nil`, `Unwrap()` must return `nil`. This is important for `errors.Is` and `errors.Unwrap` usage.
- Pay attention to clarity. A common pattern is:

```go 
func (e *MyError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}
```
