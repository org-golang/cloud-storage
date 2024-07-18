package configuration

import (
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
	"gopkg.in/yaml.v3"
	"os"
)

type Properties struct {
	Provider Provider `yaml:"providers"`
}

type Provider struct {
	Alibaba AlibabaProperties `yaml:"alibaba"`
	Baidu   BaiduProperties   `yaml:"baidu"`
}

func BuildProperties(filePath string) Properties {

	file, err := os.ReadFile(filePath)

	if lang.IsNonNil(err) {
		panic(err)
	}

	var properties Properties

	err = yaml.Unmarshal(file, &properties)

	if lang.IsNonNil(err) {
		panic(err)
	}

	return properties
}
