package cli

import (
	"fmt"
)

var logType string

// 纯命令行交互
func StartCli() {
	for {
		fmt.Printf("查询模式选择\n[1]方块日志查询\n:")
		fmt.Scanln(&logType)
		switch logType {
		case "1":
			blockLogCli()
		default:
		}
	}
}
