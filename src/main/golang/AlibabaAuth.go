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

type AlibabaAuth struct {
	properties configuration.ApplicationProperties
}

func (auth AlibabaAuth) GetAuthorizePage() string {

	uri := "oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s"

	uri = fmt.Sprintf(
		uri,
		auth.properties.AppId(),
		auth.properties.RedirectUrl(),
		strings.Join(auth.properties.Scopes(), ","),
	)

	return auth.properties.BaseUrl() + uri
}

func (auth AlibabaAuth) Authorize(code string) (collection.Map[string, string], Throwable) {

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
		return collection.NewHashMap[string, string](), throwNewException(err.Error(), Error).throw(true)
	}

	result, _ := io.ReadAll(response.Body)

	resp := collection.NewHashMap[string, string]()

	err = json.Unmarshal(result, &resp.Values)

	if err != nil {
		return nil, throwNewException(err.Error(), Error).throw(true)
	}

	return resp, throwNewException("", Success)
}
