package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 encodes string to MD5 hex value.
func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
