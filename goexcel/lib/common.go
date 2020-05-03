package goexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// Load Excelファイルをロードします。
func Load(name string) (*excelize.File, error) {
	f, err := excelize.OpenFile(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
