/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com>
*/
package edit

import (
	"github.com/spf13/cobra"

	"github.com/BenjuhminStewart/stew/types"
)

// EditCmd represents the edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a stew's name, description, or path",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		selected_stew, _ := cmd.Flags().GetString("stew")
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		path, _ := cmd.Flags().GetString("path")

		if name == "" && description == "" && path == "" {
			cmd.Help()
			return
		}

		stews := types.Stews{}

		err := stews.Load(types.GetHomeDir() + "/.stews.json")
		if err != nil {
			cmd.Println(err)
			return
		}

		if id == -1 {
			if selected_stew == "" {
				cmd.Help()
				return
			} else {
				stew, err := stews.GetByName(selected_stew)
				if err != nil {
					cmd.Println(err)
					return
				}
				err = stew.Edit(name, description, path)
				if err != nil {
					cmd.Println(err)
					return
				}
				err = stews.Save(types.GetHomeDir() + "/.stews.json")
				if err != nil {
					cmd.Println(err)
					return
				}
			}
		} else {
			s, err := stews.Get(id)
			if err != nil {
				cmd.Println(err)
				return
			}
			err = s.Edit(name, description, path)
			if err != nil {
				cmd.Println(err)
				return
			}
			err = stews.Save(types.GetHomeDir() + "/.stews.json")
			if err != nil {
				cmd.Println(err)
				return
			}
		}
	},
}

func flags() {
	EditCmd.Flags().IntP("id", "i", -1, "The id of the stew")
	EditCmd.Flags().StringP("stew", "s", "", "The stew you want to edit")
	EditCmd.Flags().StringP("name", "n", "", "The new name of the stew")
	EditCmd.Flags().StringP("description", "d", "", "The new description of the stew")
	EditCmd.Flags().StringP("path", "p", "", "The new path of the stew")
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flags()
}
