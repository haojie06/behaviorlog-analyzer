package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
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

func MetricTime(stageName string) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("%s Time cost = %v\n", stageName, tc)
	}
}
