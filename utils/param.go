package utils

import (
	"errors"
	"fmt"
	"strings"
)

// 获取分页信息,返回PageSize pageNo offset 及 mysql 分页组合
func ParamGetPageInfoSql(pageNo, pageSize int) (int, string) {
	offset := (pageNo - 1) * pageSize
	limitSql := ""
	if pageNo > 0 && pageSize > 0 {
		limitSql = fmt.Sprintf(`limit %d,%d`, offset, pageSize)
	}
	return offset, limitSql
}

//验证必要参数
func CheckRequiredParam(param map[string]interface{}, name ...string) {
	var missParamName string
	var flag bool
	for _, item := range name {
		if "" == fmt.Sprint(param[item]) || nil == param[item] {
			flag = true
			missParamName = item
			break
		}
	}
	if flag {
		err := errors.New("缺少参数名为" + missParamName + "的数据")
		panic(err)
	}
}

//复制元素到新的map
func CopyToMap(param map[string]interface{}, paramName ...string) map[string]interface{} {
	newMap := make(map[string]interface{})
	if nil != param {
		for _, item := range paramName {
			value, ok := param[item]
			if ok {
				newMap[item] = value
			}
		}
	}
	return newMap
}

// orimap 提供元素的map ; objmap 要插入元素的map, elenames 元素名字列表
func CopyMapFromOri(oriMap map[string]interface{}, objMap map[string]interface{}, names ...string) {
	if nil == oriMap || nil == objMap {
		return
	}
	for _, ele := range names {
		ele = strings.TrimSpace(ele)
		if "" == ele {
			continue
		}
		if nil == oriMap[ele] {
			continue
		}
		objMap[ele] = strings.TrimSpace(fmt.Sprint(oriMap[ele]))
	}
}
