package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"io"
	"net/http"
	"strings"
)

type alibabaAuth struct {
	properties configuration.ApplicationProperties
}

func (auth alibabaAuth) GetAuthorizePage() string {

	uri := "oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s"

	uri = fmt.Sprintf(
		uri,
		auth.properties.AppId(),
		auth.properties.RedirectUrl(),
		strings.Join(auth.properties.Scopes(), ","),
	)

	return auth.properties.BaseUrl() + uri
}

func (auth alibabaAuth) Authorize(code string) (collection.Map[string, string], Throwable) {

	uri := auth.properties.BaseUrl() + "oauth/access_token"

	body := collection.NewHashMap[string, string]()

	body.Put("client_id", auth.properties.AppId()).
		Put("client_secret", auth.properties.AppSecret()).
		Put("grant_type", "authorization_code").
		Put("code", code)

	jsonBytes, _ := json.Marshal(body)

	response, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonBytes))

	defer response.Body.Close()

	if err != nil {
		return collection.NewHashMap[string, string](), throwNewRuntimeException(err.Error(), Error)
	}

	result, _ := io.ReadAll(response.Body)

	resp := collection.NewHashMap[string, string]()

	err = json.Unmarshal(result, &resp.Values)

	if err != nil {
		return nil, throwNewRuntimeException(err.Error(), Error)
	}

	return resp, throwNewException("", Success)
}
