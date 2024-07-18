package lang

const (
	Success = 0
	Error   = 500
)

type exception struct {
	message string

	status uint16
}

func ThrowNewException(message string, status uint16) Throwable {
	return &exception{message, status}
}

func (e *exception) GetMessage() string {
	return e.message
}

func (e *exception) GetStatus() uint16 {
	return e.status
}
