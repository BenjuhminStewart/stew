/*
Copyright © 2024 Benjamin Stewart <benjuhminstewart@gmail.com
*/
package use

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BenjuhminStewart/stew/types"
)

// UseCmd represents the init command
var UseCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a stew instance to create a new project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := types.Stews{}
		err := s.Load(types.StewPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}

		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")

		if id == -1 && name == "" {
			cmd.Help()
			return
		}

		if id != -1 && name != "" {
			fmt.Println("You can only use one flag at a time")
			return
		}

		// get Stew by id or name
		if id != -1 {
			stew, err := s.Get(id)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = CopyDir(stew.Path, path)
			if err != nil {
				fmt.Println(err)
				return
			}

			newFolderName := strings.Split(stew.Path, "/")[len(strings.Split(stew.Path, "/"))-1]

			if checkIfDirExists(newFolderName) {
				fmt.Println("🎉 Project created successfully")
			} else {
				fmt.Println("\n❌ Project not created")
			}

		}

		if name != "" {
			stew, err := s.GetByName(name)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = CopyDir(stew.Path, path)

			newFolderName := strings.Split(stew.Path, "/")[len(strings.Split(stew.Path, "/"))-1]

			if checkIfDirExists(newFolderName) {
				fmt.Println("🎉 Project created successfully")
			} else {
				fmt.Println("\n❌ Project not created")
			}
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)

			return
		}

	},
}

func checkIfDirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		fmt.Println(err)
		return
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

func flags() {
	UseCmd.Flags().StringP("name", "n", "", "The name of the project")
	UseCmd.Flags().IntP("id", "i", -1, "The id of the stew")
	UseCmd.Flags().StringP("path", "p", getCWD(), "The path to the stew")
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
