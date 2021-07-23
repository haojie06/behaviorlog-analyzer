package data

import "fmt"

type result struct {
	Player string
	Count  int
}

// 获取用户列表
func GetPlayerBlockActionCount(action string) []result {
	var res []result
	DB.Model(&BlockLog{}).Select("player, count(*) as count").Where("action = ?", action).Order("count desc").Group("player").Find(&res)
	for _, r := range res {
		fmt.Println(r)
	}
	return res
}
