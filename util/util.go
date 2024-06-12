package util

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"fmt"
	"io"

	"github.com/spf13/viper"
)

const (
	red       = "\033[31m"
	pathColor = "\033[33m"
	green     = "\033[32m"
	quoted    = "\033[35m"
	reset     = "\033[0m"
)

// GetHomeDir returns the home directory
func GetHomeDir() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir

}

// GetCurrentDir returns the current directory
func GetCurrentDir() string {
	cwd, _ := os.Getwd()
	return cwd
}

// FormatTime formats a time
func FormatTime(t time.Time) string {
	// format yyyy-mm-dd hh:mm:ss
	return t.Format(viper.GetString("timeFormat"))

}

// GetPath returns the path to a file
func GetPath(path string) (string, error) {
	dirs := strings.Split(path, "/")
	homeDir := GetCurrentDir()

	// check if path is absolute
	isPathAbsolute := filepath.IsAbs(path)

	if isPathAbsolute {
		return path, nil
	}

	path = homeDir
	for _, val := range dirs {
		path = filepath.Join(path, val)
	}

	if !CheckIfDirExists(path) {
		err := fmt.Sprintf("\n%v%v%v does not exist", red, path, reset)
		return path, errors.New(err)
	}

	return path, nil
}

// CheckIfDirExists checks if a directory exists
func CheckIfDirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// getCWD returns the current working directory
func getCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

// CopyFile copies a file from one location to another
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

// UpdateProjectName updates a project name in a file
func UpdateProjectName(path string, oldString string, newString string, ignoreCase bool) (int, error) {
	// file walker that goes through all files in the directory and replaces the replaceString with the projectName

	var filesChanged []string
	count := 0
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		read, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		for scanner.Scan() {
			line := scanner.Text()
			if ignoreCase {
				if caseInsensitiveContains(line, oldString) {
					regexString := fmt.Sprintf("(?i)%v", oldString)
					re := regexp.MustCompile(regexString)
					newContents := re.ReplaceAllString(string(read), newString)
					err = os.WriteFile(path, []byte(newContents), 0644)
					if err != nil {
						return err
					}
					filesChanged = append(filesChanged, path)
					count++
				}
			} else {
				if strings.Contains(line, oldString) {
					newContents := strings.Replace(string(read), oldString, newString, -1)
					err = os.WriteFile(path, []byte(newContents), 0644)
					if err != nil {
						return err
					}

					filesChanged = append(filesChanged, path)
					count++
				}
			}
		}
		err = file.Sync()
		if err != nil {
			return err
		}
		return nil
	})
	if len(filesChanged) != 0 {
		fmt.Printf("\nReplaced in:\n%v", filesReplacedString(filesChanged))
	}
	return count, nil
}

func caseInsensitiveContains(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func filesReplacedString(filesChanged []string) string {
	var str string
	for _, file := range filesChanged {
		str += fmt.Sprintf(" - %v%v%v\n", quoted, file, reset)
	}
	return str
}
