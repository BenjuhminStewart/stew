/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all created stews",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
