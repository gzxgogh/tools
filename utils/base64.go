package utils

import (
	"encoding/base64"
)

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	out, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	} else {
		return string(out)
	}
}
