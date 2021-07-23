package data

import (
	"behaviorlog-analyzer/utils"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	SqlDB *sql.DB
)

func InitDB(memDB bool) {
	// 0 文件数据库 1 内存数据库
	var db *gorm.DB
	var err error
	if memDB {
		fmt.Println("内存数据库")
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	} else {
		fmt.Println("文件数据库")
		db, err = gorm.Open(sqlite.Open("log.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	}
	utils.CheckErr(err, "开启sqlite数据库1")
	DB = db
	SqlDB, _ := db.DB()
	err = SqlDB.Ping()
	utils.CheckErr(err, "开启sqlite数据库2")
	DB.Exec("DROP TABLE block_logs")
	DB.AutoMigrate(&BlockLog{})
	log.Println("数据库初始化完成")
}

// 一次传递的切片里面的类型
func SaveLogs(is []interface{}) (err error) {
	defer utils.MetricTime("加载日志->日志载入数据库")()
	var blockLogs []BlockLog
	for _, i := range is {
		switch logRecord := i.(type) {
		case BlockLog:
			blockLogs = append(blockLogs, logRecord)
		default:
		}
	}
	// 日志切片太长无法一次写入，所以切片处理
	DB.CreateInBatches(&blockLogs, 1000)
	return
}
