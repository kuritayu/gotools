/*
Copyright © 2020 Yusuke.Kurita

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"

	goexcel "github.com/kuritayu/gotools/goexcel/lib"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
// TODO 引数が1ならシート一覧、2ならシートダンプ、3ならセルが良いか。
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Excelから各種情報を取得します。",
	Long: `Excelから各種情報を取得します。
使用例: 
   1. シート一覧出力
   goexcel get -s Book1.xlsx
 
   2. シートダンプ
   goexcel get -d Book1.xlsx Sheet1
 
   3. セル値出力
   goexcel get -c Book1.xlsx Sheet1 A1

   `,
	RunE: func(cmd *cobra.Command, args []string) error {
		sheetConfig, err := cmd.Flags().GetString("sheet")
		if err := callPrintSheetName(sheetConfig, err); err != nil {
			return err
		}

		cellConfig, err := cmd.Flags().GetStringSlice("cell")
		if err := callPrintValue(cellConfig, err, args); err != nil {
			return err
		}

		dumpConfig, err := cmd.Flags().GetStringSlice("dump")
		if err := callDump(dumpConfig, err, args); err != nil {
			return err
		}

		if sheetConfig == "" && len(cellConfig) == 0 && len(dumpConfig) == 0 {
			return errors.New("リソースが指定されていません。")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("sheet", "s", "", "引数で指定されたExcelのシート一覧を出力します。")
	getCmd.Flags().StringSliceP("cell", "c", []string{}, "引数で指定されたExcel、シート名、セル名の値を出力します。")
	getCmd.Flags().StringSliceP("dump", "d", []string{}, "引数で指定されたExcel、シートをダンプ(全出力)します。")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func callPrintSheetName(book string, err error) error {
	if book == "" {
		return nil
	}
	if err != nil {
		return err
	}
	f, err := goexcel.Load(book)
	if err != nil {
		return err
	}
	goexcel.PrintSheetName(f)
	return nil
}

func callPrintValue(config []string, err error, args []string) error {
	if len(config) == 0 {
		return nil
	}
	if err != nil {
		return err
	}
	book := config[0]

	if len(args) != 2 {
		return errors.New("シート名またはセル名が指定されていません。")
	}
	sheet := args[0]
	cell := args[1]
	if book != "" && sheet != "" && cell != "" {
		f, err := goexcel.Load(book)
		if err != nil {
			return err
		}
		goexcel.PrintValue(f, sheet, cell)
		return nil
	}
	return nil
}

func callDump(config []string, err error, args []string) error {
	if len(config) == 0 {
		return nil
	}
	if err != nil {
		return err
	}
	book := config[0]

	if len(args) != 2 {
		return errors.New("シート名またはセパレータ文字列が指定されていません。")
	}
	sheet := args[0]
	separator := args[1]
	if book != "" && sheet != "" && separator != "" {
		f, err := goexcel.Load(book)
		if err != nil {
			return err
		}
		goexcel.Dump(f, sheet, separator)
		return nil
	}
	return nil
}
