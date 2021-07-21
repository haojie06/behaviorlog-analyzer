package utils

import (
	"log"
	"os"
	"path/filepath"
)

func PathParse(path string) string {
	if !filepath.IsAbs(path) {
		currentWd, err := os.Getwd()
		CheckErr(err, "路径解析")
		path = filepath.Join(currentWd, path)
	}
	return path
}

func CheckErr(err error, where string) {
	if err != nil {
		log.Println(where, "出现异常", err.Error())
		panic(err)
	}
}
