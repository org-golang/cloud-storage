package configuration

import (
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

	if err != nil {
		panic(err)
	}

	var properties Properties

	err = yaml.Unmarshal(file, &properties)

	if err != nil {
		panic(err)
	}

	return properties
}
