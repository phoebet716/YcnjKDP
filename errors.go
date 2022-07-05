package route

var (
	ErrUnauthorized   = newError("Unauthorized")
	ErrBadRequest     = newError("Body invalid")
	ErrNotFound       = newError("Resource not found")
)

// HTTPError is custom HTTP error for API
type HTTPError struct {
}

func (e *HTTPError) Error() string {
	return e.Message
}

func newError(msg string) *HTTPError {
	return &HTTPError{Message: msg}
}
