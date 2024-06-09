package util

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fmt"
	"io"

	"github.com/spf13/viper"
)

const (
	red = "\033[31m"
)

func GetHomeDir() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir

}

func GetCurrentDir() string {
	cwd, _ := os.Getwd()
	return cwd
}

func FormatTime(t time.Time) string {
	// format yyyy-mm-dd hh:mm:ss
	return t.Format(viper.GetString("timeFormat"))

}

func GetPath(path string) (string, error) {
	dirs := strings.Split(path, "/")
	homeDir := GetCurrentDir()

	// check if path is absolute
	isPathAbsolute := filepath.IsAbs(path)

	if isPathAbsolute {
		return path, nil
	} else {
		path = homeDir
		for _, val := range dirs {
			path = filepath.Join(path, val)
		}
	}

	if !checkIfDirExists(path) {
		err := fmt.Sprintf("\n%v%v does not exist", red, path)
		return err, errors.New(err)
	}

	return path, nil
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
