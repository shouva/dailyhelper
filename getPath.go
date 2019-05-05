package dailyhelper

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetCurrentPath :
func GetCurrentPath(isbuild bool) string {
	if isbuild {
		return getExePath()
	}
	return getSourcePath()
}
func getSourcePath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}
func getExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
