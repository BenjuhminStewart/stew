/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package add

import (
	"fmt"
	"github.com/BenjuhminStewart/stew/types"
	"github.com/spf13/cobra"

	"os"
)

// AddCmd represents the create command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a stew based off a defined directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			cmd.Help()
			return
		}
		description, _ := cmd.Flags().GetString("description")
		path, _ := cmd.Flags().GetString("path")
		usesGit, _ := cmd.Flags().GetBool("git")

		st := types.Stews{}

		if err := st.Load(types.GetHomeDir() + "/.stews.json"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		st.Add(name, description, path, usesGit)
		err := st.Save(types.GetHomeDir() + "/.stews.json")
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
	AddCmd.Flags().StringP("name", "n", "", "Name of the stew")
	AddCmd.Flags().StringP("description", "d", "no description provided", "Description of the stew")
	AddCmd.Flags().StringP("path", "p", getCWD(), "Path to the stew")
	AddCmd.Flags().BoolP("git", "g", false, "If the stew uses git")

	// required flags
	AddCmd.MarkFlagRequired("name")
	AddCmd.MarkFlagRequired("description")
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
