package goexcel

import (
	"errors"
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
	// FIXME 重複コード
	if f.GetSheetIndex(name) == -1 {
		return errors.New("シートが存在しません。")
	}
	f.DeleteSheet(name)
	return nil
}
