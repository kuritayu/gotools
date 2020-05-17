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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goaddress",
	Short: "郵便番号 <-> 住所変換ツール",
	Long: `郵便番号から住所変換、または住所から郵便番号に変換を行うツールです。

このツールは、日本郵便の郵便番号データを使用して検索を行います。
郵便番号データは日々更新されますので、使用時は更新を行ってください。
インターネット接続環境であれば更新コマンドを実行してください。
非インターネット接続環境であれば、
ダウンロードした郵便番号データをカレントディレクトリに格納してください。
<https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/ken_all.zip>
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goaddress.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}