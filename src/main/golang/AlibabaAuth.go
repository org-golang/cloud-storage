package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
	"io"
	"net/http"
	"net/url"
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
		url.QueryEscape(auth.properties.RedirectUrl()),
		strings.Join(auth.properties.Scopes(), ","),
	)

	return auth.properties.BaseUrl() + uri
}

func (auth alibabaAuth) Authorize(code string) (collection.Map[string, string], lang.Throwable) {

	uri := auth.properties.BaseUrl() + "oauth/access_token"

	body := collection.NewHashMap[string, string]()

	body.Put("client_id", auth.properties.AppId()).
		Put("client_secret", auth.properties.AppSecret()).
		Put("grant_type", "authorization_code").
		Put("code", code)

	jsonBytes, _ := json.Marshal(body.Values)

	response, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonBytes))

	defer response.Body.Close()

	if lang.IsNonNil(err) {
		return nil, lang.ThrowNewException(err.Error(), lang.Error)
	}

	result, _ := io.ReadAll(response.Body)

	resp := collection.NewHashMap[string, string]()

	err = json.Unmarshal(result, &resp.Values)

	if lang.IsNonNil(err) {
		return nil, lang.ThrowNewException(err.Error(), lang.Error)
	}

	if resp.Contains("code") {
		return nil, lang.ThrowNewException(resp.Get("code"), lang.Error)
	}

	return resp, nil
}

const (
	AppNotExists        = "应用不存在"
	InvalidClientSecret = "应用密钥不对"
	InvalidCode         = "授权码为空或过期"
	InvalidClientId     = "应用ID和构造授权链接时填的不一致"
	InvalidGrantType    = "无效的担保类型，目前仅支持 authorization_code 和 refresh_token"
)
