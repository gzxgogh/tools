package utils

import (
	"fmt"
	"time"
)

func Now() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

func NowDate() string {
	return time.Now().Local().Format("2006-01-02")
}

func Timestamp() int64 {
	return time.Now().Unix()
}

func NowNum() string {
	return time.Now().Local().Format("20060102150405")
}

func NowYear() int {
	return time.Now().Year()
}

func NowMonth() string {
	m := fmt.Sprint(time.Now().Month())
	month := ""
	switch m {
	case "January":
		month = "01"
		break
	case "February":
		month = "02"
		break
	case "March":
		month = "03"
		break
	case "April":
		month = "04"
		break
	case "May":
		month = "05"
		break
	case "June":
		month = "06"
		break
	case "July":
		month = "07"
		break
	case "August":
		month = "08"
		break
	case "September":
		month = "09"
		break
	case "October":
		month = "10"
		break
	case "November":
		month = "11"
		break
	case "December":
		month = "12"
		break
	default:
		month = "不存在的"
	}
	return month
}

func GetMonthDays(year, month int) (num int) {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		num = 30
	} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		num = 31
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			num = 29
		} else {
			num = 28
		}
	}

	return num
}
