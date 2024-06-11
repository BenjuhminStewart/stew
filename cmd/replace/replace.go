/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package replace

import (
	"fmt"

	"github.com/BenjuhminStewart/stew/util"
	"github.com/spf13/cobra"
)

const (
	green      = "\033[32m"
	red        = "\033[31m"
	path_color = "\033[33m"
	quoted     = "\033[35m"
	reset      = "\033[0m"
)

// ReplaceCmd represents the replace command
var ReplaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace a project name in a stew",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			cmd.Help()
			return
		}
		old_string := args[0]
		new_string := args[1]

		path, _ := cmd.Flags().GetString("path")
		ignore_case, _ := cmd.Flags().GetBool("ignore-case")

		path, err := util.GetPath(path)
		if err != nil {
			cmd.Println(err)
			return
		}

		if path == "" {
			path = util.GetCurrentDir()
		}

		count, err := util.UpdateProjectName(path, old_string, new_string, ignore_case)
		if err != nil {
			cmd.Println(err)
			return
		}

		fmt.Printf("\nReplaced %v%v%v instances of %v%v%v with %v%v%v\n", green, count, reset, red, old_string, reset, quoted, new_string, reset)

	},
}

func flags() {
	ReplaceCmd.Flags().StringP("path", "p", "", "The path to the stew (defaults to current directory)")
	ReplaceCmd.Flags().BoolP("ignore-case", "i", false, "Whether to replace the string case sensitively")
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	flags()
}
