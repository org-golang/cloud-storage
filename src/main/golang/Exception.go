package golang

const (
	Success = 0
	Error   = 500
)

type exception struct {
	message string

	status uint16

	error bool
}

func throwNewException(message string, status uint16) *exception {
	return &exception{message, status, false}
}

func (e *exception) GetMessage() string {
	return e.message
}

func (e *exception) GetStatus() uint16 {
	return e.status
}

func (e *exception) Error() bool {
	return e.error
}

func (e *exception) throw(throw bool) Throwable {
	e.error = throw

	return e
}
