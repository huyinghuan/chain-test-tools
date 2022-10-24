package utils

import (
	"encoding/json"
	mrand "math/rand"
)

func BaseJsonEncode(data interface{}) string {
	mjson, _ := json.Marshal(data)
	mString := string(mjson)
	return mString
}

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// RandString
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[mrand.Intn(len(chars))])
	}
	return string(bs)
}
