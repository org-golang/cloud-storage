package golang

import (
	"github.com/org-golang/cloud-storage/src/main/golang"
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {

	properties := configuration.BuildProperties("./cloud.yaml")

	alibaba := golang.NewAlibabaAuth(properties)

	if !strings.Contains(alibaba.GetAuthorizePage(), "openapi.alipan.com") {
		t.Errorf("Not included: [%s]", "openapi.alipan.com")
	}

	baidu := golang.NewBaiduAuth(properties)

	if !strings.Contains(baidu.GetAuthorizePage(), "openapi.baidu.com") {
		t.Errorf("Not included: [%s]", "openapi.baidu.com")
	}
}
