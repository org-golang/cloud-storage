package golang

import (
	"github.com/org-golang/cloud-storage/src/main/golang/configuration"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestProperties(t *testing.T) {

	file, err := os.ReadFile("./cloud.yaml")

	if err != nil {
		t.Logf("Open file failed, err:%v\n", err)
	}

	var properties configuration.Properties

	err = yaml.Unmarshal(file, &properties)

	if err != nil {
		t.Errorf("Unable resolve file,err: %s", err.Error())
	}

	if len(properties.Provider.Baidu.Scope) <= 0 {
		t.Errorf("provider object is null: %v", properties)
	}
}
