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
	goaddress "github.com/kuritayu/gotools/goaddress/lib"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "郵便番号データ検索コマンド",
	Long: `郵便番号データから郵便番号または住所を検索します。

example:
	goaddress search 1810002
	goaddress search -r 東京都三鷹市牟礼
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := cmd.Flags().GetString("reverse")
		if err != nil {
			return err
		}
		if s != "" {
			err := goaddress.AddToCode(s)
			if err != nil {
				return err
			}
			return nil
		}

		err = goaddress.CodeToAdd(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("reverse", "r", "", "逆引き(住所->郵便番号)オプション")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
