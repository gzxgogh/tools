package utils

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//保留小数位数
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

//保留四位小数
func RoundInterface4(val interface{}) float64 {
	f := Math2float64(val)
	return Round(f, 4)
}

//保留两位小数
func RoundInterface2(val interface{}) float64 {
	f := Math2float64(val)
	return Round(f, 2)
}

func RoundMap4(mp map[string]interface{}, paramName ...string) {
	if nil == mp {
		return
	}
	for _, key := range paramName {
		if nil == mp[key] {
			continue
		}
		if "" == fmt.Sprint(mp[key]) {
			continue
		}
		mp[key] = RoundInterface4(mp[key])
	}
}

func RoundMap2(mp map[string]interface{}, paramName ...string) {
	if nil == mp {
		return
	}
	for _, key := range paramName {
		if nil == mp[key] {
			continue
		}

		if "" == fmt.Sprint(mp[key]) {
			continue
		}

		mp[key] = RoundInterface2(mp[key])
	}
}

func RoundMapLst4(mpLst []map[string]interface{}, paramName ...string) {
	if len(mpLst) <= 0 {
		return
	}
	for _, mp := range mpLst {
		if nil == mp {
			continue
		}

		for _, key := range paramName {
			if nil == mp[key] {
				continue
			}

			if "" == fmt.Sprint(mp[key]) {
				continue
			}

			mp[key] = RoundInterface4(mp[key])
		}
	}
}

func RoundMapLst2(mpLst []map[string]interface{}, paramName ...string) {
	if len(mpLst) <= 0 {
		return
	}
	for _, mp := range mpLst {
		if nil == mp {
			continue
		}

		for _, key := range paramName {
			if nil == mp[key] {
				continue
			}

			if "" == fmt.Sprint(mp[key]) {
				continue
			}

			mp[key] = RoundInterface2(mp[key])
		}
	}
}

//interface 转float64
func Math2float64(param interface{}) float64 {
	if nil == param {
		return 0
	}
	str := fmt.Sprint(param)
	str = strings.Replace(str, ",", "", -1)
	if "" == param {
		return 0
	}
	fre, err := strconv.ParseFloat(str, 64)
	if nil != err {
		return 0
	}
	return fre
}

func Math2int(param interface{}) int {
	fre := Math2float64(param)
	return int(fre)
}

// 字符串递增,长度一致
func StrIncrease(strNo string) string {
	intNo := Math2float64(strNo)
	intNo = intNo + 1
	strNewNo := fmt.Sprint(intNo)
	if len(strNewNo) >= len(strNo) {
		return strNewNo
	}
	zNum := len(strNo) - len(strNewNo)
	for i := 0; i < zNum; i++ {
		strNewNo = "0" + strNewNo
	}

	return strNewNo
}

//数字转换成中文大写
func ConvertNumToCny(num float64) string {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	// log.Println(sliceUnit[:len(sliceUnit)-2])
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]

	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}
	reg, err := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, err = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, err = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")

	reg, err = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		log.Fatal(err)
	}
	return str
}

//数字转换为百分数
func ConvertNumToPercent(num float64) string {
	return fmt.Sprint(Round(num*100, 4)) + "%"
}
