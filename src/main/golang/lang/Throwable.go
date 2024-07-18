package lang

type Throwable interface {

	// GetMessage Get exception error message.
	GetMessage() string

	// GetStatus Get exception error code.
	GetStatus() uint16
}
