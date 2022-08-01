package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mathRand "math/rand"
	"time"
)

//32位
func UUID() string {
	return getGuid()
}

//50位
func OriUUID() string {
	b := getGuid()
	uuid := fmt.Sprintf("%v-%v-%v-%v-%v", b[0:8], b[8:12], b[12:16], b[16:20], b[20:])
	return uuid
}

//生成32位md5字串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func getGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	return GenerateRandString(str, l)
}

func GetRandomCaseString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&_="
	return GenerateRandString(str, l)
}

func GetRandomHexString(l int) string {
	str := "0123456789abcdef"
	return GenerateRandString(str, l)
}

func GetRandomIntString(l int) string {
	str := "0123456789"
	return GenerateRandString(str, l)
}

func GenerateRandString(source string, l int) string {
	bytes := []byte(source)
	result := []byte{}
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
