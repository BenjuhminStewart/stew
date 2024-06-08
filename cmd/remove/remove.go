/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"github.com/spf13/cobra"
)

// RemoveCmd represents the delete command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a stew",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
