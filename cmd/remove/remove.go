/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package remove

import (
	"github.com/BenjuhminStewart/stew/types"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the delete command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a stew",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := types.Stews{}
		err := s.Load(types.GetHomeDir() + "/.stews.json")
		if err != nil {
			cmd.Println(err)
			return
		}

		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")

		if id == -1 && name == "" {
			cmd.Help()
			return
		}

		if id != -1 && name != "" {
			cmd.Println("Either search by id OR name, not both")
			return
		}

		if id != -1 {
			err := s.Remove(id)
			if err != nil {
				cmd.Println(err)
				return
			}
		}

		if name != "" {
			err := s.RemoveByName(name)
			if err != nil {
				cmd.Println(err)
				return
			}
		}

		err = s.Save(types.GetHomeDir() + "/.stews.json")
		if err != nil {
			cmd.Println(err)
			return
		}
	},
}

func flags() {
	RemoveCmd.Flags().IntP("id", "i", -1, "The id of the stew you want to remove")
	RemoveCmd.Flags().StringP("name", "n", "", "The name of the stew you want to remove")
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flags()
}
