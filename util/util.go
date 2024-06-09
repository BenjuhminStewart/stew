package util

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

func GetHomeDir() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir

}

func FormatTime(t time.Time) string {
	// format yyyy-mm-dd hh:mm:ss
	return t.Format(viper.GetString("timeFormat"))

}
