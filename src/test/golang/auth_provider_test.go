package golang

import (
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {

	properties := configuration.BuildProperties("./cloud.yaml")

	alibaba, err := golang.NewAuthProvider(properties, golang.ALIBABA)

	if err.Error() {
		t.Error(err.GetMessage())
	}

	if !strings.Contains(alibaba.GetAuthorizePage(), "openapi.alipan.com") {
		t.Errorf("Not included: [%s]", "openapi.alipan.com")
	}

	baidu, err := golang.NewAuthProvider(properties, golang.BAIDU)

	if err.Error() {
		t.Error(err.GetMessage())
	}

	if !strings.Contains(baidu.GetAuthorizePage(), "openapi.baidu.com") {
		t.Errorf("Not included: [%s]", "openapi.baidu.com")
	}

	fmt.Println(alibaba.GetAuthorizePage())
	fmt.Println(baidu.GetAuthorizePage())
}
