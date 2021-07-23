package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 通用部分
func askForTime() (startTime *time.Time, endTime *time.Time) {
	for {
		var input, timeStr1, timeStr2 string
		fmt.Printf("是否需要设置时间区间?(y/n)\n:")
		fmt.Scanln(&input)
		if input == "y" {
			fmt.Printf("请输入时间区间起点 格式 2006-01-02 15:04:05\n:")
			// 注意scanln不能读取一串带空格的字符串（会分成两部分）
			var tmp1, tmp2 string
			fmt.Scanf("%s %s\n", &tmp1, &tmp2)
			timeStr1 = tmp1 + " " + tmp2
			fmt.Printf("请输入时间区间终点 格式 2006-01-02 15:04:05\n:")
			fmt.Scanf("%s %s\n", &tmp1, &tmp2)
			timeStr2 = tmp1 + " " + tmp2
			time1, err1 := time.Parse("2006-01-02 15:04:05", timeStr1)
			time2, err2 := time.Parse("2006-01-02 15:04:05", timeStr2)
			if err1 == nil && err2 == nil {
				if time1.Unix() <= time2.Unix() {
					startTime, endTime = &time1, &time2
				} else {
					startTime, endTime = &time2, &time1
				}
			} else {
				fmt.Println("时间解析失败")
				startTime, endTime = nil, nil
			}
		} else {
			fmt.Println("不设置时间区间")
			startTime, endTime = nil, nil
		}
		return
	}
}

func askForPlayer() (player string) {
	fmt.Printf("查询的玩家名(留空则表示所有玩家)\n")
	// 玩家名可能包含空格，需要特殊处理!
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		player = scanner.Text()
	}
	return
}

func askForDimension() (dimension string) {
	fmt.Printf("输入要查询的维度:\n[0]主世界\n[1]下界\n[2]末地\n[留空]所有\n:")
	fmt.Scanln(&dimension)
	if dimension != "0" && dimension != "1" && dimension != "2" {
		dimension = ""
	}
	return
}

func askForPosition() (sx int64, sy int64, sz int64, ex int64, ey int64, ez int64, valid bool) {
	fmt.Printf("输入要查询的操作对象(方块名,留空表示任何对象)\n:")
	fmt.Scanln(&block)

	fmt.Printf("输入查询范围的起点(格式: x,y,z) 留空使用-999999,0,-999999\n:")
	fmt.Scanln(&startPosition)
	sx, sy, sz, valid = positionParser(startPosition)
	if !valid {
		fmt.Println("输入无效,使用默认值 -999999,0,-999999")
		sx, sy, sz = -999999, 0, -999999
	}
	fmt.Printf("输入查询范围的终点(格式: x,y,z) 留空使用999999, 1024, 999999\n:")
	fmt.Scanln(&endPosition)
	ex, ey, ez, valid = positionParser(endPosition)
	if !valid {
		fmt.Println("输入无效,使用默认值 999999,1024,999999")
		ex, ey, ez = 999999, 1024, 999999
	}
	return
}

func positionParser(pos string) (x int64, y int64, z int64, valid bool) {
	valid = true
	var err1, err2, err3 error
	poss := strings.Split(pos, ",")
	if len(poss) != 3 {
		fmt.Println("无效坐标输入")
		valid = false
	} else {
		x, err1 = strconv.ParseInt(poss[0], 10, 64)
		y, err2 = strconv.ParseInt(poss[1], 10, 64)
		z, err3 = strconv.ParseInt(poss[2], 10, 64)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("无效坐标输入,坐标转换失败")
			valid = false
		}
	}
	return
}
