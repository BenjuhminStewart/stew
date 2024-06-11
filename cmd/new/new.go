package new

import (
	"fmt"
	"github.com/BenjuhminStewart/stew/types"
	"github.com/BenjuhminStewart/stew/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmd represents the init command
var NewCmd = &cobra.Command{
	Use:   "new <name_of_stew>",
	Short: "Create a new instance of a stew template to create a new project",
	Long:  `stew new <name_of_stew> [flags]`,
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
		force, _ := cmd.Flags().GetBool("force")

		// get Stew by name
		stew, err := s.GetByName(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		if path == "" {
			path = util.GetCurrentDir()
		} else {
			path, err = util.GetPath(path)

			if err != nil && !viper.GetBool("allowFileCreation") && !force {
				fmt.Println(err)
				fmt.Println("To overwrite this, run the command with the -f or --force flag")
				fmt.Println("If you want to overwrite this by default, set allowFileCreation to true in your config file")
				return
			}
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
	NewCmd.Flags().StringP("path", "p", "", "The path to the stew (defaults to current directory)")
	NewCmd.Flags().BoolP("force", "f", false, "Force the creation of the project")
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
