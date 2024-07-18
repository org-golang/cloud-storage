package golang

import "testing"

func TestEnum(t *testing.T) {

}

const (
	AppNotExists        = "应用不存在"
	InvalidClientSecret = "应用密钥不对"
	InvalidCode         = "授权码为空或过期"
	InvalidClientId     = "应用ID和构造授权链接时填的不一致"
	InvalidGrantType    = "无效的担保类型，目前仅支持 authorization_code 和 refresh_token"
)
