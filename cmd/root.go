/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/BenjuhminStewart/stew/cmd/add"
	"github.com/BenjuhminStewart/stew/cmd/edit"
	"github.com/BenjuhminStewart/stew/cmd/get"
	"github.com/BenjuhminStewart/stew/cmd/list"
	"github.com/BenjuhminStewart/stew/cmd/new"
	"github.com/BenjuhminStewart/stew/cmd/remove"
	"github.com/BenjuhminStewart/stew/util"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stew",
	Short: "A project template manager",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setDefaults() {
	viper.SetDefault("stewsPath", util.GetHomeDir()+"/.stews.json")
	viper.SetDefault("timeFormat", "2006-01-02 15:04:05")
	viper.SetDefault("allowFileCreation", false)
}

func addSubCommands() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(edit.EditCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(get.GetCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/stew/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add sub commands
	addSubCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := home + "/.config/stew"

		// Search config in home directory with name ".stew" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		setDefaults()
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
