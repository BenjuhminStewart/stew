/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com>
*/
package edit

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BenjuhminStewart/stew/types"
)

// EditCmd represents the edit command
var EditCmd = &cobra.Command{
	Use:   "edit <name_of_stew>",
	Short: "Edit a stew's name, description, or path",
	Long:  `stew edit <name_of_stew> [flags]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		selected_stew := args[0]
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		path, _ := cmd.Flags().GetString("path")

		if name == "" && description == "" && path == "" {
			cmd.Help()
			return
		}

		stews := types.Stews{}

		err := stews.Load(viper.GetString("stewsPath"))
		if err != nil {
			cmd.Println(err)
			return
		}

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
		err = stews.Save(viper.GetString("stewsPath"))
		if err != nil {
			cmd.Println(err)
			return
		}
	},
}

func flags() {
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
