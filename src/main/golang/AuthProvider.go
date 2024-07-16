package golang

import "github.com/org-golang/cloud-storage/src/main/golang/collection"

type AuthProvider interface {

	// GetAuthorizePage Get user authorize page.
	GetAuthorizePage() string

	// Authorize User authorize action.
	Authorize(code string) (collection.Map[string, string], Throwable)
}
