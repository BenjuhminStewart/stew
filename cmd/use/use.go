/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package use

import (
	"fmt"
	"github.com/BenjuhminStewart/stew/types"
	"github.com/BenjuhminStewart/stew/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// UseCmd represents the init command
var UseCmd = &cobra.Command{
	Use:   "use <name_of_stew>",
	Short: "Use a stew instance to create a new project",
	Long:  `stew use <name_of_stew> [flags]`,
	Run: func(cmd *cobra.Command, args []string) {
		s := types.Stews{}
		err := s.Load(viper.GetString("stewsPath"))
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(args) == 0 {
			cmd.Help()
			return
		}

		name := args[0]
		path, _ := cmd.Flags().GetString("path")

		// get Stew by name
		stew, err := s.GetByName(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		if path == "" {
			path = util.GetCurrentDir()
		} else {
			path, _ = util.GetPath(path)
		}

		err = util.CopyDir(stew.Path, path)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("🎉 Project created successfully")

	},
}

func flags() {
	UseCmd.Flags().StringP("path", "p", "", "The path to the stew (defaults to current directory)")
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flags()
}
