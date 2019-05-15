package gitconfig

type errNotFound struct {
	msg string
}

func (e *errNotFound) Error() string {
	return e.msg
}

func (e *errNotFound) NotFound() bool {
	return true
}

func notFound(msg string) error {
	return &errNotFound{msg: msg}
}

type notFounder interface {
	Error() string
	NotFound() bool
}

// IsNotFound returns a boolean indicating whether the error is known to
// report that a value does not found.
func IsNotFound(err error) bool {
	if nerr, ok := err.(notFounder); ok {
		return nerr.NotFound()
	}
	return false
}
