package goexcel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// PrintValue セルの値を出力します。
func PrintValue(f *excelize.File, name string, axis string) error {
	cell, err := f.GetCellValue(name, axis)
	if err != nil {
		return err
	}
	fmt.Println(cell)
	return nil
}

// Dump シートを全出力します。
func Dump(f *excelize.File, name string, separator string) error {
	rows, err := f.GetRows(name)
	if err != nil {
		return err
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, separator)
		}
		fmt.Println()
	}
	return nil
}

// PrintSheetName シート一覧を出力します。
func PrintSheetName(f *excelize.File) error {
	for _, name := range f.GetSheetList() {
		fmt.Println(name)
	}
	return nil
}
