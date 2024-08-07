package golang

import (
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
	"strings"
)

type baiduAuth struct {
	properties configuration.CloudProperties
}

func (receiver baiduAuth) GetAuthorizePage() string {

	uri := "oauth/2.0/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&device_id=%s"

	scope := strings.Join(receiver.properties.Scopes(), ",")

	uri = fmt.Sprintf(
		uri,
		receiver.properties.AppKey(),
		receiver.properties.RedirectUrl(),
		scope,
		receiver.properties.AppId(),
	)

	return receiver.properties.BaseUrl() + uri
}

func (receiver baiduAuth) Authorize(code string) (collection.Map[string, string], lang.Throwable) {
	return nil, nil
}
