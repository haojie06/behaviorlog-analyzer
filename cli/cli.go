package cli

import (
	"behaviorlog-analyzer/data"
	"fmt"
)

var logType string

// 纯命令行交互
func StartCli() {
	for {
		fmt.Printf("查询模式选择\n[1]方块日志查询\n[2]破坏、放置榜单\n:")
		fmt.Scanln(&logType)
		switch logType {
		case "1":
			blockLogCli()
		case "2":
			fmt.Println("-----------------放置榜-----------------")
			data.GetPlayerBlockActionCount("Place")
			fmt.Println("-----------------破坏榜-----------------")
			data.GetPlayerBlockActionCount("Place")
		default:
		}
	}
}
