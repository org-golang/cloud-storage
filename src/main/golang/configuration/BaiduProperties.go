package configuration

type BaiduProperties struct {
	Appid string `json:"appid"`

	SecretKey string `json:"secret-key"`

	React string `yaml:"display"`

	ApplicationKey string `yaml:"app-key"`

	Signature string `yaml:"sign-key"`

	Scope []string `yaml:"scopes"`

	Callback string `yaml:"callback"`

	Url string `yaml:"base-url"`
}

func (receiver BaiduProperties) Display() string {
	return receiver.React
}

func (receiver BaiduProperties) AppId() string {
	return receiver.Appid
}

func (receiver BaiduProperties) AppSecret() string {
	return receiver.SecretKey
}

func (receiver BaiduProperties) RedirectUrl() string {
	return receiver.Callback
}

func (receiver BaiduProperties) BaseUrl() string {
	return prefix(receiver.Url)
}

func (receiver BaiduProperties) Scopes() []string {
	return receiver.Scope
}

func (receiver BaiduProperties) AppKey() string {
	return receiver.ApplicationKey
}

func (receiver BaiduProperties) SignKey() string {
	return receiver.Signature
}
