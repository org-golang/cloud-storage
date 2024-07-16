package configuration

type Properties struct {
	Provider Provider `yaml:"providers"`
}

type Provider struct {
	Alibaba AlibabaProperties `yaml:"alibaba"`
	Baidu   BaiduProperties   `yaml:"baidu"`
}
