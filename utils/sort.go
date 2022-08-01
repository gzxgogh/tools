package utils

import "fmt"

func SortLstDesc(sortLst []map[string]interface{}, sortField string) {
	temp := make(map[string]interface{})
	length := len(sortLst)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if fmt.Sprint(sortLst[j][sortField]) < fmt.Sprint(sortLst[j+1][sortField]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortLstAsc(sortLst []map[string]interface{}, sortField string) {
	length := len(sortLst)
	temp := make(map[string]interface{})
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if fmt.Sprint(sortLst[j][sortField]) > fmt.Sprint(sortLst[j+1][sortField]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortNumLstDesc(sortLst []map[string]interface{}, sortField string) {
	length := len(sortLst)
	temp := make(map[string]interface{})
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if Math2float64(sortLst[j][sortField]) < Math2float64(sortLst[j+1][sortField]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortNumLstAsc(sortLst []map[string]interface{}, sortField string) {
	length := len(sortLst)
	temp := make(map[string]interface{})
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if Math2float64(sortLst[j][sortField]) > Math2float64(sortLst[j+1][sortField]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortArrDesc(sortLst []string) {
	length := len(sortLst)
	temp := ""
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if fmt.Sprint(sortLst[j]) < fmt.Sprint(sortLst[j+1]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortArrAsc(sortLst []string) {
	length := len(sortLst)
	temp := ""
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if fmt.Sprint(sortLst[j]) > fmt.Sprint(sortLst[j+1]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}

func SortNumArrAsc(sortLst []string) {
	length := len(sortLst)
	temp := ""
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if Math2float64(sortLst[j]) > Math2float64(sortLst[j+1]) {
				temp = sortLst[j]
				sortLst[j] = sortLst[j+1]
				sortLst[j+1] = temp
			}
		}
	}
}
