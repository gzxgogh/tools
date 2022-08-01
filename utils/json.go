package utils

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJSON(o interface{}) string {
	j, err := json.Marshal(o)
	if err != nil {
		return "{}"
	} else {
		js := string(j)
		js = strings.Replace(js, "\\u003c", "<", -1)
		js = strings.Replace(js, "\\u003e", ">", -1)
		js = strings.Replace(js, "\\u0026", "&", -1)
		return js
	}
}

func FromJSON(j string, o interface{}) (*interface{}, error) {
	err := json.Unmarshal([]byte(j), &o)
	if err != nil {
		return nil, err
	} else {
		return &o, nil
	}
}
