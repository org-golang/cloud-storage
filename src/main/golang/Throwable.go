package golang

type Throwable interface {

	// GetMessage Get exception error message.
	GetMessage() string

	// GetStatus Get exception error code.
	GetStatus() uint16

	// Error Determine if an exception is thrown.
	Error() bool
}
