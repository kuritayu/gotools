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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Excelの各種情報を更新します。",
	Long: `Excelの各種情報を更新します。
使用例: 
	1. シート名更新(Sheet1をSheet2更新)
	goexcel set Book1.xlsx Sheet1 Sheet2
  
	2. セル値更新(Sheet1のA1を100に更新)
	goexcel set Book1.xlsx Sheet1 A1 100
 
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 && len(args) != 4 {
			return errors.New("引数が不正です。")
		}
		f, err := goexcel.Load(args[0])
		if err != nil {
			return err
		}
		switch len(args) {
		case 3:
			if err := goexcel.RenameSheet(f, args[1], args[2]); err != nil {
				return err
			}
		case 4:
			if err := goexcel.SetValue(f, args[1], args[2], args[3]); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
