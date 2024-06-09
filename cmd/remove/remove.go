/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package remove

import (
	"github.com/BenjuhminStewart/stew/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RemoveCmd represents the delete command
var RemoveCmd = &cobra.Command{
	Use:   "remove <name_of_stew>",
	Short: "Remove a stew",
	Long:  `stew remove <name_of_stew> [flags]`,
	Run: func(cmd *cobra.Command, args []string) {
		s := types.Stews{}
		err := s.Load(viper.GetString("stewsPath"))
		if err != nil {
			cmd.Println(err)
			return
		}

		if len(args) == 0 {
			cmd.Help()
			return
		}

		name := args[0]
		_, err = s.GetByName(name)
		if err != nil {
			cmd.Println(err)
			return
		}

		err = s.RemoveByName(name)
		if err != nil {
			cmd.Println(err)
			return
		}

		err = s.Save(viper.GetString("stewsPath"))
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
