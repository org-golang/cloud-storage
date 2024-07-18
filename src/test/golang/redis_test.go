package golang

import (
	"context"
	"fmt"
	"github.com/org-golang/cloud-storage/src/main/golang/collection"
	"github.com/org-golang/cloud-storage/src/main/golang/lang"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestRedis(t *testing.T) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	fmt.Println(rdb)

	ctx := context.Background()

	hashMap := collection.NewHashMap[string, any]()

	hashMap.Put("token_type", "Bearer")
	hashMap.Put("access_token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ")
	hashMap.Put("refresh_token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ")
	hashMap.Put("expires_in", 7200)

	val, err := rdb.HSet(
		ctx,
		"user:1",
		"token_type", "Bearer",
		"access_token",
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9",
		"refresh_token",
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9",
		"expires_in",
		7200,
	).Result()

	if lang.IsNonNil(err) {
		t.Errorf("Store error: %s \n", err.Error())
	}
	fmt.Println(val, err)
}
