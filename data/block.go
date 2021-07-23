package data

import (
	"behaviorlog-analyzer/utils"
	"time"

	"gorm.io/gorm"
)

// 方块日志相关查询
// 根据范围查询
func GetBlockLog(sx int64, sy int64, sz int64, ex int64, ey int64, ez int64, startTime *time.Time, endTime *time.Time, blockLog *BlockLog) (blockLogs []BlockLog) {
	defer utils.MetricTime("--------------------查询记录------------------")()
	sx, ex = utils.ComparePos(sx, ex)
	sy, ey = utils.ComparePos(sy, ey)
	sz, ez = utils.ComparePos(sz, ez)
	var tempDB *gorm.DB
	if startTime != nil && endTime != nil {
		tempDB = DB.Table("block_logs").Where("time >= ? and time <= ?", startTime, endTime)
	} else {
		tempDB = DB.Table("block_logs")
	}
	if blockLog != nil {
		// 结构体条件 GORM会忽略零值 除了时间范围和坐标范围的查询条件都可以使用
		tempDB.Where("position_x >= ? and position_y >= ? and position_z >= ? and position_x <= ? and position_y <= ? and position_z <= ?", sx, sy, sz, ex, ey, ez).Where(blockLog)
	} else {
		tempDB.Where("position_x >= ? and position_y >= ? and position_z >= ? and position_x <= ? and position_y <= ? and position_z <= ?", sx, sy, sz, ex, ey, ez)
	}
	tempDB.Find(&blockLogs)
	return
}

// 根据范围 时间范围查询

// 根据范围 时间范围 和 玩家名查询
