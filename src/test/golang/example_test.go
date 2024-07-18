package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestExample(t *testing.T) {
	fmt.Println("Hello World")

	authUrl := "https://openapi.alipan.com"

	accessToken := "/oauth/access_token"

	var body = make(map[string]string)

	body["client_id"] = "508dcedc4b6e41c1b120355ae7b8c550"
	body["client_secret"] = "f122bc6a7b2a48c79c81f77eaa3ac49c"
	body["grant_type"] = "authorization_code"
	body["code"] = "cf73b577bb7f4761a2a0b2c5d0d5acd8"

	jsonData, _ := json.Marshal(body)

	resp, err := http.Post(authUrl+accessToken, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	if resp.StatusCode != 200 {
		panic("response status code: " + resp.Status)
	}

	result, _ := io.ReadAll(resp.Body)
	fmt.Println(string(result))

}
