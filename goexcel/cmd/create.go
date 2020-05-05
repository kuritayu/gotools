/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Excelの各種情報を作成します。",
	Long: `Excelの各種情報を作成します。
使用例: 
	1. ブック作成
	goexcel create Book1.xlsx
  
	2. シート作成
	goexcel create Book1.xlsx Sheet1
 
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 && len(args) != 2 {
			return errors.New("引数が不正です。")
		}
		switch len(args) {
		case 1:
			if err := goexcel.CreateEmptyFile(args[0]); err != nil {
				return err
			}
		case 2:
			f, err := goexcel.Load(args[0])
			if err != nil {
				return err
			}
			if err := goexcel.CreateSheet(f, args[1]); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
