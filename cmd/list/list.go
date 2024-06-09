/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package list

import (
	"fmt"
	"os"

	"github.com/BenjuhminStewart/stew/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all created stews",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		st := types.Stews{}
		if err := st.Load(viper.GetString("stewsPath")); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		st.List()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
