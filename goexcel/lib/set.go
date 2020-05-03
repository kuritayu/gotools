package goexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// SetValue セルに値を設定します。
func SetValue(f *excelize.File, name string, axis string, value interface{}) error {
	if err := f.SetCellValue(name, axis, value); err != nil {
		return err
	}
	return nil
}
