package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func CheckHttpSign(sign, appId, none, signKey string, timestamp int64) bool {
	str := fmt.Sprintf("appid=%s&none=%s&timestamp=%d&secret_key=%s", appId, none, timestamp, signKey)
	str = strings.ToLower(str)
	hash := md5.Sum([]byte(str))
	md := hex.EncodeToString(hash[:])
	return sign == md
}
