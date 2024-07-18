package golang

import (
	"reflect"
)

type Response map[string]any

func Json[T string | ~map[string]any](data T, code int) Response {

	r := make(Response)
	r["code"] = code
	r["data"] = nil
	r["msg"] = "success"

	reflector := reflect.TypeOf(data)

	var key string

	if reflector.Kind() == reflect.String {
		key = "msg"
	} else {
		key = "data"
	}

	r[key] = data

	return r
}
