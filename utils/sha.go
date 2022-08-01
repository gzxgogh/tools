package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func Sha1(src string) string {
	h := sha1.New()
	h.Write([]byte(src))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func Sha256(src string) string {
	sha256Hash := sha256.New()
	data := []byte(src)
	sha256Hash.Write(data)
	hashed := sha256Hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

func FileSha256(filename string) (string, error) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	sha256Hash := sha256.New()
	sha256Hash.Write(src)
	hashed := sha256Hash.Sum(nil)
	return fmt.Sprintf("%x", hashed), nil
}
