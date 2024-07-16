package golang

import (
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
)

type AuthProvider interface {

	// GetAuthorizePage Get user authorize page.
	GetAuthorizePage() string

	// Authorize User authorize action.
	Authorize(code string) (collection.Map[string, string], Throwable)
}

// NewAuthProvider Create a new auth provider instance.
func NewAuthProvider(configuration configuration.Properties, scene string) (AuthProvider, Throwable) {

	if scene != ALIBABA && scene != BAIDU {
		return nil,
			throwNewRuntimeException(fmt.Sprintf("sence must be %s or %s ", ALIBABA, BAIDU), Error)
	}

	if scene == ALIBABA {
		return &alibabaAuth{properties: configuration.Provider.Alibaba}, throwNewException("", Success)
	}

	return &baiduAuth{properties: configuration.Provider.Baidu}, throwNewException("", Success)
}

const (
	ALIBABA = "alibaba"
	BAIDU   = "baidu"
)
