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

// TODO saveメソッドを再定義して、先頭シートのA1を選択させるようにする
