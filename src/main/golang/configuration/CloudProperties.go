package configuration

type CloudProperties interface {
	ApplicationProperties
	// Display Auth page show style.
	Display() string

	// AppKey Client id.
	AppKey() string

	// SignKey Signature key.
	SignKey() string
}
