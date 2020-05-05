package goexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

// CreateEmptyFile 空ファイルを作成します。
func CreateEmptyFile(name string) error {
	f := excelize.NewFile()
	if err := f.SaveAs(name); err != nil {
		return err
	}
	return nil
}

// CreateSheet シートを作成します。
func CreateSheet(f *excelize.File, name string) error {
	_ = f.NewSheet(name)
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
