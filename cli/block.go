package cli

import (
	"behaviorlog-analyzer/data"
	"fmt"
)

// 方块日志相关的查询
var (
	actionType, block, action, dimension string
	startPosition, endPosition           string
)

func blockLogCli() {
	for {
		fmt.Println("开始进行方块日志查询")
		startTime, endTime := askForTime()
		player := askForPlayer()
		dimension := askForDimension()
		fmt.Printf("输入要查询的操作类型:\n[1]放置\n[2]破坏\n[留空]所有\n:")
		fmt.Scanln(&actionType)
		switch actionType {
		case "1":
			action = "Place"
		case "2":
			action = "Destroy"
		case "3":
			action = ""
		default:
			action = ""
		}
		sx, sy, sz, ex, ey, ez, _ := askForPosition()
		fmt.Printf("\n---------------------------------------------\n")
		if startTime != nil && endTime != nil {
			fmt.Printf("当前输入:\n日志类型:%s\n查询的玩家:%s\n操作类型:%s\n范围:(%d,%d,%d) - (%d,%d,%d)\n时间区间:[%s] ~ [%s]\n请确认是否进行查询,输入q重新进行设置\n:", logType, player, action, sx, sy, sz, ex, ey, ez, startTime, endTime)
		} else {
			fmt.Printf("当前输入:\n日志类型:%s\n查询的玩家:%s\n操作类型:%s\n范围:(%d,%d,%d) - (%d,%d,%d)\n时间区间:全时段\n请确认是否进行查询,输入q重新进行设置\n:", logType, player, action, sx, sy, sz, ex, ey, ez)
		}
		var input string
		// CheckValid
		if fmt.Scanln(&input); input == "q" {
			// valid = false
			break
		}
		queryLog := &data.BlockLog{
			Player:    player,
			Action:    action,
			Dimension: dimension,
		}
		blockLogs := data.GetBlockLog(sx, sy, sz, ex, ey, ez, nil, nil, queryLog)
		fmt.Println("查询到以下记录:")
		for _, bLog := range blockLogs {
			fmt.Println(bLog)
		}
	}
}
