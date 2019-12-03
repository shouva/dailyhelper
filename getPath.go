package dailyhelper

import (
	"fmt"
	"mime/multipart"
	"net/http"
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

// IsValidImage :
func IsValidImage(file multipart.File) bool {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		fmt.Println(err) // do something with that error
		return false
	}
	tipe := http.DetectContentType(buff)
	switch tipe {
	case "image/png", "image/jpeg", "image/jpg":
		return true
	default:
		fmt.Println(tipe)
		return false
	}
}
