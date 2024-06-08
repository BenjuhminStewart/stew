/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/BenjuhminStewart/stew/types"
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stew from a given name or id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// check id flag
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")

		s := types.Stews{}
		err := s.Load(types.StewPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		if id == -1 && name == "" {
			cmd.Help()
			return
		}

		if id != -1 && name != "" {
			fmt.Println("You can only use one flag at a time")
			return
		}

		if id != -1 {
			_, err := s.Get(id)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if name != "" {
			_, err := s.GetByName(name)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func flags() {
	GetCmd.Flags().IntP("id", "i", -1, "The id of the stew")
	GetCmd.Flags().StringP("name", "n", "", "The name of the stew")
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	flags()
}
