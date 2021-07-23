package file

import (
	"behaviorlog-analyzer/data"
	"behaviorlog-analyzer/utils"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 读取目录下的所有文件，提取有效信息 储存进sqlite 当前每次运行程序都要重新载入所有的数据，之后改为只添加增量
func LoadLogs(logDir string) (logs []interface{}) {
	defer utils.MetricTime("加载日志")()
	files, err := ioutil.ReadDir(logDir)
	utils.CheckErr(err, "加载日志")
	readCount := 0
	for _, file := range files {
		// 暂时不读取容器日志
		reg := regexp.MustCompile(`BlockLog.*`)
		if !file.IsDir() && reg.MatchString(file.Name()) {
			fmt.Println(file.Name())
			csvFile, err := os.Open(filepath.Join(logDir, file.Name()))
			utils.CheckErr(err, "打开CSV文件")
			reader := csv.NewReader(bufio.NewReader(csvFile))
			// var logs []data.LogItem
			for {
				line, err := reader.Read()
				readCount++
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("出现错误,当前读取文件:%s\n", file.Name())
					fmt.Println(line)
					utils.CheckErr(err, "读取CSV记录")
				}
				t, err := parseTime(line[0])
				utils.CheckErr(err, "时间解析"+line[0])
				// fmt.Println(line, t)
				// [2021-06-30 20:44:53 Place TanisGon 0 (-7438, 46, -11905) 放置 stone]
				posX, posY, posZ := parseLocation(line[4])
				logs = append(logs, data.BlockLog{
					Time:      t,
					Action:    line[1],
					Player:    line[2],
					Dimension: line[3],
					Position:  line[4],
					Target:    line[6],
					PositionX: posX,
					PositionY: posY,
					PositionZ: posZ,
				})
			}
		}
	}
	fmt.Printf("完成日志文件读取，共读取%d行\n开始载入数据库...", readCount)
	data.SaveLogs(logs)
	return
}

// 暂时没有处理时区问题
func parseTime(rawTime string) (t time.Time, err error) {
	rawTime = strings.Replace(rawTime, "\uFEFF", "", -1)
	t, err = time.Parse("2006-01-02 15:04:05", rawTime)
	return
}

func parseLocation(location string) (int64, int64, int64) {
	pos := strings.Split(location[1:len(location)-1], ", ")
	x, err := strconv.ParseInt(pos[0], 10, 32)
	utils.CheckErr(err, "x坐标转换")
	y, err := strconv.ParseInt(pos[1], 10, 32)
	utils.CheckErr(err, "y坐标转换")
	z, err := strconv.ParseInt(pos[2], 10, 32)
	utils.CheckErr(err, "z坐标转换")

	return x, y, z
}
