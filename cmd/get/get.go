package get

import (
	"fmt"

	"github.com/BenjuhminStewart/stew/types"
	"github.com/BenjuhminStewart/stew/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get <name_of_stew>",
	Short: "Get a stew from a given name or id",
	Long:  `stew get <name_of_stew> [flags]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}

		name := args[0]

		s := types.Stews{}
		err := s.Load(viper.GetString("stewsPath"))
		if err != nil {
			fmt.Println(err)
			return
		}

		stew, err := s.GetByName(name)

		if err != nil {
			fmt.Println(err)
			return
		}

		stew.Print()

		if cmd.Flag("tree").Value.String() == "true" {
			util.PrintTree(stew.Path)
		}
	},
}

func flags() {
	GetCmd.Flags().BoolP("tree", "t", false, "Print the tree of the stew")
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
