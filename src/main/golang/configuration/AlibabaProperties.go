package configuration

type AlibabaProperties struct {
	Appid    string   `yaml:"appid"`
	Secret   string   `yaml:"secret"`
	Callback string   `yaml:"callback"`
	Url      string   `yaml:"base-url"`
	Scope    []string `yaml:"scopes"`
}

func (receiver AlibabaProperties) AppId() string {
	return receiver.Appid
}

func (receiver AlibabaProperties) AppSecret() string {
	return receiver.Secret
}

func (receiver AlibabaProperties) RedirectUrl() string {
	return receiver.Callback
}

func (receiver AlibabaProperties) BaseUrl() string {
	return prefix(receiver.Url)
}

func (receiver AlibabaProperties) Scopes() []string {
	return receiver.Scope
}
