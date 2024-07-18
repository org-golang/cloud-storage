package golang

type Token struct {
	TokenType string

	AccessToken string

	RefreshToken string

	ExpiresIn uint16
}
