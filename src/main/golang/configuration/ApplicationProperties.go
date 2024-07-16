package configuration

import "strings"

type ApplicationProperties interface {

	// AppId Get developer app id.
	AppId() string

	// AppSecret Get developer app Secret.
	AppSecret() string

	// RedirectUrl Obtain the Callback address used to receive the authorization code.
	RedirectUrl() string

	// BaseUrl Base Url.
	BaseUrl() string

	// Scopes Authorize Scope.
	Scopes() []string
}

// Interface default implementation.
func prefix(haystack string) string {

	if strings.HasSuffix(haystack, "/") {
		return haystack
	}

	return haystack + "/"
}
