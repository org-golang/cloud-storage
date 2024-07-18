package golang

import (
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
)

type AuthProvider interface {

	// GetAuthorizePage Get user authorize page
	GetAuthorizePage() string

	// Authorize User authorize action.
	Authorize(code string) (collection.Map[string, string], lang.Throwable)
}

// NewAuthProvider Create a new auth provider instance.
func NewAuthProvider(configuration configuration.Properties, scene string) (AuthProvider, lang.Throwable) {

	if scene != ALIBABA && scene != BAIDU {
		return nil, lang.ThrowNewException(fmt.Sprintf("Sence must be %s or %s ", ALIBABA, BAIDU), lang.Error)
	}

	if scene == ALIBABA {
		return &alibabaAuth{properties: configuration.Provider.Alibaba}, nil
	}

	return &baiduAuth{properties: configuration.Provider.Baidu}, nil
}

const (
	ALIBABA = "alibaba"
	BAIDU   = "baidu"
)
