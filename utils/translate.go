package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Translate(oriLang, aimLang, content string) (string, error) {
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s", oriLang, aimLang, url.QueryEscape(content))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str := string(body)
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, "null,", "")
	str = strings.ReplaceAll(str, "\\u200b", "")
	str = strings.Trim(str, `"`)
	arr := strings.Split(str, `","`)
	str = arr[0]
	if strings.Contains(arr[0], "href=//www.google.com") {
		return "", errors.New("无效的语言类型")
	}

	return str, nil
}
