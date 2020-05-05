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

// IsExistedSheet シートが存在しているか確認します。
func IsExistedSheet(f *excelize.File, name string) error {
	if f.GetSheetIndex(name) == -1 {
		return excelize.ErrSheetNotExist{
			SheetName: name,
		}
	}
	return nil
}
