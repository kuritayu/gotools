package goexcel

import (
	"errors"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// SetValue セルに値を設定します。
func SetValue(f *excelize.File, name string, axis string, value string) error {
	if err := f.SetCellDefault(name, axis, value); err != nil {
		return err
	}
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}

// RenameSheet シート名をリネームします。
func RenameSheet(f *excelize.File, old string, new string) error {
	if f.GetSheetIndex(old) == -1 {
		return errors.New("シートが存在しません。")
	}
	f.SetSheetName(old, new)
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
