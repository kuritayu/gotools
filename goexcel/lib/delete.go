package goexcel

import (
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// DeleteBook ブックを削除します。
func DeleteBook(name string) error {
	if err := os.Remove(name); err != nil {
		return nil
	}
	return nil
}

// DeleteSheet シートを削除します。
func DeleteSheet(f *excelize.File, name string) error {
	if err := IsExistedSheet(f, name); err != nil {
		return err
	}
	f.DeleteSheet(name)
	return nil
}
