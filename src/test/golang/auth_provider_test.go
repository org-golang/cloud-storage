package golang

import (
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
	"strings"
	"testing"
)

var properties configuration.Properties

func init() {
	properties = configuration.BuildProperties("./cloud.yaml")
}

func TestCreateAuthorizePageUrl(t *testing.T) {

	alibaba, err := golang.NewAuthProvider(properties, golang.ALIBABA)

	if lang.IsNonNil(err) {
		t.Error(err.GetMessage())
	}

	if !strings.Contains(alibaba.GetAuthorizePage(), "openapi.alipan.com") {
		t.Errorf("Not included: [%s]", "openapi.alipan.com")
	}

	baidu, err := golang.NewAuthProvider(properties, golang.BAIDU)

	if lang.IsNonNil(err) {
		t.Error(err.GetMessage())
	}

	if !strings.Contains(baidu.GetAuthorizePage(), "openapi.baidu.com") {
		t.Errorf("Not included: [%s]", "openapi.baidu.com")
	}

	fmt.Println(alibaba.GetAuthorizePage())
	fmt.Println(baidu.GetAuthorizePage())

}

func TestGetAccessToken(t *testing.T) {

	alibaba, err := golang.NewAuthProvider(properties, golang.ALIBABA)

	if lang.IsNonNil(err) {
		t.Error(err.GetMessage())
	}

	res, err := alibaba.Authorize("cf73b577bb7f4761a2a0b2c5d0d5acd8")

	if lang.IsNonNil(err) {
		t.Error(err.GetMessage())
	}

	fmt.Println(res)
}
