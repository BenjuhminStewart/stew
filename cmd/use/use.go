/*
Copyright © 2024 NAME HERE benjuhminstewart@gmail.com
*/
package use

import (
	"github.com/spf13/cobra"
)

// UseCmd represents the init command
var UseCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a stew instance to create a new project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
