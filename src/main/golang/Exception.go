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

func throwNewException(message string, status uint16) Throwable {
	return &exception{message, status, false}
}

func throwNewRuntimeException(message string, status uint16) Throwable {
	return &exception{message, status, true}
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
