/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package add

import (
	"github.com/spf13/cobra"
)

// AddCmd represents the create command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a stew based off a defined directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
