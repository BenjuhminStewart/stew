/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package add

import (
	"fmt"

	"github.com/BenjuhminStewart/stew/types"
	"github.com/BenjuhminStewart/stew/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"os"
)

// AddCmd represents the create command
var AddCmd = &cobra.Command{
	Use:   "add <name_of_stew>",
	Short: "Add a stew based off a defined directory",
	Long:  `stew add <name_of_stew> [flags]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		name := args[0]
		description, _ := cmd.Flags().GetString("description")
		path, _ := cmd.Flags().GetString("path")

		st := types.Stews{}

		if err := st.Load(viper.GetString("stewsPath")); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// get absolute path
		path, _ = util.GetPath(path)

		// if description is empty, set it to no description provided
		if description == "" {
			description = "no description provided"
		}

		st.Add(name, description, path)
		err := st.Save(viper.GetString("stewsPath"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func getCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

func addFlag() {
	AddCmd.Flags().StringP("description", "d", "no description provided", "Description of the stew")
	AddCmd.Flags().StringP("path", "p", getCWD(), "Path to the stew")

}

func init() {
	addFlag()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
