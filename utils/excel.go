package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

//读取excel的内容
func GetExcelMapLst(startRow int, filePath string, colCaptionMap map[string]string) ([]map[string]interface{}, error) {
	isXlsx := strings.HasSuffix(filePath, "xlsx")
	isXls := strings.HasSuffix(filePath, "xls")
	if isXlsx != true && isXls != true {
		return nil, errors.New("错误的文件格式")
	}
	xlFile, err := xlsx.OpenFile(filePath)
	if fmt.Sprint(err) == `strconv.ParseInt: parsing "true": invalid syntax` {
		return nil, errors.New("请将单元格格式重新进行保存")
	}
	if err != nil {
		return nil, err
	}
	var finalList []map[string]interface{}

	// 英文-列索引
	colIndexMap := make(map[string]interface{})
	// 不存在的列名
	var unExistColumnName []string
	unExistColNameExistMap := make(map[string]bool)

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex < startRow {
				continue
			} else if startRow == rowIndex {
				// 读到的 标题-索引 map
				captionIndexMap := make(map[string]interface{})
				for colIndex, cell := range row.Cells {
					cst := cell.String()
					// 去空格后的单元格内容
					value := strings.TrimSpace(cst)
					if strings.HasSuffix(value, `*`) {
						value = value[:len(value)-1]
					}
					captionIndexMap[value] = colIndex
				}
				//  colCaptionMap  为 属性名英文-属性中文
				for k, v := range colCaptionMap {
					if strings.HasSuffix(v, `*`) {
						v = v[:len(v)-1]
					}
					colIndexMap[k] = captionIndexMap[v]
				}
			} else {
				objMap := make(map[string]interface{})

				// 判断是否为空，若整行为空，则不再读取
				bIsEmpty := true

				for k, interV := range colIndexMap {
					if nil == interV {
						// 若值为空，说明不存在该列，过
						if !unExistColNameExistMap[colCaptionMap[k]] {
							unExistColumnName = append(unExistColumnName, colCaptionMap[k])
							unExistColNameExistMap[colCaptionMap[k]] = true
						}
						continue
					}
					v, _ := strconv.Atoi(fmt.Sprint(interV))
					if len(row.Cells) < (v + 1) {
						continue
					}
					cel := row.Cells[v]
					if nil == cel {
						continue
					}
					str := cel.String()

					if strings.Contains(str, "%") {
						str = strings.Replace(str, "%", " ", -1)
					}
					celStr := strings.TrimSpace(str)
					objMap[k] = celStr
					if "" != celStr {
						bIsEmpty = false
					}
				}
				if bIsEmpty {
					break
				}

				finalList = append(finalList, objMap)
			}
		}
	}

	return finalList, nil
}

//生成excel
func MakeExcelFile(filePath string, colCaptionMap map[string]string, content []map[string]interface{}) error {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return err
	}

	// 添加第一行 说明
	row0 := sheet.AddRow()
	var keys []string
	for k, _ := range colCaptionMap {
		keys = append(keys, k)
		cell := row0.AddCell()
		cell.Value = k
	}

	kLen := len(keys)
	for _, v := range content {
		tmpRow := sheet.AddRow()
		for i := 0; i < kLen; i++ {
			tmCell := tmpRow.AddCell()
			tmCell.Value = fmt.Sprint(v[colCaptionMap[keys[i]]])
		}
	}
	err = file.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}
