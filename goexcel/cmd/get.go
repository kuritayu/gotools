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
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Excelから各種情報を取得します。",
	Long: `Excelから各種情報を取得します。
使用例: 
   1. シート一覧出力
   goexcel get Book1.xlsx
 
   2. シートダンプ
   goexcel get Book1.xlsx Sheet1
 
   3. セル値出力
   goexcel get Book1.xlsx Sheet1 A1

   `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || len(args) > 3 {
			return errors.New("引数が不正です。")
		}
		f, err := goexcel.Load(args[0])
		if err != nil {
			return err
		}
		switch len(args) {
		case 1:
			if err := goexcel.PrintSheetName(f); err != nil {
				return err
			}
		case 2:
			if err := goexcel.Dump(f, args[1], ","); err != nil {
				return err
			}
		case 3:
			if err := goexcel.PrintValue(f, args[1], args[2]); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
