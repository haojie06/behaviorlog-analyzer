package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// 排序两个值，用于范围查询的条件
func ComparePos(x int64, y int64) (smallOne int64, bigOne int64) {
	if x <= y {
		smallOne = x
		bigOne = y
	} else {
		smallOne = y
		bigOne = x
	}
	return
}
