package util

type HttpError interface {
	Status() int
	Error() string
}

func NewHttpError(status int, err error) HttpError {
	return &httpError{
		status: status,
		err:    err,
	}
}

type httpError struct {
	status int
	err    error
}

func (e *httpError) Error() string {
	return e.err.Error()
}

func (e *httpError) Status() int {
	return e.status
}
